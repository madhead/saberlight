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

// Timings gets active device's timings
func Timings() {
	device, err := util.OpenHCI()

	if err != nil {
		os.Exit(util.ExitStatusHCIError)
	}

	done := make(chan bool)

	device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
		if strings.ToUpper(*cli.TimingsTarget) == strings.ToUpper(peripheral.ID()) {
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

				chunks := make(chan []byte)
				var data []byte

				peripheral.SetNotifyValue(listen, func(characteristic *gatt.Characteristic, data []byte, err error) {
					chunks <- data

					// 0x52 indicated end of timings data
					if data[len(data)-1] == 0x52 {
						close(chunks)
					}
				})
				peripheral.WriteCharacteristic(characteristic, []byte{0x24, 0x2A, 0x2B, 0x42}, false)

				// Wait for the timings data
				for chunk := range chunks {
					data = append(data, chunk...)
				}
				for i := 0; i < 6; i++ {
					log.Info.Printf("Timing #%v: Days: %v, Hour: %02v, Minute: %02v, Turn on?: %v, Open?: %v, \n", i+1, days(data[i*14+8]), data[i*14+5], data[i*14+6], data[i*14+14] == 240, data[i*14+1] == 240)
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
		log.Error.Println("Failed to get timings for target device")
		os.Exit(util.ExitStatusGenericError)
	case <-done:
		log.Info.Println("Timings queried successfully")
	}
}

/*
 */
const (
	all = 254
	mon = 2
	tue = 4
	wed = 8
	thu = 16
	fri = 32
	sat = 64
	sun = 128
)

func days(days byte) []string {
	if days == all {
		return []string{"ALL"}
	}

	var result []string

	if days&mon != 0 {
		result = append(result, "MON")
	}
	if days&tue != 0 {
		result = append(result, "TUE")
	}
	if days&wed != 0 {
		result = append(result, "WED")
	}
	if days&thu != 0 {
		result = append(result, "THU")
	}
	if days&fri != 0 {
		result = append(result, "FRI")
	}
	if days&sat != 0 {
		result = append(result, "SAT")
	}
	if days&sun != 0 {
		result = append(result, "SUN")
	}

	return result
}
