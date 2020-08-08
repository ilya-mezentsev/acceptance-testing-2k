package tests_reader

import "errors"

var (
	TestsFileNotFound = errors.New("tests-file-not-found")
	UnknownError      = errors.New("unknown-error")
)
