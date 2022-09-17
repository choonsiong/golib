package json

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // 1MB

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	d := json.NewDecoder(r.Body)

	err := d.Decode(data)
	if err != nil {
		return err
	}

	err = d.Decode(&struct{}{})
	if err != io.EOF {
		return ErrMultipleJSONValue
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}
