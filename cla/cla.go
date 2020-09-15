/*
MIT License

Copyright (c) 2020 Lee Choon Siong

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
