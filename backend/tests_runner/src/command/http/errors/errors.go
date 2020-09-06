package errors

import (
	"errors"
	"fmt"
)

var (
	CommandNotFound   = errors.New("command-not-found")
	NoJSONInResponse  = errors.New("no-json-in-response")
	NoJSONInArguments = errors.New("no-json-in-arguments")
)

func UnsuccessfulStatus(code int) error {
	return fmt.Errorf("unsuccessful-status: %d", code)
}
