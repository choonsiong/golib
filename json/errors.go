package json

import "errors"

var (
	ErrMultipleJSONValue = errors.New("json: body must have only a single json value")
)
