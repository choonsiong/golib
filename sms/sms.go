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
// NGM AGW server.
package sms

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Sms type provides a simplified way for the user to define all the
// information required to send SMS via AGW server. Furthermore, user
// can enter normal text message (e.g. with spaces and newline characters)
// without worrying about formatting it correctly.
type Sms struct {
	Host       string // hostname or ip of AGW server
	Port       string // AGW listening port
	Sender     string // sms sender
	Text       string // sms text
	Recipients []string // msisdn list in international format
}

type SmsRecipient struct {
	MSISDN string
}

// Public Methods

func (s Sms) Process() {
	smsRecipient := make(chan SmsRecipient)
	smsCount := make(chan int)
	go s.processSmsRecipient(smsRecipient)
	go s.sendSMS(smsRecipient, smsCount)

}

// Private Methods

// normalizedSmsText normalized the SMS text from user input.
func normalizedSmsText(text string) string {
	s := strings.Replace(text, " ", "+", -1) // replace whitespace with '+'
	s = strings.Replace(s, "\n", "+%A+", -1) // replace newline with '%A'

	return s
}

func (s Sms) checkCount(in <-chan int) {
	var runningCount int

	for i := range in {
		runningCount = i
	}

	if runningCount != len(s.Recipients) {
		fmt.Fprintf(os.Stdout, "Failed to send to all recipients: %i/%i\n", runningCount, len(s.Recipients))
	} else {
		fmt.Fprintf(os.Stdout, "SMS sent to all recipients successfully.\n")
	}
}

func (s Sms) processSmsRecipient(out chan<- SmsRecipient) {
	for _, v := range s.Recipients {
		var r SmsRecipient
		r.MSISDN = v
		out<- r
	}

	close(out)
}

func (s Sms) sendSMS(in <-chan SmsRecipient, out chan<- int) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var count int

	for smsRecipient := range in {
		smsContent := "http://" + s.Host + ":" + s.Port + "/send?sms_dest=" + smsRecipient.MSISDN + "&sms_source=" + s.Sender + "&sms_valid_rel=500&sms_text=" + normalizedSmsText(s.Text) + " HTTP/1.0"

		cmd := exec.Command("curl", smsContent)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			fmt.Fprintf(os.Stderr, "sms.sendSMS: %v: %v\n", err, stderr.String())
			continue
		}

		count += 1
		out<- count
	}

	close(out)
}

// Send sends a sms text message using the value of Sms struct.
// Note: The caller is required to initialize Sms struct with correct values for all fields.
func (s Sms) Send() {
	// Normalized sms text (AGW expect the text in certain format)
	smsText := strings.Replace(s.Text, " ", "+", -1) // replace whitespace with '+'
	smsText = strings.Replace(smsText, "\n", "+%A+", -1) // replace newline with '%A'

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	for _, r := range s.Recipients {
		content := "http://" + s.Host + ":" + s.Port + "/send?sms_dest=" + r + "&sms_source=" + s.Sender + "&sms_valid_rel=500&sms_text=" + smsText + " HTTP/1.0"

		cmd := exec.Command("curl", content)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			fmt.Fprintf(os.Stderr, "sms.Send: %v: %v\n", err, stderr.String())
			continue
		}
	}
}