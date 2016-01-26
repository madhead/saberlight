package commands

import (
	"fmt"
)

var version = "dev"

func Version() {
	fmt.Println(version)
}
