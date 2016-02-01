package main

import (
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/madhead/saberlight/app/cli"
	"github.com/madhead/saberlight/app/commands"
)

func main() {
	switch kingpin.MustParse(cli.App.Parse(os.Args[1:])) {
	case cli.Version.FullCommand():
		commands.Version()
	case cli.Scan.FullCommand():
		commands.Scan()
	case cli.Dump.FullCommand():
		commands.Dump()
	case cli.SetColor.FullCommand():
		commands.SetColor()
	}
}
