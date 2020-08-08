package request_decoder

import (
	"test_utils"
	"testing"
)

type (
	T1 struct {
		X int
	}

	T2 struct {
		X, Y int
	}
)

func TestDecodeSuccessT1(t *testing.T) {
	var x T1
	err := Decode(test_utils.GetReadCloser(`{"x": 1}`), &x)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(1, x.X, t)
}

func TestDecodeSuccessT2(t *testing.T) {
	var x T2
	err := Decode(test_utils.GetReadCloser(`{"x": 1, "y": 2}`), &x)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(1, x.X, t)
	test_utils.AssertEqual(2, x.Y, t)
}

func TestDecodeError(t *testing.T) {
	err := Decode(test_utils.GetReadCloser(`1`), &T1{})

	test_utils.AssertNotNil(err, t)
}
