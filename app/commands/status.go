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

// Status queries for bulb's status
func Status() {
	device, err := util.OpenHCI()

	if err != nil {
		os.Exit(util.ExitStatusHCIError)
	}

	done := make(chan bool)

	device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
		if strings.ToUpper(*cli.StatusTarget) == strings.ToUpper(peripheral.ID()) {
			log.Info.Println("Device found")

			device.Handle(gatt.PeripheralConnected(func(peripheral gatt.Peripheral, err error) {
				defer device.CancelConnection(peripheral)

				characteristic, err := util.GetCharacteristic(peripheral, gatt.MustParseUUID("FFD5"), gatt.MustParseUUID("FFD9"))

				if (err != nil) || (nil == characteristic) {
					log.Error.Printf("Failed to get characteristic: %v\n", err)
					os.Exit(util.ExitStatusGenericError)
				}

				listen, err := util.GetCharacteristicWithDescriptors(peripheral, gatt.MustParseUUID("FFD0"), gatt.MustParseUUID("FFD4"))

				if (err != nil) || (nil == listen) {
					log.Error.Printf("Failed to get listen characteristic: %v\n", err)
					os.Exit(util.ExitStatusGenericError)
				}

				peripheral.SetNotifyValue(listen, func(characteristic *gatt.Characteristic, data []byte, err error) {
					if data[2] == 0x23 {
						if data[3] == 0x41 {
							if data[9] != 0 {
								// White
								log.Info.Printf("White: %#02X\n", data[9])
							} else {
								// Color
								log.Info.Printf("Red: %#02X, Green: %#02X, Blue: %#02X\n", data[6], data[7], data[8])
							}
						} else {
							// Built-in mode
							log.Info.Printf("Built-in mode: %#02X, Speed: %#02X\n", data[3], data[5])
						}
					} else {
						log.Info.Println("The bulb is powered off")
					}

					done <- true
				})
				peripheral.WriteCharacteristic(characteristic, []byte{0xEF, 0x01, 0x77}, false)
			}))

			device.StopScanning()
			device.Connect(peripheral)
		}
	}))

	log.Info.Println("Scanning devices")
	device.Scan([]gatt.UUID{}, false)

	select {
	case <-time.After(*cli.OperationTimeout):
		log.Error.Println("Failed to query status for target device")
		os.Exit(util.ExitStatusGenericError)
	case <-done:
		log.Info.Println("Status queried successfully")
	}
}
