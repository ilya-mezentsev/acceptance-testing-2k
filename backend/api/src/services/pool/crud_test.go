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
	pool := NewCRUD()
	pool.AddService("test", services.CRUDServiceMock{})

	service, response := pool.Get("test")

	test_utils.AssertNil(response, t)
	test_utils.AssertNotNil(service, t)
}

func TestCRUDServicesPool_GetError(t *testing.T) {
	pool := NewCRUD()

	service, response := pool.Get("test")

	responseData := response.GetData().(errors.ServiceError)
	test_utils.AssertEqual("error", response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(noServiceErrorCode, responseData.Code, t)
	test_utils.AssertEqual(
		pool.getNoServiceErrorDescription("test"),
		responseData.Description,
		t,
	)
	test_utils.AssertNil(service, t)
}
