package commands

import (
	"os"

	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/util"
	"github.com/madhead/saberlight/app/util/log"
)

// Mode sets mode
func Mode() {
	if (*cli.ModeMode < 0x25) || (*cli.ModeMode > 0x38) {
		log.Error.Println("Unknown mode")
		os.Exit(util.ExitStatusGenericError)
	}

	util.Write(*cli.ModeTarget, "FFD5", "FFD9", []byte{0xBB, *cli.ModeMode, *cli.ModeSpeed, 0x44})
}
