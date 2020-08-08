package db_connector

import "errors"

var (
	DBFileNotFound = errors.New("db-file-not-found")
	UnknownError   = errors.New("unknown-error")
)
