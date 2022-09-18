// Package sms provides helpers to work with sms.
package sms

import (
	"bytes"
	"fmt"
	"github.com/choonsiong/golib/v2/logger"
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
	Logger     logger.Logger
}

func New(l logger.Logger) *SMS {
	return &SMS{
		Logger: l,
	}
}

// SendSMS sends sms text message via the configured SMS gateway.
func (s *SMS) SendSMS() error {
	str := strings.Replace(s.Content, " ", "+", -1) // Replace whitespace with '+'

	proto := "http"

	if s.UseHTTPS {
		proto = "https"
	}

	s.Logger.PrintDebug("SMS.SendSMS()", map[string]string{
		"s": fmt.Sprintf("%v", s),
	})

	for _, r := range s.Recipients {
		message := proto + "://" + s.Host + ":" + s.Port + "/send?sms_dest=" + r + "&sms_source=" + s.Sender + "&sms_valid_rel=500&sms_text=" + str + " HTTP/1.0"

		s.Logger.PrintDebug("SMS.SendSMS()", map[string]string{
			"message": message,
		})

		cmd := exec.Command("curl", message)

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			return fmt.Errorf("SendSMS(): %w: %v: %v", ErrRunningCurl, err, stderr.String())
		}
	}

	return nil
}
