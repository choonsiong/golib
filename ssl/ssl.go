// Package ssl provides helpers to work with SSL.
package ssl

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/choonsiong/golib/logger/jsonlog"
	"os"
	"os/exec"
)

type SSL struct {
	CertName     string
	CertPath     string
	KeyName      string
	KeyPath      string
	Country      string
	State        string
	Location     string
	Organization string
	CommonName   string
	Days         string
	Logger       *jsonlog.Logger
}

// New returns a new SSL.
func New(logger *jsonlog.Logger) *SSL {
	return &SSL{
		Logger: logger,
	}
}

// Generate generates the SSL/TLS certificate and private key.
func (s *SSL) Generate() error {
	if _, err := os.Stat(s.CertPath); err != nil {
		err = os.Mkdir(s.CertPath, 0755)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(s.KeyPath); err != nil {
		err = os.Mkdir(s.KeyPath, 0755)
		if err != nil {
			return err
		}
	}

	s.Logger.PrintDebug("SSL.Generate()", map[string]string{
		"CertPath": s.CertPath,
		"KeyPath":  s.KeyPath,
	})

	//if _, err := file.Exists("openssl"); err != nil {
	//	return errors.New(fmt.Sprintf("SSL.Generate(): %v", err))
	//}
	
	if _, err := exec.LookPath("openssl"); err != nil {
		return errors.New(fmt.Sprintf("SSL.Generate(): %v", err))
	}

	// Generate SSL cert and private key file using openssl command
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	certPath := s.CertPath + "/" + s.CertName
	keyPath := s.KeyPath + "/" + s.KeyName
	subject := "/C=" + s.Country + "/ST=" + s.State + "/L=" + s.Location + "/O=" + s.Organization + "/CN=" + s.CommonName

	cmd := exec.Command("openssl", "genrsa", "-out", keyPath, "2048")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("openssl", "req", "-x509", "-new", "-nodes", "-key", keyPath, "-sha256", "-days", s.Days, "-subj", subject, "-out", certPath)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	// Make sure the cert and key files has correct permission
	err = os.Chmod(certPath, 0600)
	if err != nil {
		return err
	}

	err = os.Chmod(keyPath, 0644)
	if err != nil {
		return err
	}

	return nil
}
