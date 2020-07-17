package arguments_builder

import (
	"command/http/errors"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"utils"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestBuild_EmptyArguments(t *testing.T) {
	args := Build(``)

	ampersandSeparated, err := args.AmpersandSeparated()

	utils.AssertNil(err, t)
	utils.AssertEqual(``, args.Value(), t)
	utils.AssertEqual(``, ampersandSeparated, t)
}

func TestBuild_WithFlatJSONArguments(t *testing.T) {
	args := Build(`{"x": 1}`)

	ampersandSeparated, err := args.AmpersandSeparated()

	utils.AssertNil(err, t)
	utils.AssertEqual(`{"x": 1}`, args.Value(), t)
	utils.AssertEqual(`x=1`, ampersandSeparated, t)
}

func TestArguments_AmpersandSeparatedWithComplexJSON(t *testing.T) {
	args := Build(`{"x": [1, 2, 3], "y": {"t1": 1, "t2": 2}}`)

	ampersandSeparated, err := args.AmpersandSeparated()

	utils.AssertNil(err, t)
	utils.AssertEqual(`{"x": [1, 2, 3], "y": {"t1": 1, "t2": 2}}`, args.Value(), t)
	utils.AssertEqual(`x=[1,2,3]&y={"t1":1,"t2":2}`, ampersandSeparated, t)
}

func TestBuild_WithNoJSONArguments(t *testing.T) {
	args := Build(`1`)

	_, err := args.AmpersandSeparated()

	utils.AssertErrorsEqual(errors.NoJSONInArguments, err, t)
	utils.AssertEqual(`1`, args.Value(), t)
}
