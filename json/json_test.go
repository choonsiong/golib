package json

import (
	"bytes"
	"errors"
	"github.com/choonsiong/golib/logger"
	"github.com/choonsiong/golib/logger/commonlog"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func TestJSON_ReadJSON(t *testing.T) {
	j := &JSON{}
	j.Logger = commonlog.New(os.Stdout, logger.LevelError)

	tests := []struct {
		name               string
		body               string
		maxBytes           int
		allowUnknownFields bool
		want               *Person
		wantErr            error
	}{
		{
			name:    "valid json body",
			body:    `{"name": "foobar","age": 42,"email":"foobar@example.com"}`,
			want:    &Person{"foobar", 42, "foobar@example.com"},
			wantErr: nil,
		},
		{
			name:    "valid json body with extra field",
			body:    `{"name": "foobar","age": 42,"email":"foobar@example.com","mobile":"1234567"}`,
			want:    &Person{"foobar", 42, "foobar@example.com"},
			wantErr: nil,
		},
		{
			name:    "badly formatted json",
			body:    `{"name":}`,
			want:    new(Person),
			wantErr: nil,
		},
		{
			name:    "multiple json body",
			body:    `{"name": "foobar","age": 42,"email":"foobar@example.com"}{"name": "alice","age": 28,"email":"alice@example.com"}`,
			want:    &Person{"foobar", 42, "foobar@example.com"},
			wantErr: ErrMultipleJSONValue,
		},
		{
			name:     "maximum bytes",
			body:     `{"name": "foobar","age": 42,"email":"foobar@example.com"}`,
			maxBytes: 1,
			want:     new(Person),
			wantErr:  ErrDecodeJSON,
		},
		{
			name:               "disallow unknown fields",
			body:               `{"name": "foobar","age": 42,"email":"foobar@example.com","unknown":"unknown"}`,
			want:               &Person{"foobar", 42, "foobar@example.com"},
			allowUnknownFields: false,
			wantErr:            ErrDecodeJSON,
		},
		{
			name:               "allow unknown fields",
			body:               `{"name": "foobar","age": 42,"email":"foobar@example.com","unknown":"unknown"}`,
			want:               &Person{"foobar", 42, "foobar@example.com"},
			allowUnknownFields: true,
			wantErr:            nil,
		},
		{
			name:    "empty json body",
			body:    ``,
			want:    new(Person),
			wantErr: ErrDecodeJSON,
		},
		{
			name:    "incorrect json field type",
			body:    `{"name": 42}`,
			want:    new(Person),
			wantErr: ErrDecodeJSON,
		},
		{
			name:    "syntax error in json body",
			body:    `{"name": "bar"`,
			want:    new(Person),
			wantErr: ErrDecodeJSON,
		},
		{
			name:    "missing field name in json body",
			body:    `{name: "fobar"}`,
			want:    new(Person),
			wantErr: ErrDecodeJSON,
		},
		{
			name:    "not json",
			body:    `hello world`,
			want:    new(Person),
			wantErr: ErrDecodeJSON,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			p := new(Person)

			r, err := http.NewRequest(http.MethodGet, "/", strings.NewReader(tt.body))
			r.Header.Set("Content-Type", "application/json")

			if err != nil {
				t.Fatal(err)
			}

			j.AllowUnknownFields = tt.allowUnknownFields
			j.MaxBytes = tt.maxBytes

			err = j.ReadJSON(w, r, p)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("JSON.ReadJSON(): want error %v; got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("JSON.ReadJSON(): want error %v; got %v", tt.wantErr, err)
				}
			}

			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("JSON.ReadJSON(): want %v; got %v", tt.want, p)
			}

			p = nil
		})
	}
}

func TestJSON_ReadJSONDecode(t *testing.T) {
	j := &JSON{}
	j.Logger = commonlog.New(os.Stdout, logger.LevelError)

	w := httptest.NewRecorder()

	tests := []struct {
		name    string
		body    string
		want    *Person
		wantErr error
	}{
		{
			name:    "nil receiver",
			body:    `{"name": "foobar","age": 42,"email":"foobar@example.com"}`,
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodGet, "/", bytes.NewReader([]byte(tt.body)))
			r.Header.Set("Content-Type", "application/json")

			if err != nil {
				t.Fatal(err)
			}

			err = j.ReadJSON(w, r, nil)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("JSON.ReadJSON(): want error %v; got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("JSON.ReadJSON(): want error %v; got %v", tt.wantErr, err)
				}
			}
		})
	}
}

func TestJSON_WriteJSON(t *testing.T) {
	j := &JSON{}
	j.Logger = commonlog.New(os.Stdout, logger.LevelError)

	h := http.Header{}
	h["testing"] = []string{"header testing"}

	tests := []struct {
		name    string
		status  int
		data    any
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "valid struct",
			status:  http.StatusOK,
			data:    &Person{"foobar", 42, "foobar@example.com"},
			headers: nil,
			want:    "{\n\t\"name\": \"foobar\",\n\t\"age\": 42,\n\t\"email\": \"foobar@example.com\"\n}",
			wantErr: nil,
		},
		{
			name:    "integer data",
			status:  http.StatusOK,
			data:    1,
			headers: nil,
			want:    "1",
			wantErr: nil,
		},
		{
			name:   "valid struct with missing field",
			status: http.StatusOK,
			data: &Person{
				Name: "foobar",
				Age:  42,
			},
			headers: nil,
			want:    "{\n\t\"name\": \"foobar\",\n\t\"age\": 42,\n\t\"email\": \"\"\n}",
			wantErr: nil,
		},
		{
			name:    "nil data",
			status:  http.StatusOK,
			data:    nil,
			headers: nil,
			want:    "null",
			wantErr: nil,
		},
		{
			name:    "with header",
			status:  http.StatusOK,
			data:    nil,
			headers: h,
			want:    "null",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			_, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			err = j.WriteJSON(w, tt.status, tt.data, tt.headers)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("JSON.WriteJSON(): want error %v; got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("JSON.WriteJSON(): want error %v; got %v", tt.wantErr, err)
				}
			}

			got := w.Result()

			if got.StatusCode != http.StatusOK {
				t.Errorf("JSON.WriteJSON(): want HTTP status code %v; got %v", http.StatusOK, got.StatusCode)
			}

			if got.Header.Get("Content-Type") != "application/json" {
				t.Errorf("JSON.WriteJSON(): want HTTP Content-Type %v; got %v", "application/json", got.Header.Get("Content-Type"))
			}

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatal(err)
				}
			}(got.Body)
			//defer got.Body.Close()

			body, err := io.ReadAll(got.Body)
			if err != nil {
				t.Fatal(err)
			}
			bytes.TrimSpace(body)

			if string(body) != tt.want {
				t.Errorf("JSON.WriteJSON(): want %q; got %q", tt.want, string(body))
			}
		})
	}
}
