package sms

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type SmsContent struct {
	host       string
	port       string
	sender     string
	text       string
	recipients []string
}

// SendSMS sends a text message using the content of SmsContent struct.
// Preconfigured SMS server is required.
func SendSMS(sc SmsContent) {
	// Normalize sms text
	t := strings.Replace(sc.text, " ", "+", -1) // replace space with '+'

	for _, r := range sc.recipients {
		content := "http://" + sc.host + ":" + sc.port + "/send?sms_dest=" + r + "&sms_source=" + sc.sender + "&sms_valid_rel=500&sms_text=" + t + " HTTP/1.0"

		cmd := exec.Command("curl", content)

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			msg := fmt.Sprint(err) + ": " + stderr.String()
			io.WriteString(os.Stderr, msg)
			io.WriteString(os.Stderr, "\n")
		}
	}
}
