package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/choonsiong/golib/v2/logger"
	"io"
	"net/http"
)

type JSON struct {
	AllowUnknownFields bool
	MaxBytes           int
	Logger             logger.Logger
}

// JSONResponse is the type used for sending JSON.
type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ReadJSON tries to read the body of a request and coverts from json into
// a go data variable.
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

	j.Logger.PrintDebug("ReadJSON()", map[string]string{
		"AllowUnknownFields": fmt.Sprintf("%v", j.AllowUnknownFields),
		"MaxBytes":           fmt.Sprintf("%v", maxBytes),
	})

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

	j.Logger.PrintDebug("ReadJSON()", map[string]string{
		"Data": fmt.Sprintf("%v", data),
	})

	err = d.Decode(&struct{}{})
	if err != io.EOF {
		return fmt.Errorf("ReadJSON(): %w: %v", ErrMultipleJSONValue, err)
	}

	return nil
}

// WriteJSON takes a response status code and arbitrary data and writes json
// to the client.
func (j *JSON) WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
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

// ErrorJSON takes an error and optionally a status code, generates and
// send a JSON error message.
func (j *JSON) ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return j.WriteJSON(w, statusCode, payload)
}

// PushJSONToRemote posts arbitrary data to a remote URL as JSON, and returns
// the response, status code and error if any. The final parameter client is
// optional. If none is specified, we use the standard http.Client.
func (j *JSON) PushJSONToRemote(uri string, data any, client ...*http.Client) (*http.Response, int, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	httpClient := &http.Client{}

	if len(client) > 0 {
		// client is a variadic variable, which means it allows us to send either 0 or more
		// values
		httpClient = client[0]
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	return resp, resp.StatusCode, nil
}
