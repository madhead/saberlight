package cli

import (
	"github.com/alecthomas/kingpin"
)

var (
	App           = kingpin.New("saberlight", `A tool to control some "smart" BLE bulbs.`)
	DeviceTimeout = App.Flag("deviceTimeout", "Maximum wait time for HCI to start").Default("3s").Duration()
	ScanPeriod    = App.Flag("scanPeriod", "Scan period").Default("5s").Duration()
	Version       = App.Command("version", "Print app's version")
	Scan          = App.Command("scan", "Scan for BLE devices")
	Dump          = App.Command("dump", "Dump BLE device")
	DumpTarget    = Dump.Arg("id", "Device to dump").Required().String()
	SetColor      = App.Command("color", "Dump BLE device")
	ColorTarget   = SetColor.Arg("id", "Target device").Required().String()
	ColorRed      = SetColor.Arg("red", "Red color component").Required().Uint8()
	ColorGreen    = SetColor.Arg("green", "Green color component").Required().Uint8()
	ColorBlue     = SetColor.Arg("blue", "Blue color component").Required().Uint8()
)
