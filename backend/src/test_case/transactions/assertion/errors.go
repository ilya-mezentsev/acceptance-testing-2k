package assertion

import "errors"

var (
	cannotAccessValueByPath = errors.New("cannot-access-value-by-path")
	invalidPath             = errors.New("invalid-path")
	variableIsNotDefined    = errors.New("variable-is-not-defined")
	assertionFailed         = errors.New("assertion-failed")
	indexOutOfBounds        = errors.New("index-out-of-bounds")
	invalidNumberForIndex   = errors.New("invalid-number-for-index")
)
