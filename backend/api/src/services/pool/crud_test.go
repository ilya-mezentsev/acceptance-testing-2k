package pool

import (
	"api_meta/mock/services"
	"io/ioutil"
	"log"
	"os"
	"services/errors"
	"test_utils"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestCRUDServicesPool_GetSuccess(t *testing.T) {
	pool := New()
	pool.AddService("test", services.CRUDServiceMock{})

	service := pool.Get("test")

	test_utils.AssertNotNil(service, t)
}

func TestCRUDServicesPool_GetNotExistsService(t *testing.T) {
	pool := New()

	service := pool.Get("test")

	_, ok := service.(defaultCRUDService)
	test_utils.AssertTrue(ok, t)

	var (
		response     = service.Create(nil)
		responseData errors.ServiceError
	)
	responseData = response.GetData().(errors.ServiceError)
	test_utils.AssertEqual("error", response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(noServiceErrorCode, responseData.Code, t)
	test_utils.AssertEqual(
		pool.getNoServiceErrorDescription("test"),
		responseData.Description,
		t,
	)

	response = service.GetAll("")
	responseData = response.GetData().(errors.ServiceError)
	test_utils.AssertEqual("error", response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(noServiceErrorCode, responseData.Code, t)
	test_utils.AssertEqual(
		pool.getNoServiceErrorDescription("test"),
		responseData.Description,
		t,
	)

	response = service.Get("", "")
	responseData = response.GetData().(errors.ServiceError)
	test_utils.AssertEqual("error", response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(noServiceErrorCode, responseData.Code, t)
	test_utils.AssertEqual(
		pool.getNoServiceErrorDescription("test"),
		responseData.Description,
		t,
	)

	response = service.Update(nil)
	responseData = response.GetData().(errors.ServiceError)
	test_utils.AssertEqual("error", response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(noServiceErrorCode, responseData.Code, t)
	test_utils.AssertEqual(
		pool.getNoServiceErrorDescription("test"),
		responseData.Description,
		t,
	)

	response = service.Delete("", "")
	responseData = response.GetData().(errors.ServiceError)
	test_utils.AssertEqual("error", response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(noServiceErrorCode, responseData.Code, t)
	test_utils.AssertEqual(
		pool.getNoServiceErrorDescription("test"),
		responseData.Description,
		t,
	)
}
