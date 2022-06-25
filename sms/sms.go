// Package sms provides helpers to work with sms.
package sms

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// SMS implements various information needed for the SMS type.
type SMS struct {
	Host       string   // sms server hostname or ip
	Port       string   // sms server port number
	Sender     string   // sms sender
	Content    string   // sms content
	Recipients []string // list of sms recipients
	UseHTTPS   bool     // use https?
}

// SendSMS sends sms text message via the configured SMS gateway.
func (s *SMS) SendSMS() error {
	str := strings.Replace(s.Content, " ", "+", -1) // Replace whitespace with '+'

	proto := "http"

	if s.UseHTTPS {
		proto = "https"
	}

	for _, r := range s.Recipients {
		message := proto + "://" + s.Host + ":" + s.Port + "/send?sms_dest=" + r + "&sms_source=" + s.Sender + "&sms_valid_rel=500&sms_text=" + str + " HTTP/1.0"

		cmd := exec.Command("curl", message)

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			return errors.New(fmt.Sprintf("SendSMS(): %v: %v", err, stderr.String()))
		}
	}

	return nil
}
