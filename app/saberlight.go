package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/paypal/gatt"
)

var (
	app = kingpin.New("saberlight", `A tool to control some "smart" BLE bulbs.`)

	// Commong flags
	deviceTimeout = app.Flag("deviceTimeout", "Maximum wait time for HCI to start").Default("3s").Duration()

	//TODO: should scan only for supported BLE devices
	scan       = app.Command("scan", "Scan for BLE devices")
	scanPeriod = scan.Flag("scanPeriod", "Scan period").Default("5s").Duration()

	dump       = app.Command("dump", "Dump BLE device")
	dumpTarget = dump.Arg("id", "Device to dump").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case scan.FullCommand():
		//TODO: Sometimes gatt.NewDevice blocks indefinitely:
		//  h, err := linux.NewHCI(d.devID, d.chkLE, d.maxConn)
		//  h.resetDevice()
		//  if err := h.c.SendAndCheckResp(s, []byte{0x00}); err != nil {
		//..due to device doesn't respond to a command (e.g. opReset = hostCtl<<10 | 0x0003 // Reset)
		device, err := gatt.NewDevice([]gatt.Option{
			gatt.LnxDeviceID(-1, true),
			gatt.LnxMaxConnections(1),
		}...)

		if err != nil {
			log.Fatalf("Failed to open device: %v\n", err)
		}

		deviceReady := make(chan bool)

		device.Init(func(device gatt.Device, state gatt.State) {
			if state == gatt.StatePoweredOn {
				deviceReady <- true
				close(deviceReady)
			}
		})

		select {
		case <-time.After(*deviceTimeout):
			log.Fatalln("HCI device timed out")
		case <-deviceReady:
			log.Println("HCI device is ready")
		}

		peripherals := make(map[string]gatt.Peripheral)

		device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
			peripherals[peripheral.ID()] = peripheral
		}))

		log.Println("Scanning devices")
		device.Scan([]gatt.UUID{}, false)
		time.Sleep(*scanPeriod)
		device.StopScanning()
		log.Println("Scan results:")
		for id, peripheral := range peripherals {
			log.Printf("ID: %s, Name: %v\n", id, peripheral.Name())
		}
	case dump.FullCommand():
		//TODO
		fmt.Println("Dumping")
	}
}
