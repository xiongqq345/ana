package cmd

import (
	"fmt"
	"os"
)

const (
	// http://tldp.org/LDP/abs/html/exitcodes.html
	ExitSuccess = iota
	ExitError
	ExitBadConnection
	ExitInvalidInput // for txn, watch command
	ExitBadFeature   // provided a valid flag with an unsupported value
	ExitInterrupted
	ExitIO
	ExitBadArgs = 128
)

func ExitWithError(code int, err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	
	//if cerr, ok := err.(*client.ClusterError); ok {
	//	fmt.Fprintln(os.Stderr, cerr.Detail())
	//}
	os.Exit(code)
}

