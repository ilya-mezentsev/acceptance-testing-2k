package command

import "errors"

type BadReadCloser struct {
}

func (b BadReadCloser) Read([]byte) (n int, err error) {
	return 0, errors.New("some-error")
}

func (b BadReadCloser) Close() error {
	return errors.New("some-error")
}
