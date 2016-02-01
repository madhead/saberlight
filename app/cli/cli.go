package cli

import (
	"github.com/alecthomas/kingpin"
)

var (
	App           = kingpin.New("saberlight", `A tool to control some "smart" BLE bulbs.`)
	DeviceTimeout = App.Flag("deviceTimeout", "Maximum wait time for HCI to start").Default("3s").Duration()
	ScanPeriod    = App.Flag("scanPeriod", "Scan period").Default("5s").Duration()
	Scan          = App.Command("scan", "Scan for BLE devices")
	Dump          = App.Command("dump", "Dump BLE device")
	DumpTarget    = Dump.Arg("id", "Device to dump").Required().String()
	Server        = App.Command("server", "Start REST server")
	Version       = App.Command("version", "Print app's version")
)
