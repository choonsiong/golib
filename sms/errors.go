package sms

import "errors"

var (
	ErrRunningCurl = errors.New("sms: error running curl command")
)
