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

// Package sms provides method for sending sms text message using the
// NGM's AGW server.
package sms

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Sms struct {
	Host       string // hostname or ip of AGW server
	Port       string // AGW listening port
	Sender     string // sms sender
	Text       string // sms text
	Recipients []string // msisdn list in international format
}

// Send sends a sms text message using the content of Sms struct.
// Note: The caller is required to initialize Sms struct with correct values for all fields.
func (s Sms) Send() (string, error) {
	// Normalized sms text
	t := strings.Replace(s.Text, " ", "+", -1) // replace whitespace with '+'
	t = strings.Replace(t, "\n", "+%A+", -1) // replace newline with '%A'

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	for _, r := range s.Recipients {
		content := "http://" + s.Host + ":" + s.Port + "/send?sms_dest=" + r + "&sms_source=" + s.Sender + "&sms_valid_rel=500&sms_text=" + t + " HTTP/1.0"

		cmd := exec.Command("curl", content)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			fmt.Fprintf(os.Stderr, stderr.String())
			return stdout.String(), err
		}
	}

	return stdout.String(), nil
}