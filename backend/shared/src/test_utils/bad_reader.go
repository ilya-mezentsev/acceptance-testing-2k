package test_utils

import (
	"errors"
	"io"
)

type badReader struct {
}

func BadReader() io.Reader {
	return badReader{}
}

func (r badReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("some-error")
}
