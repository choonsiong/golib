// Package cli provides helpers to work with command-line interface.
package cli

import (
	"bufio"
	"fmt"
	"os"
	"path"
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

// GetStringsWithPrompt prompts text to the user and accept number of size
// strings from standard input, and returns a slice of string.
// Use case: You want to read some fixed number of strings from stdin, but
// you don't want to declare any variables for those strings, instead this
// function returns a slice of string which you can then further process
// accordingly.
func GetStringsWithPrompt(text string, size int) (result []string) {
	result = make([]string, size)
	vals := make([]interface{}, size)

	for i := 0; i < len(result); i++ {
		vals[i] = &result[i]
	}

	fmt.Print(text + ": ")
	fmt.Scan(vals...) // we cannot use vals... here

	return
}

// ProgName returns the program name from the command-line arguments.
func ProgName(args []string) string {
	if len(args) == 0 {
		return ""
	}

	return path.Base(args[0])
}
