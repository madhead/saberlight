package commands

import (
	"fmt"
	"net/http"
	"time"

	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/util"

	"github.com/paypal/gatt"
)

// Server serves!
func Server() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		device, err := util.OpenHCI()

		if err != nil {
			response := &http.Response{}

			response.StatusCode = 503
			response.Write(responseWriter)
			return
		}

		peripherals := make(map[string]gatt.Peripheral)

		device.Handle(gatt.PeripheralDiscovered(func(peripheral gatt.Peripheral, advertisement *gatt.Advertisement, rssi int) {
			peripherals[peripheral.ID()] = peripheral
		}))

		device.Scan([]gatt.UUID{}, false)
		time.Sleep(*cli.ScanPeriod)
		device.StopScanning()

		for id, peripheral := range peripherals {
			responseWriter.Write([]byte(fmt.Sprintf("ID: %s, Name: %v\n", id, peripheral.Name())))
		}

		device.Stop()
	})
	http.ListenAndServe(":8080", nil)
}
