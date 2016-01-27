package log

import (
	"log"
	"os"
)

var (
	// Error logs to stderr
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.LUTC|log.Ltime)

	// Info logs to stdout
	Info = log.New(os.Stdout, "INFO:  ", log.Ldate|log.LUTC|log.Ltime)
)
