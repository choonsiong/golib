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

package sms

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type SmsContent struct {
	Host       string // hostname or ip
	Port       string
	Sender     string   // to identify the sms sender
	Text       string   // sms text
	Recipients []string // msisdn list of recipients
}

// SendSMS sends a text message using the content of SmsContent struct.
// The caller is required to initialize SmsContent with correct values for all fields.
// Furthermore, a preconfigured SMS server is required.
func SendSMS(sc SmsContent) {
	// Normalized sms text
	t := strings.Replace(sc.Text, " ", "+", -1) // replace whitespace with '+'

	for _, r := range sc.Recipients {
		content := "http://" + sc.Host + ":" + sc.Port + "/send?sms_dest=" + r + "&sms_source=" + sc.Sender + "&sms_valid_rel=500&sms_text=" + t + " HTTP/1.0"

		cmd := exec.Command("curl", content)

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			msg := fmt.Sprint(err) + ": " + stderr.String()
			fmt.Fprintf(os.Stderr, "%v\n", msg)
		}
	}
}
