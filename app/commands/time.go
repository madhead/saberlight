package commands

import (
	"time"

	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/util"
)

// Time sets device's time
func Time() {
	now := time.Unix(*cli.TimeTime, 0)
	dayOfWeek := now.Weekday()

	if dayOfWeek == 0 {
		dayOfWeek = 7
	}

	util.Write(*cli.OnTarget, "FFD5", "FFD9", []byte{
		0x10,
		byte(now.Year() / 100),
		byte(now.Year() % 100),
		byte(now.Month()),
		byte(now.Day()),
		byte(now.Hour()),
		byte(now.Minute()),
		byte(now.Second()),
		byte(dayOfWeek),
		0x00,
		0x01,
	})
}
