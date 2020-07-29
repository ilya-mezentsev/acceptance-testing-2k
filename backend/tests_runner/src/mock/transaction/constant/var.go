package constant

import "errors"

var (
	BuildCommandError = errors.New("build-error")
	RunCommandError   = errors.New("run-error")
)
