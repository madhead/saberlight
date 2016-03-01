package util

const (
	// ExitStatusGenericError signals about generic error
	ExitStatusGenericError = 1
	// ExitStatusHCIError signals about errors with HCI device
	ExitStatusHCIError = 2
	// ExitStatusTimeout signals that operation failed to complete in time
	ExitStatusTimeout = 3
)
