// Package sms provides types and methods to work with SMS delivery.
package sms

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/choonsiong/golib/v2/logger"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// SMS contains properties required for sending SMS to an SMS gateway.
type SMS struct {
	UseHTTPS   bool     // Use secure http
	Content    string   // SMS content
	Host       string   // SMS server hostname or ip
	Port       string   // SMS server port number
	Sender     string   // SMS sender
	Recipients []string // List of sms recipients
	Logger     logger.Logger
}

func New(l logger.Logger) *SMS {
	return &SMS{
		Logger: l,
	}
}

// SendNewSMS sends sms text message via the configured SMS gateway.
func (s *SMS) SendNewSMS() {
	httpClient := http.Client{
		Timeout: time.Second * 3,
	}

	smsContent := strings.Replace(s.Content, " ", "+", -1) // Replace whitespace with '+'

	proto := "http"

	if s.UseHTTPS {
		proto = "https"
	}

	s.Logger.PrintDebug("SMS.SendNewSMS()", map[string]string{
		"s": fmt.Sprintf("%+v", s),
	})

	for _, r := range s.Recipients {
		message := proto + "://" + s.Host + ":" + s.Port + "/send?sms_dest=" + r + "&sms_source=" + s.Sender + "&sms_valid_rel=500&sms_text=" + smsContent + " HTTP/1.0"

		s.Logger.PrintDebug("SMS.SendNewSMS()", map[string]string{
			"message": message,
		})

		req, err := http.NewRequest(http.MethodGet, message, nil)
		if err != nil {
			s.Logger.PrintWarning("SMS.SendNewSMS()", map[string]string{
				"message": message,
				"err":     err.Error(),
			})
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			s.Logger.PrintWarning("SMS.SendNewSMS()", map[string]string{
				"message": message,
				"err":     err.Error(),
			})
		}

		if strings.ToLower(resp.Status) != "200 ok" {
			s.Logger.PrintWarning("SMS.SendNewSMS()", map[string]string{
				"message":     message,
				"resp.Status": resp.Status,
			})
		}

		bufferedReader := bufio.NewReader(resp.Body)

		var sb strings.Builder
		slice := make([]byte, 16)

		for {
			count, err := bufferedReader.Read(slice)
			if count > 0 {
				sb.Write(slice[0:count])
			}
			if err != nil {
				break
			}
		}

		if sb.String() != "" {
			s.Logger.PrintDebug("SMS.SendNewSMS()", map[string]string{
				"response": sb.String(),
			})
		}

		resp.Body.Close()
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
