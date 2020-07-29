package arguments_builder

import (
	"command/http/errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"test_utils"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestBuild_EmptyArguments(t *testing.T) {
	args := Build(``)

	ampersandSeparated, err := args.AmpersandSeparated()

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(``, args.Value(), t)
	test_utils.AssertEqual(``, ampersandSeparated, t)
	test_utils.AssertFalse(args.IsSlashSeparated(), t)
}

func TestBuild_WithFlatJSONArguments(t *testing.T) {
	args := Build(`{"x": 1}`)

	ampersandSeparated, err := args.AmpersandSeparated()

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(`{"x": 1}`, args.Value(), t)
	test_utils.AssertEqual(`x=1`, ampersandSeparated, t)
	test_utils.AssertFalse(args.IsSlashSeparated(), t)
}

func TestArguments_AmpersandSeparatedWithComplexJSON(t *testing.T) {
	args := Build(`{"x": [1, 2, 3], "y": {"t1": 1, "t2": 2}}`)

	ampersandSeparated, err := args.AmpersandSeparated()

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(`{"x": [1, 2, 3], "y": {"t1": 1, "t2": 2}}`, args.Value(), t)
	test_utils.AssertTrue(strings.Contains(ampersandSeparated, `x=[1,2,3]`), t)
	test_utils.AssertTrue(strings.Contains(ampersandSeparated, `y={"t1":1,"t2":2}`), t)
	test_utils.AssertTrue(strings.Contains(ampersandSeparated, `&`), t)
}

func TestBuild_WithNoJSONArguments(t *testing.T) {
	args := Build(`1`)

	_, err := args.AmpersandSeparated()

	test_utils.AssertErrorsEqual(errors.NoJSONInArguments, err, t)
	test_utils.AssertEqual(`1`, args.Value(), t)
}

func TestBuild_WithSlashSeparatedArguments(t *testing.T) {
	args := Build(`id/nickname`)

	test_utils.AssertTrue(args.IsSlashSeparated(), t)
}

func TestBuild_WithSlashSeparatedOneArgument(t *testing.T) {
	test_utils.AssertTrue(Build(`hash-1`).IsSlashSeparated(), t)
	test_utils.AssertTrue(Build(`hash_1`).IsSlashSeparated(), t)
}
