// Package cli provides helpers to work with command-line interface.
package cli

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
