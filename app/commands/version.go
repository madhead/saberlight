package commands

import (
	"fmt"
)

var version = "dev"

// Version prints application version
func Version() {
	fmt.Println(version)
}
