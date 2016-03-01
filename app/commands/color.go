package commands

import (
	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/util"
)

var zero = byte(0)

// Color sets color
func Color() {
	util.Write(*cli.ColorTarget, "FFD5", "FFD9", []byte{0x56, *cli.ColorRed, *cli.ColorGreen, *cli.ColorBlue, 0x00, 0xF0, 0xAA})
}

// White sets white color with given intensity
func White() {
	util.Write(*cli.WhiteTarget, "FFD5", "FFD9", []byte{0x56, 0, 0, 0, *cli.WhiteIntensity, 0x0F, 0xAA})
}
