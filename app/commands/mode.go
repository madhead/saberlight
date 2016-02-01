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
			device.Handle(gatt.PeripheralConnected(func(peripheral gatt.Peripheral, err error) {
				defer device.CancelConnection(peripheral)

				services, err := peripheral.DiscoverServices(nil)

				if err != nil {
					log.Error.Printf("Failed to discover services: %v\n", err)
					os.Exit(util.ExitStatusGenericError)
				}

				for _, service := range services {
					if service.UUID().Equal(gatt.MustParseUUID("FFD5")) {
						characteristics, err := peripheral.DiscoverCharacteristics(nil, service)

						if err != nil {
							log.Error.Printf("Failed to discover characteristics: %v\n", err)
							os.Exit(util.ExitStatusGenericError)
						}

						for _, characteristic := range characteristics {
							if characteristic.UUID().Equal(gatt.MustParseUUID("FFD9")) {
								peripheral.WriteCharacteristic(characteristic, []byte{0xBB, *cli.ModeMode, *cli.ModeSpeed, 0x44}, false)

								break
							}
						}

						break
					}
				}

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
