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
	Host       string
	Port       string
	Sender     string
	Text       string
	Recipients []string
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
			io.WriteString(os.Stderr, msg)
			io.WriteString(os.Stderr, "\n")
		}
	}
}
