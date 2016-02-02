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

// On turns bulb on
var On = powerManagementFunction(cli.OnTarget, 0x23)

// Off turns bulb off
var Off = powerManagementFunction(cli.OffTarget, 0x24)

func powerManagementFunction(target *string, status byte) func() {
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

					peripheral.WriteCharacteristic(characteristic, []byte{0xCC, status, 0x33}, false)

					done <- true
				}))

				device.StopScanning()
				device.Connect(peripheral)
			}
		}))

		log.Info.Println("Scanning devices")
		device.Scan([]gatt.UUID{}, false)

		select {
		case <-time.After(*cli.ScanPeriod):
			log.Error.Println("Failed to change bulb's power status")
			os.Exit(util.ExitStatusGenericError)
		case <-done:
			log.Info.Println("Power status changed successfully")
		}
	}
}
