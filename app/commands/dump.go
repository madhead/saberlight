package commands

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/util"
	"github.com/madhead/saberlight/app/util/log"

	"github.com/paypal/gatt"
)

// Dump dumps target BLE bulb to stdout
func Dump() {
	device, err := util.OpenHCI()

	if err != nil {
		os.Exit(util.ExitStatusHCIError)
	}

	done := make(chan bool)

	device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
		if strings.ToUpper(*cli.DumpTarget) == strings.ToUpper(peripheral.ID()) {
			log.Info.Println("Target found:")
			log.Info.Printf("\tID: %s, Name: %v\n", peripheral.ID(), peripheral.Name())
			log.Info.Printf("\tLocal Name: %v\n", advertisement.LocalName)
			log.Info.Printf("\tTX Power Level: %v\n", advertisement.TxPowerLevel)
			log.Info.Printf("\tManufacturer Data: %v\n", advertisement.ManufacturerData)
			log.Info.Printf("\tService Data: %v\n", advertisement.ServiceData)

			device.Handle(gatt.PeripheralConnected(func(peripheral gatt.Peripheral, err error) {
				defer device.CancelConnection(peripheral)

				services, err := peripheral.DiscoverServices(nil)

				if err != nil {
					log.Error.Printf("Service discovering error: %v\n", err)
					return
				}

				for _, service := range services {
					msg := fmt.Sprintf("Service: %v [%+v] (%v)", service.UUID(), service.UUID(), service.Name())

					log.Info.Println(msg)

					characteristics, err := peripheral.DiscoverCharacteristics(nil, service)

					if err != nil {
						log.Error.Printf("Characteristics discovering error: %v\n", err)
						continue
					}

					for _, characteristic := range characteristics {
						msg := fmt.Sprintf("\tCharacteristic: %v [%+v] (%v)", characteristic.UUID(), characteristic.UUID(), characteristic.Name())

						log.Info.Println(msg)
						log.Info.Printf("\t\tProperties: %s\n", characteristic.Properties().String())

						if (characteristic.Properties() & gatt.CharRead) != 0 {
							value, err := peripheral.ReadCharacteristic(characteristic)

							if err != nil {
								log.Error.Printf("Failed to read characteristic: %v\n", err)
								continue
							}

							log.Info.Printf("\t\tValue: %X (%q)\n", value, value)
						}

						// Discovery descriptors
						descriptors, err := peripheral.DiscoverDescriptors(nil, characteristic)

						if err != nil {
							log.Error.Printf("Failed to discover descriptors: %v\n", err)
							continue
						}

						for _, descriptor := range descriptors {
							msg := fmt.Sprintf("\t\tDescriptor: %v [%+v] (%v)", descriptor.UUID(), descriptor.UUID(), descriptor.Name())

							log.Info.Println(msg)

							value, err := peripheral.ReadDescriptor(descriptor)

							if err != nil {
								fmt.Printf("Failed to read descriptor, err: %s\n", err)
								continue
							}
							log.Info.Printf("\t\t\tValue: %X (%q)\n", value, value)
						}
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
	case <-time.After(*cli.OperationTimeout):
		log.Error.Println("Failed to dump target device")
		os.Exit(util.ExitStatusGenericError)
	case <-done:
		log.Info.Println("Target device dumped")
	}
}
