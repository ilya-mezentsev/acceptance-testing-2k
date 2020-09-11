package value_path

import "errors"

var (
	CannotAccessValueByPath = errors.New("cannot-access-value-by-path")
	invalidPath             = errors.New("invalid-path")
	indexOutOfBounds        = errors.New("index-out-of-bounds")
	invalidNumberForIndex   = errors.New("invalid-number-for-index")
)
