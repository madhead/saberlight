package commands

import (
	"os"
	"strings"
	"time"

	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/util"
	"github.com/madhead/saberlight/app/util/log"

	"github.com/paypal/gatt"
)

var zero = byte(0)

// Color sets color
var Color = colorFunction(cli.ColorTarget, cli.ColorRed, cli.ColorGreen, cli.ColorBlue, &zero, false)

// White sets white color with given intensity
var White = colorFunction(cli.WhiteTarget, &zero, &zero, &zero, cli.WhiteIntensity, true)

func colorFunction(target *string, red, green, blue, intensity *byte, white bool) func() {
	whiteIndicator := byte(0xF0)
	if white {
		whiteIndicator = byte(0x0F)
	}

	return func() {
		device, err := util.OpenHCI()

		if err != nil {
			os.Exit(util.ExitStatusHCIError)
		}

		done := make(chan bool)

		device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
			if strings.ToUpper(*target) == strings.ToUpper(peripheral.ID()) {
				log.Info.Println("Device found")

				device.Handle(gatt.PeripheralConnected(func(peripheral gatt.Peripheral, err error) {
					defer device.CancelConnection(peripheral)

					characteristic, err := util.GetCharacteristic(peripheral, gatt.MustParseUUID("FFD5"), gatt.MustParseUUID("FFD9"))

					if (err != nil) || (nil == characteristic) {
						log.Error.Printf("Failed to get characteristic: %v\n", err)
						os.Exit(util.ExitStatusGenericError)
					}

					peripheral.WriteCharacteristic(characteristic, []byte{0x56, *red, *green, *blue, *intensity, whiteIndicator, 0xAA}, false)

					done <- true
				}))

				device.StopScanning()
				device.Connect(peripheral)
			}
		}))

		log.Info.Println("Scanning devices")
		device.Scan([]gatt.UUID{}, false)

		select {
		case <-time.After(*cli.OperationTimeout):
			log.Error.Println("Failed to set bulb's color")
			os.Exit(util.ExitStatusGenericError)
		case <-done:
			log.Info.Println("Color changed successfully")
		}
	}
}
