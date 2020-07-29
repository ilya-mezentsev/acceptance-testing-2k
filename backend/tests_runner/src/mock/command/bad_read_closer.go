package command

import "errors"

type BadReadCLoser struct {
}

func (b BadReadCLoser) Read(p []byte) (n int, err error) {
	return 0, errors.New("some-error")
}

func (b BadReadCLoser) Close() error {
	return nil
}
