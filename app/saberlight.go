package main

import (
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"log"
	"fmt"
)

// TODO: Scan for 5 seconds then stop. Report only unique
func main() {
	device, err := gatt.NewDevice(option.DefaultClientOptions...)

	if err != nil {
		log.Fatalf("Failed to open device: %s", err)
	}

	device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, a *gatt.Advertisement, rssi int) {
		fmt.Printf("\nPeripheral ID:%s, NAME:(%s)\n", peripheral.ID(), peripheral.Name())
		fmt.Println("  RSSI              =", rssi)
		fmt.Println("  Connectable       =", a.Connectable)
		fmt.Println("  Local Name        =", a.LocalName)
		fmt.Println("  Manufacturer Data =", a.ManufacturerData)
		fmt.Println("  Overflow Service  =", a.OverflowService)
		fmt.Println("  Service Data      =", a.ServiceData)
		fmt.Println("  Services          =", a.Services)
		fmt.Println("  Solicited Service =", a.SolicitedService)
		fmt.Println("  TX Power Level    =", a.TxPowerLevel)
	}))
	device.Init(func(device gatt.Device, state gatt.State) {
		fmt.Println("State:", state)
		switch state {
		case gatt.StatePoweredOn:
			fmt.Println("scanning...")
			device.Scan(make([]gatt.UUID, 0), false)
			return
		default:
			device.StopScanning()
		}
	})
	select {}
}
