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

type peripheralResult struct {
	peripheral    gatt.Peripheral
	advertisement *gatt.Advertisement
	rssi          int
}

// Dump dumps target BLE bulb to stdout
func Dump() {
	device, err := util.OpenHCI()

	if err != nil {
		os.Exit(util.ExitStatusHCIError)
	}

	peripheralFound := make(chan *peripheralResult)
	done := make(chan bool)

	device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
		if strings.ToUpper(*cli.DumpTarget) == strings.ToUpper(peripheral.ID()) {
			peripheralFound <- &peripheralResult{
				peripheral:    peripheral,
				advertisement: advertisement,
				rssi:          rssi,
			}
		}
	}))
	device.Handle(gatt.PeripheralConnected(func(peripheral gatt.Peripheral, err error) {
		services, err := peripheral.DiscoverServices(nil)

		if err != nil {
			log.Error.Printf("Service discovering error: %v\n", err)
			return
		}

		for _, service := range services {
			msg := "Service: " + service.UUID().String()

			if len(service.Name()) > 0 {
				msg += " (" + service.Name() + ")"
			}
			log.Info.Println(msg)

			characteristics, err := peripheral.DiscoverCharacteristics(nil, service)

			if err != nil {
				log.Error.Printf("Characteristics discovering error: %v\n", err)
				continue
			}

			for _, characteristic := range characteristics {
				msg := "\tCharacteristic:  " + characteristic.UUID().String()

				if len(characteristic.Name()) > 0 {
					msg += " (" + characteristic.Name() + ")"
				}
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
					msg := "\t\tDescriptor: " + descriptor.UUID().String()

					if len(descriptor.Name()) > 0 {
						msg += " (" + descriptor.Name() + ")"
					}
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

		peripheral.Device().CancelConnection(peripheral)
		done <- true
	}))

	log.Info.Println("Scanning devices")
	device.Scan([]gatt.UUID{}, false)

	select {
	case <-time.After(*cli.ScanPeriod):
		log.Error.Println("Target device not found")
		device.StopScanning()
		os.Exit(util.ExitStatusTargetDeviceNotFound)
	case peripheralResult := <-peripheralFound:
		device.StopScanning()
		log.Info.Println("Target found:")
		log.Info.Printf("\tID: %s, Name: %v\n", peripheralResult.peripheral.ID(), peripheralResult.peripheral.Name())
		log.Info.Printf("\tLocal Name: %v\n", peripheralResult.advertisement.LocalName)
		log.Info.Printf("\tTX Power Level: %v\n", peripheralResult.advertisement.TxPowerLevel)
		log.Info.Printf("\tManufacturer Data: %v\n", peripheralResult.advertisement.ManufacturerData)
		log.Info.Printf("\tService Data: %v\n", peripheralResult.advertisement.ServiceData)
		device.Connect(peripheralResult.peripheral)
	}
	select {
	case <-time.After(*cli.ScanPeriod):
		log.Error.Println("Target device failed to be dumped")
		os.Exit(util.ExitStatusTargetDeviceFailedToDump)
	case <-done:
	}
}
