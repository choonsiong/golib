package cla

import (
	"fmt"
	"os"
)

// GetFileName returns a file name from the command line argument.
func GetFileName() string {
	if len(os.Args) != 2 {
		fmt.Printf("Invalid argument, please provide a filename")
		os.Exit(22)
	}

	return os.Args[1]
}
