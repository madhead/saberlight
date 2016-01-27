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

	device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
		if strings.ToUpper(*cli.DumpTarget) == strings.ToUpper(peripheral.ID()) {
			peripheralFound <- &peripheralResult{
				peripheral:    peripheral,
				advertisement: advertisement,
				rssi:          rssi,
			}
		}
	}))
	device.Handle(gatt.PeripheralConnected(func(p gatt.Peripheral, err error) {
		fmt.Println("Connected")
		defer p.Device().CancelConnection(p)

		if err := p.SetMTU(500); err != nil {
			fmt.Printf("Failed to set MTU, err: %s\n", err)
		}

		// Discovery services
		ss, err := p.DiscoverServices(nil)
		if err != nil {
			fmt.Printf("Failed to discover services, err: %s\n", err)
			return
		}

		for _, s := range ss {
			msg := "Service: " + s.UUID().String()
			if len(s.Name()) > 0 {
				msg += " (" + s.Name() + ")"
			}
			fmt.Println(msg)

			// Discovery characteristics
			cs, err := p.DiscoverCharacteristics(nil, s)
			if err != nil {
				fmt.Printf("Failed to discover characteristics, err: %s\n", err)
				continue
			}

			for _, c := range cs {
				msg := "  Characteristic  " + c.UUID().String()
				if len(c.Name()) > 0 {
					msg += " (" + c.Name() + ")"
				}
				msg += "\n    properties    " + c.Properties().String()
				fmt.Println(msg)

				// Read the characteristic, if possible.
				if (c.Properties() & gatt.CharRead) != 0 {
					b, err := p.ReadCharacteristic(c)
					if err != nil {
						fmt.Printf("Failed to read characteristic, err: %s\n", err)
						continue
					}
					fmt.Printf("    value         %x | %q\n", b, b)
				}

				// Discovery descriptors
				ds, err := p.DiscoverDescriptors(nil, c)
				if err != nil {
					fmt.Printf("Failed to discover descriptors, err: %s\n", err)
					continue
				}

				for _, d := range ds {
					msg := "  Descriptor      " + d.UUID().String()
					if len(d.Name()) > 0 {
						msg += " (" + d.Name() + ")"
					}
					fmt.Println(msg)

					// Read descriptor (could fail, if it's not readable)
					b, err := p.ReadDescriptor(d)
					if err != nil {
						fmt.Printf("Failed to read descriptor, err: %s\n", err)
						continue
					}
					fmt.Printf("    value         %x | %q\n", b, b)
				}

				// Subscribe the characteristic, if possible.
				if (c.Properties() & (gatt.CharNotify | gatt.CharIndicate)) != 0 {
					f := func(c *gatt.Characteristic, b []byte, err error) {
						fmt.Printf("notified: % X | %q\n", b, b)
					}
					if err := p.SetNotifyValue(c, f); err != nil {
						fmt.Printf("Failed to subscribe characteristic, err: %s\n", err)
						continue
					}
				}

			}
			fmt.Println()
		}

		fmt.Printf("Waiting for 5 seconds to get some notifiations, if any.\n")
		time.Sleep(5 * time.Second)
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

	time.Sleep(10 * time.Second)
}
