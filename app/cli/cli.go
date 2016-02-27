package cli

import (
	"github.com/alecthomas/kingpin"
)

var (
	App            = kingpin.New("saberlight", `A tool to control Triones "smart" BLE bulbs.`)
	DeviceTimeout  = App.Flag("deviceTimeout", "Maximum wait time for HCI to start").Default("3s").Duration()
	ScanPeriod     = App.Flag("scanPeriod", "Scan period").Default("5s").Duration()
	Version        = App.Command("version", "Print app's version")
	Scan           = App.Command("scan", "Scan for bulbs")
	On             = App.Command("on", "Power bulb on")
	OnTarget       = On.Arg("id", "Target bulb").Required().String()
	Off            = App.Command("off", "Power bulb off")
	OffTarget      = Off.Arg("id", "Target bulb").Required().String()
	Dump           = App.Command("dump", "Dump bulb")
	DumpTarget     = Dump.Arg("id", "Bulb to dump").Required().String()
	Color          = App.Command("color", "Set static color")
	ColorTarget    = Color.Arg("id", "Target bulb").Required().String()
	ColorRed       = Color.Arg("red", "Red color component").Required().Uint8()
	ColorGreen     = Color.Arg("green", "Green color component").Required().Uint8()
	ColorBlue      = Color.Arg("blue", "Blue color component").Required().Uint8()
	White          = App.Command("white", "Set white color with given intensity")
	WhiteTarget    = White.Arg("id", "Target bulb").Required().String()
	WhiteIntensity = White.Arg("intensity", "White color intensity").Required().Uint8()
	Mode           = App.Command("mode", "Set predefined mode")
	ModeTarget     = Mode.Arg("id", "Target bulb").Required().String()
	ModeMode       = Mode.Arg("mode", `Mode number:
		0x25: Seven color cross fade
		0x26: Red gradual change
		0x27: Green gradual change
		0x28: Blue gradual change
		0x29: Yellow gradual change
		0x2A: Cyan gradual change
		0x2B: Purple gradual change
		0x2C: White gradual change
		0x2D: Red,Green cross fade
		0x2E: Red,Blue cross fade
		0x2F: Green,Blue cross fade
		0x30: Seven color stobe flash
		0x31: Red strobe flash
		0x32: Green strobe flash
		0x33: Blue strobe flash
		0x34: Yellow strobe flash
		0x35: Cyan strobe flash
		0x36: Purple strobe flash
		0x37: White strobe flash
		0x38: Seven color jumping change
	`).Required().Uint8()
	ModeSpeed = Mode.Arg("speed", `Mode speed:
		0x01: The fastests
		...
		0xFF: The slowest
	`).Required().Uint8()
	Time          = App.Command("time", "Set bulb's time")
	TimeTarget    = Time.Arg("id", "Target bulb").Required().String()
	TimeTime      = Time.Arg("time", "Time to set in seconds since 1970-01-01 00:00:00 UTC").Required().Int64()
	Timings       = App.Command("timings", "Get active timings")
	TimingsTarget = Timings.Arg("id", "Target bulb").Required().String()
)
