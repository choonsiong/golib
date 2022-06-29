package sms

import (
	"testing"
)

func TestSMS_SendSMS(t *testing.T) {
	sms := SMS{
		Host:       "127.0.0.1",
		Port:       "8000",
		Sender:     "12345",
		Content:    "HELLO",
		Recipients: []string{"222", "333", "444"},
		UseHTTPS:   false,
	}
	sms.SendSMS()
}
