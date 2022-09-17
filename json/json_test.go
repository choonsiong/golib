package json

import (
	"bytes"
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
			w := httptest.NewRecorder()

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
	h := http.Header{}
	h["testing"] = []string{"testing"}

	tests := []struct {
		name    string
		status  int
		data    *Person
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "valid struct",
			status:  http.StatusOK,
			data:    &Person{"foobar", 42, "foobar@example.com"},
			headers: http.Header{},
			want:    "{\n\t\"name\": \"foobar\",\n\t\"age\": 42,\n\t\"email\": \"foobar@example.com\"\n}",
			wantErr: nil,
		},
		{
			name:   "valid struct with missing field",
			status: http.StatusOK,
			data: &Person{
				Name: "foobar",
				Age:  42,
			},
			headers: http.Header{},
			want:    "{\n\t\"name\": \"foobar\",\n\t\"age\": 42,\n\t\"email\": \"\"\n}",
			wantErr: nil,
		},
		{
			name:    "nil data",
			status:  http.StatusOK,
			data:    nil,
			headers: http.Header{},
			want:    "null",
			wantErr: nil,
		},
		{
			name:    "want header",
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

			err = WriteJSON(w, tt.status, tt.data, tt.headers)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("want error %q; got %q", tt.wantErr, err)
				}
			}

			got := w.Result()

			if got.StatusCode != http.StatusOK {
				t.Errorf("want HTTP status code %v; got %v", http.StatusOK, got.StatusCode)
			}

			if got.Header.Get("Content-Type") != "application/json" {
				t.Errorf("want HTTP Content-Type %v; got %v", "application/json", got.Header.Get("Content-Type"))
			}

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatal(err)
				}
			}(got.Body)

			body, err := io.ReadAll(got.Body)
			if err != nil {
				t.Fatal(err)
			}
			bytes.TrimSpace(body)

			if string(body) != tt.want {
				t.Errorf("want %q; got %q", tt.want, string(body))
			}
		})
	}
}
