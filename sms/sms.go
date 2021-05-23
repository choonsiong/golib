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

// Package sms provides methods for sending SMS text message.
package sms

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

// Sms contains all the required fields for sending SMS text message
// via the AGW server. Note that the caller is required to initialize
// the Sms type with all the required fields.
type Sms struct {
	Host       string // AGW server hostname/ip
	Port       string // AGW listening port
	Sender     string // SMS sender
	Text       string // SMS text
	Recipients []string // MSISDN list in international format
}

// SmsRecipient contains the information related to a SMS recipient.
type SmsRecipient struct {
	MSISDN string
}

var wg sync.WaitGroup

// Process process the sms send request.
func (s Sms) Process() {
	smsRecipient := make(chan SmsRecipient)
	smsStatus := make(chan SmsRecipient)

	go s.processSmsRecipient(smsRecipient)
	go s.sendSms(smsRecipient, smsStatus)

	s.printStatus(smsStatus)

	wg.Wait()
}

// printStatus prints the status of the sms sending request.
// Note that due to how the AGW works, the program will not know
// the delivery status of the SMS, therefore, caller should not
// assume the status here means the SMS is delivered successfully
// to the recipients.
func (s Sms) printStatus(in <-chan SmsRecipient) {
	var count int

	for smsRecipient := range in {
		io.WriteString(os.Stdout, smsRecipient.MSISDN + "\n")
		count++
	}

	fmt.Fprintf(os.Stdout, "Processed %d sms request.\n", count)
}

// processSmsRecipient verify and process all the sms recipients.
func (s Sms) processSmsRecipient(out chan<- SmsRecipient) {
	// Make sure MSISDN is valid, i.e. 601xxxxxxxx
	grammar := "(601[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]?)"
	match := regexp.MustCompile(grammar)

	for _, msisdn := range s.Recipients {
		var r SmsRecipient

		if match.MatchString(msisdn) {
			r.MSISDN = msisdn
		} else {
			fmt.Fprintf(os.Stderr, "sms.processSmsRecipient: invalid msisdn: %v\n", msisdn)
			continue
		}

		out<- r
	}

	close(out)
}

// sendSms send the sms text message to the sms recipient.
func (s Sms) sendSms(in <-chan SmsRecipient, out chan<- SmsRecipient) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	for smsRecipient := range in {
		wg.Add(1)

		go func() {
			smsContent := "http://" + s.Host + ":" + s.Port + "/send?sms_dest=" + smsRecipient.MSISDN + "&sms_source=" + s.Sender + "&sms_valid_rel=500&sms_text=" + normalizedSmsText(s.Text) + " HTTP/1.0"

			cmd := exec.Command("curl", smsContent)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()

			if err != nil {
				fmt.Fprintf(os.Stderr, "sms.sendSms: %v: %v\n", err, stderr.String())
			}
			
			wg.Done()
		}()

		out<- smsRecipient
	}

	close(out)
}

// normalizedSmsText normalized the SMS text received from user.
// The AGW server expected the SMS text in certain format, i.e. white space with '+'
// and newline with '%A'.
func normalizedSmsText(text string) string {
	return strings.Replace(strings.Replace(text, " ", "+", -1), "\n", "+%A+", -1)
}

// Deprecated: 1.0.18
// Send sends a sms text message.
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