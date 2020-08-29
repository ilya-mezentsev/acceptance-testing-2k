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
	expectedService := services.CRUDServiceMock{}
	pool.AddCRUDService("test", expectedService)

	createService := pool.GetCreateService("test")
	readService := pool.GetReadService("test")
	updateService := pool.GetUpdateService("test")
	deleteService := pool.GetDeleteService("test")

	test_utils.AssertEqual(expectedService, createService, t)
	test_utils.AssertEqual(expectedService, readService, t)
	test_utils.AssertEqual(expectedService, updateService, t)
	test_utils.AssertEqual(expectedService, deleteService, t)
}

func TestCRUDServicesPool_AddServiceSuccess(t *testing.T) {
	pool := New()
	expectedService := services.CRUDServiceMock{}
	pool.AddService(
		"test",
		[]string{CreateServiceOperationType, ReadServiceOperationType, UpdateServiceOperationType},
		expectedService,
	)

	createService := pool.GetCreateService("test")
	readService := pool.GetReadService("test")
	updateService := pool.GetUpdateService("test")
	deleteService := pool.GetDeleteService("test")

	test_utils.AssertEqual(expectedService, createService, t)
	test_utils.AssertEqual(expectedService, readService, t)
	test_utils.AssertEqual(expectedService, updateService, t)

	_, ok := deleteService.(defaultCRUDService)
	test_utils.AssertTrue(ok, t)

	pool.AddService(
		"test",
		[]string{DeleteServiceOperationType},
		expectedService,
	)
	deleteService = pool.GetDeleteService("test")

	test_utils.AssertEqual(expectedService, deleteService, t)
}

func TestCRUDServicesPool_AddService(t *testing.T) {
	defer func() {
		p := recover()
		test_utils.AssertEqual(
			"Unexpected operation type: bad-type",
			p.(string),
			t,
		)
	}()

	New().AddService("test", []string{"bad-type"}, services.CRUDServiceMock{})
	test_utils.AssertTrue(false, t)
}

func TestCRUDServicesPool_GetNotExistsService(t *testing.T) {
	pool := New()

	createService := pool.GetCreateService("test")

	_, ok := createService.(defaultCRUDService)
	test_utils.AssertTrue(ok, t)

	var (
		response     = createService.Create(nil)
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

	readService := pool.GetReadService("test")
	response = readService.GetAll("")
	responseData = response.GetData().(errors.ServiceError)
	test_utils.AssertEqual("error", response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(noServiceErrorCode, responseData.Code, t)
	test_utils.AssertEqual(
		pool.getNoServiceErrorDescription("test"),
		responseData.Description,
		t,
	)

	response = readService.Get("", "")
	responseData = response.GetData().(errors.ServiceError)
	test_utils.AssertEqual("error", response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(noServiceErrorCode, responseData.Code, t)
	test_utils.AssertEqual(
		pool.getNoServiceErrorDescription("test"),
		responseData.Description,
		t,
	)

	updateService := pool.GetUpdateService("test")
	response = updateService.Update(nil)
	responseData = response.GetData().(errors.ServiceError)
	test_utils.AssertEqual("error", response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(noServiceErrorCode, responseData.Code, t)
	test_utils.AssertEqual(
		pool.getNoServiceErrorDescription("test"),
		responseData.Description,
		t,
	)

	deleteService := pool.GetDeleteService("test")
	response = deleteService.Delete("", "")
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
