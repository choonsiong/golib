package cla

import (
	"fmt"
	"os"
)

// GetFileName returns a filename from the command line argument.
func GetFileName() string {
	if len(os.Args) != 2 {
		fmt.Printf("Invalid argument, please provide a filename")
		os.Exit(22)
	}

	return os.Args[1]
}

// GetFileNames returns number of filenames specified in count.
func GetFileNames(count int) []string {
	total := count + 1

	if len(os.Args) != total {
		fmt.Printf("Invalid arguments, please provide %d filenames", count)
		os.Exit(22)
	}

	return os.Args[1:]
}
