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

// Time sets device's time
func Time() {
	device, err := util.OpenHCI()

	if err != nil {
		os.Exit(util.ExitStatusHCIError)
	}

	done := make(chan bool)

	device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
		if strings.ToUpper(*cli.TimeTarget) == strings.ToUpper(peripheral.ID()) {
			log.Info.Println("Device found")

			device.Handle(gatt.PeripheralConnected(func(peripheral gatt.Peripheral, err error) {
				defer device.CancelConnection(peripheral)

				characteristic, err := util.GetCharacteristic(peripheral, gatt.MustParseUUID("FFD5"), gatt.MustParseUUID("FFD9"))

				if (err != nil) || (nil == characteristic) {
					log.Error.Printf("Failed to get characteristic: %v\n", err)
					os.Exit(util.ExitStatusGenericError)
				}

				now := time.Unix(*cli.TimeTime, 0)
				dayOfWeek := now.Weekday()

				if dayOfWeek == 0 {
					dayOfWeek = 7
				}

				log.Info.Printf("Time will be set to %v\n", now)
				peripheral.WriteCharacteristic(characteristic, []byte{
					0x10,
					byte(now.Year() / 100),
					byte(now.Year() % 100),
					byte(now.Month()),
					byte(now.Day()),
					byte(now.Hour()),
					byte(now.Minute()),
					byte(now.Second()),
					byte(dayOfWeek),
					0x00,
					0x01,
				}, false)

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
		log.Error.Println("Failed to set time for target device")
		os.Exit(util.ExitStatusGenericError)
	case <-done:
		log.Info.Println("Time set successfully")
	}
}
