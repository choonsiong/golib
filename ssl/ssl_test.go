package ssl

import (
	"github.com/choonsiong/golib/logger/jsonlog"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	logger := &jsonlog.Logger{}
	want := &SSL{
		Logger: logger,
	}
	got := New(logger)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v; got %v", got, want)
	}
}

func TestSSL_Generate(t *testing.T) {
	ssl := SSL{
		CertPath:     "/tmp/ssl",
		KeyPath:      "/tmp/ssl",
		CertName:     "cert.pem",
		KeyName:      "key.pem",
		Country:      "MY",
		State:        "Kuala Lumpur",
		Location:     "Kuala Lumpur",
		Organization: "Foo Bar",
		CommonName:   "com.example",
		Days:         "30",
		Logger:       jsonlog.New(os.Stdout, jsonlog.LevelError),
	}

	err := ssl.Generate()
	if err != nil {
		t.Errorf("SSL.Generate(): want error nil; got %v", err)
	}

	fileInfo, err := os.Stat("/tmp/ssl/cert.pem")
	if err != nil {
		t.Errorf("SSL.Generate(): file %s not found", "/tmp/ssl/cert.pem")
	}
	if fileInfo.Mode().String() != "-rw-------" {
		t.Errorf("SSL.Generate(): want %s; got %s", "-rw-------", fileInfo.Mode().String())
	}

	fileInfo, err = os.Stat("/tmp/ssl/key.pem")
	if err != nil {
		t.Errorf("SSL.Generate(): file %s not found", "/tmp/ssl/key.pem")
	}
	if fileInfo.Mode().String() != "-rw-r--r--" {
		t.Errorf("SSL.Generate(): want %s; got %s", "-rw-r--r--", fileInfo.Mode().String())
	}
}
