package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JSON struct {
	MaxBytes           int
	AllowUnknownFields bool
}

func (j *JSON) ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1024 * 1024 // 1MB

	if j.MaxBytes != 0 {
		maxBytes = j.MaxBytes
	}

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	d := json.NewDecoder(r.Body)

	if !j.AllowUnknownFields {
		d.DisallowUnknownFields()
	}

	err := d.Decode(data)
	if err != nil {
		return fmt.Errorf("READJSON(): %w: %v", ErrDecodeJSON, err)
	}
	//if err != nil {
	//	var syntaxError *json.SyntaxError
	//	var unmarshalTypeError *json.UnmarshalTypeError
	//	var invalidUnmarshalError *json.InvalidUnmarshalError
	//
	//	switch {
	//	case errors.As(err, &syntaxError):
	//		return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
	//	case errors.Is(err, io.ErrUnexpectedEOF):
	//		return errors.New("body contains badly-formed JSON")
	//	case errors.As(err, &unmarshalTypeError):
	//		if unmarshalTypeError.Field != "" {
	//			return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
	//		}
	//		return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
	//	case errors.Is(err, io.EOF):
	//		return errors.New("body must not be empty")
	//	case strings.HasPrefix(err.Error(), "json: unknown field"):
	//		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
	//		return fmt.Errorf("body contains unknown key %s", fieldName)
	//	case err.Error() == "http: request body too large":
	//		return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
	//	case errors.As(err, &invalidUnmarshalError):
	//		return fmt.Errorf("error unmarshalling JSON: %s", err.Error())
	//	default:
	//		return fmt.Errorf("READJSON(): %w: %v", ErrDecodeJSON, err)
	//	}
	//}

	err = d.Decode(&struct{}{})
	if err != io.EOF {
		return fmt.Errorf("ReadJSON(): %w: %v", ErrMultipleJSONValue, err)
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
