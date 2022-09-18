package sms

import (
	"errors"
	"github.com/choonsiong/golib/v2/logger"
	"github.com/choonsiong/golib/v2/logger/commonlog"
	"github.com/choonsiong/golib/v2/logger/jsonlog"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	l := jsonlog.New(nil, logger.LevelDebug)

	s1 := New(l)
	s2 := &SMS{
		Logger: l,
	}

	if !reflect.DeepEqual(s1, s2) {
		t.Errorf("want %v; got %v", s1, s2)
	}
}

func TestSMS_SendSMS(t *testing.T) {
	sms := SMS{
		Host:       "127.0.0.1",
		Port:       "8000",
		Sender:     "12345",
		Content:    "HELLO",
		Recipients: []string{"222", "333", "444"},
		UseHTTPS:   false,
		Logger:     commonlog.New(os.Stdout, logger.LevelError),
	}

	err := sms.SendSMS()
	if err != nil {
		if !errors.Is(err, ErrRunningCurl) {
			t.Errorf("SendSMS() == %v; want %v", err, ErrRunningCurl)
		}
	}
}

func TestSMS_SendSMSWithHTTPS(t *testing.T) {
	sms := SMS{
		Host:       "127.0.0.1",
		Port:       "8000",
		Sender:     "12345",
		Content:    "HELLO",
		Recipients: []string{"222", "333", "444"},
		UseHTTPS:   true,
		Logger:     commonlog.New(os.Stdout, logger.LevelError),
	}

	err := sms.SendSMS()
	if err != nil {
		if !errors.Is(err, ErrRunningCurl) {
			t.Errorf("SendSMS() == %v; want %v", err, ErrRunningCurl)
		}
	}
}
