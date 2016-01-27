package util

import (
	"errors"
	"time"

	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/util/log"
	"github.com/paypal/gatt"
)

type deviceResult struct {
	device gatt.Device
	err    error
}

// OpenHCI opens HCI device for future interactions
func OpenHCI() (gatt.Device, error) {
	// Sometimes gatt.NewDevice blocks indefinitely:
	// h, err := linux.NewHCI(d.devID, d.chkLE, d.maxConn)
	// h.resetDevice()
	// if err := h.c.SendAndCheckResp(s, []byte{0x00}); err != nil {
	// ...due to device doesn't respond to a command (e.g. opReset = hostCtl<<10 | 0x0003 // Reset)
	// so wrap this in gorutine and timeout it

	deviceReady := make(chan *deviceResult)

	// TODO: when running webserver, this gouroutine can run forever, causing leaks!
	go func() {
		device, err := gatt.NewDevice([]gatt.Option{
			gatt.LnxDeviceID(-1, true),
			gatt.LnxMaxConnections(1),
		}...)

		if err != nil {
			log.Error.Printf("Failed to open device: %v\n", err)
			deviceReady <- &deviceResult{
				device: nil,
				err:    err,
			}
			return
		}

		device.Init(func(device gatt.Device, state gatt.State) {
			if state == gatt.StatePoweredOn {
				deviceReady <- &deviceResult{
					device: device,
					err:    nil,
				}
			}
		})
	}()

	select {
	case <-time.After(*cli.DeviceTimeout):
		log.Error.Println("HCI device timed out")
		return nil, errors.New("HCI device timed out")
	case device := <-deviceReady:
		return device.device, device.err
	}
}
