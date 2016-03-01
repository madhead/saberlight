package commands

import (
	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/util"
)

// On turns bulb on
func On() {
	util.Write(*cli.OnTarget, "FFD5", "FFD9", []byte{0xCC, 0x23, 0x33})
}

// Off turns bulb off
func Off() {
	util.Write(*cli.OffTarget, "FFD5", "FFD9", []byte{0xCC, 0x24, 0x33})
}
