package response_factory

import (
	"test_utils"
	"testing"
)

func TestDefaultResponse(t *testing.T) {
	response := DefaultResponse()

	test_utils.AssertEqual(statusOk, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)
}

func TestSuccessResponse(t *testing.T) {
	someData := `data`
	response := SuccessResponse(someData)

	test_utils.AssertEqual(statusOk, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(someData, response.GetData(), t)
}

func TestErrorResponse(t *testing.T) {
	someData := `data`
	response := ErrorResponse(someData)

	test_utils.AssertEqual(statusError, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(someData, response.GetData(), t)
}
