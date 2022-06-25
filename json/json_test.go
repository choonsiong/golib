package json

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func TestReadJSON(t *testing.T) {
	w := httptest.NewRecorder()

	tests := []struct {
		name    string
		body    io.Reader
		want    *Person
		wantErr error
	}{
		{
			name:    "valid json body",
			body:    strings.NewReader(`{"name": "foobar","age": 42,"email":"foobar@example.com"}`),
			want:    &Person{"foobar", 42, "foobar@example.com"},
			wantErr: nil,
		},
		{
			name:    "valid json body with extra field",
			body:    strings.NewReader(`{"name": "foobar","age": 42,"email":"foobar@example.com","mobile":"1234567"}`),
			want:    &Person{"foobar", 42, "foobar@example.com"},
			wantErr: nil,
		},
		{
			name:    "multiple json body",
			body:    strings.NewReader(`{"name": "foobar","age": 42,"email":"foobar@example.com"}{"name": "alice","age": 28,"email":"alice@example.com"}`),
			want:    &Person{"foobar", 42, "foobar@example.com"},
			wantErr: ErrMultipleJSONValue,
		},
	}

	p := new(Person)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodGet, "/", tt.body)
			r.Header.Set("Content-Type", "application/json")

			if err != nil {
				t.Fatal(err)
			}

			err = ReadJSON(w, r, p)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("want error %q; got %q", tt.wantErr, err)
				}
			}

			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("want %v; got %v", tt.want, p)
			}
		})
	}
}

func TestReadJSON_Decode(t *testing.T) {
	w := httptest.NewRecorder()

	tests := []struct {
		name    string
		body    io.Reader
		want    *Person
		wantErr error
	}{
		{
			name:    "nil receiver",
			body:    strings.NewReader(`{"name": "foobar","age": 42,"email":"foobar@example.com"}`),
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodGet, "/", tt.body)
			r.Header.Set("Content-Type", "application/json")

			if err != nil {
				t.Fatal(err)
			}

			err = ReadJSON(w, r, nil)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("want error %q; got %q", tt.wantErr, err)
				}
			}
		})
	}
}

func TestWriteJSON(t *testing.T) {

}
