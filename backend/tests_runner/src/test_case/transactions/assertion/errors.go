package assertion

import "errors"

var (
	variableIsNotDefined = errors.New("variable-is-not-defined")
	assertionFailed      = errors.New("assertion-failed")
)
