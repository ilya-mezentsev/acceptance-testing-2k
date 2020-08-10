package validation

import (
	"api_meta/models"
	"io/ioutil"
	"log"
	"os"
	"services/plugins/hash"
	"strings"
	"test_utils"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestIsValidStruct(t *testing.T) {
	t.Run("valid struct TestObject", func(t *testing.T) {
		test_utils.AssertTrue(IsValid(&models.TestObject{
			Name: "some-name",
			Hash: hash.GetHashWithTimeAsKey("hash"),
		}), t)
	})

	t.Run("invalid hash", func(t *testing.T) {
		test_utils.AssertFalse(IsValid(&models.TestObject{
			Name: "some-name",
			Hash: "some-hash",
		}), t)

		test_utils.AssertFalse(IsValid(&models.TestObject{
			Name: "some-name",
			Hash: strings.Repeat(hash.GetHashWithTimeAsKey("hash"), 1000),
		}), t)
	})

	t.Run("invalid test object name", func(t *testing.T) {
		test_utils.AssertFalse(IsValid(&models.TestObject{
			Name: "",
			Hash: hash.GetHashWithTimeAsKey("hash"),
		}), t)

		test_utils.AssertFalse(IsValid(&models.TestObject{
			Name: strings.Repeat("1", 1000),
			Hash: hash.GetHashWithTimeAsKey("hash"),
		}), t)
	})

	t.Run("valid struct UpdateTestObjectRequest", func(t *testing.T) {
		test_utils.AssertTrue(IsValid(&models.UpdateTestObjectRequest{
			AccountHash: hash.GetHashWithTimeAsKey("hash"),
			UpdatePayload: []models.UpdateModel{
				{
					Hash:      hash.GetHashWithTimeAsKey("hash"),
					FieldName: "field_name",
				},
			},
		}), t)

		test_utils.AssertTrue(IsValid(&models.UpdateTestObjectRequest{
			AccountHash: hash.GetHashWithTimeAsKey("hash"),
			UpdatePayload: []models.UpdateModel{
				{
					Hash:      hash.GetHashWithTimeAsKey("hash"),
					FieldName: "update_target:field_name",
				},
			},
		}), t)
	})

	t.Run("invalid field name", func(t *testing.T) {
		test_utils.AssertFalse(IsValid(&models.UpdateTestObjectRequest{
			AccountHash: hash.GetHashWithTimeAsKey("hash"),
			UpdatePayload: []models.UpdateModel{
				{
					Hash:      hash.GetHashWithTimeAsKey("hash"),
					FieldName: "field-name",
				},
			},
		}), t)
	})

	t.Run("invalid slice argument", func(t *testing.T) {
		defer func() {
			test_utils.AssertEqual("slice argument is not struct", recover(), t)
		}()
		type x struct {
			y []int
		}

		IsValid(&x{y: []int{1, 2}})
		test_utils.AssertTrue(false, t)
	})

	t.Run("valid inner struct", func(t *testing.T) {
		type x struct {
			y models.TestObject
		}

		test_utils.AssertTrue(IsValid(&x{y: models.TestObject{
			Name: "some-name",
			Hash: hash.GetHashWithTimeAsKey("hash"),
		}}), t)
	})

	t.Run("invalid inner struct", func(t *testing.T) {
		type x struct {
			y models.TestObject
		}

		test_utils.AssertFalse(IsValid(&x{y: models.TestObject{
			Name: "",
			Hash: hash.GetHashWithTimeAsKey("hash"),
		}}), t)
	})

	t.Run("no need validation", func(t *testing.T) {
		type x struct {
			y int
		}

		test_utils.AssertTrue(IsValid(&x{y: 1}), t)
	})
}

func TestIsValidPanicNotPointer(t *testing.T) {
	defer func() {
		test_utils.AssertEqual("passed struct is not pointer", recover(), t)
	}()

	IsValid(models.TestObject{
		Name: strings.Repeat("1", 1000),
		Hash: hash.GetHashWithTimeAsKey("hash"),
	})
	test_utils.AssertTrue(false, t)
}

func TestIsValidPanicNotStruct(t *testing.T) {
	defer func() {
		test_utils.AssertEqual("cannot validate type: int", recover().(error).Error(), t)
	}()

	x := 1
	IsValid(&x)
	test_utils.AssertTrue(false, t)
}
