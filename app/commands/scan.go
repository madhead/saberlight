package commands

import (
	"log"
	"time"

	"github.com/madhead/saberlight/app/cli"

	"github.com/paypal/gatt"
)

//TODO: should scan only for supported BLE devices
func Scan() {
	deviceReady := make(chan gatt.Device)

	// Sometimes gatt.NewDevice blocks indefinitely:
	// h, err := linux.NewHCI(d.devID, d.chkLE, d.maxConn)
	// h.resetDevice()
	// if err := h.c.SendAndCheckResp(s, []byte{0x00}); err != nil {
	// ...due to device doesn't respond to a command (e.g. opReset = hostCtl<<10 | 0x0003 // Reset)
	// so wrap this in gorutine and timeout it
	go func() {
		device, err := gatt.NewDevice([]gatt.Option{
			gatt.LnxDeviceID(-1, true),
			gatt.LnxMaxConnections(1),
		}...)

		if err != nil {
			log.Fatalf("Failed to open device: %v\n", err)
		}

		device.Init(func(device gatt.Device, state gatt.State) {
			if state == gatt.StatePoweredOn {
				deviceReady <- device
				close(deviceReady)
			}
		})
	}()

	select {
	case <-time.After(*cli.DeviceTimeout):
		log.Fatalln("HCI device timed out")
	case device := <-deviceReady:
		peripherals := make(map[string]gatt.Peripheral)

		device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
			peripherals[peripheral.ID()] = peripheral
		}))

		log.Println("Scanning devices")
		device.Scan([]gatt.UUID{}, false)
		time.Sleep(*cli.ScanPeriod)
		device.StopScanning()
		log.Println("Scan results:")
		for id, peripheral := range peripherals {
			log.Printf("ID: %s, Name: %v\n", id, peripheral.Name())
		}
	}
}
