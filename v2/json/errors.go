package json

import "errors"

var (
	ErrDecodeJSON        = errors.New("json: error decoding json data")
	ErrMultipleJSONValue = errors.New("json: body must have only a single json value")
)
