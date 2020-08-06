package errors

import (
	"errors"
	"models"
)

var (
	DBFileNotFound    = errors.New("db-file-not-found")
	TestsFileNotFound = errors.New("tests-file-not-found")
	UnknownError      = errors.New("unknown-error")
)

var EmptyApplicationError models.ApplicationError
