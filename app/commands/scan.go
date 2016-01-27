package commands

import (
	"os"
	"time"

	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/util"
	"github.com/madhead/saberlight/app/util/log"

	"github.com/paypal/gatt"
)

// Scan scans for nearby BLE bulbs
func Scan() {
	device, err := util.OpenHCI()

	if err != nil {
		os.Exit(util.ExitStatusHCIError)
	}

	peripherals := make(map[string]gatt.Peripheral)

	device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
		peripherals[peripheral.ID()] = peripheral
	}))

	log.Info.Println("Scanning devices")
	device.Scan([]gatt.UUID{}, false)
	time.Sleep(*cli.ScanPeriod)
	device.StopScanning()
	log.Info.Println("Scan results:")
	for id, peripheral := range peripherals {
		log.Info.Printf("ID: %s, Name: %v\n", id, peripheral.Name())
	}
}
