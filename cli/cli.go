// Package cli provides helpers to work with command-line interface.
package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Filename returns the value of args[1] as the required filename, os.Args
// is required to pass in args.
func Filename(args []string) (string, error) {
	if len(args) > 2 {
		return "", ErrTooManyArguments
	}

	if len(args) != 2 {
		return "", ErrInsufficientArguments
	}

	if args[1] == "" {
		return "", ErrEmptyFilename
	}

	return args[1], nil
}

// GetFloat reads a floating-point number from os.Stdin. It returns
// the number read and any error encountered.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrReadString, err)
	}

	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrParseFloat, err)
	}

	return number, nil
}
