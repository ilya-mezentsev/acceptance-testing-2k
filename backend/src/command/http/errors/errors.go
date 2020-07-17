package errors

import "errors"

var (
	NoJSONInResponse  = errors.New("no-json-in-response")
	NoJSONInArguments = errors.New("no-json-in-arguments")
)
