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

// Mode sets mode
func Mode() {
	if (*cli.ModeMode < 0x25) || (*cli.ModeMode > 0x38) {
		log.Error.Println("Unknown mode")
		os.Exit(util.ExitStatusGenericError)
	}

	device, err := util.OpenHCI()

	if err != nil {
		os.Exit(util.ExitStatusHCIError)
	}

	done := make(chan bool)

	device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
		if strings.ToUpper(*cli.ModeTarget) == strings.ToUpper(peripheral.ID()) {
			log.Info.Println("Device found")

			device.Handle(gatt.PeripheralConnected(func(peripheral gatt.Peripheral, err error) {
				defer device.CancelConnection(peripheral)

				characteristic, err := util.GetCharacteristic(peripheral, gatt.MustParseUUID("FFD5"), gatt.MustParseUUID("FFD9"))

				if (err != nil) || (nil == characteristic) {
					log.Error.Printf("Failed to get characteristic: %v\n", err)
					os.Exit(util.ExitStatusGenericError)
				}

				peripheral.WriteCharacteristic(characteristic, []byte{0xBB, *cli.ModeMode, *cli.ModeSpeed, 0x44}, false)

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
		log.Error.Println("Failed to set mode for target device")
		os.Exit(util.ExitStatusGenericError)
	case <-done:
		log.Info.Println("Mode set successfully")
	}
}
