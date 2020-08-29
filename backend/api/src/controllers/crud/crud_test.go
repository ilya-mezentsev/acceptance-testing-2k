package crud

import (
	"api_meta/mock/controllers"
	"api_meta/mock/services"
	"controllers/plugins/response_writer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"services/plugins/response_factory"
	"services/pool"
	"test_utils"
	"testing"
)

var (
	r           *mux.Router
	p           pool.CRUDServicesPool
	serviceMock controllers.CRUDServiceMock
)

func init() {
	r = mux.NewRouter()
	p = pool.New()

	serviceMock = controllers.CRUDServiceMock{CalledWith: map[string]interface{}{}}
	p.AddCRUDService("test", &serviceMock)

	Init(r, p)
}

func TestController_GetAllSuccess(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestGet(fmt.Sprintf(
		"%s/entity/test/%s/",
		server.URL, services.SomeHash,
	))

	var response response_writer.Response
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.DefaultResponse().GetStatus(), response.Status, t)
	test_utils.AssertEqual(services.SomeHash, serviceMock.CalledWith["GetAll"].(string), t)
}

func TestController_GetAllError(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestGet(fmt.Sprintf(
		"%s/entity/test/%s/",
		server.URL, controllers.BadAccountHash,
	))

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.ErrorResponse(nil).GetStatus(), response.Status, t)
	test_utils.AssertEqual(controllers.ErrorResponseCode, response.Data.(map[string]interface{})["code"], t)
	test_utils.AssertEqual(controllers.BadAccountHash, serviceMock.CalledWith["GetAll"].(string), t)
}

func TestController_GetOneSuccess(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestGet(fmt.Sprintf(
		"%s/entity/test/%s/%s/",
		server.URL, services.SomeHash, services.PredefinedTestCommand1.Hash,
	))

	var response response_writer.Response
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.DefaultResponse().GetStatus(), response.Status, t)
	test_utils.AssertEqual(services.SomeHash, serviceMock.CalledWith["Get"].(map[string]string)["account_hash"], t)
	test_utils.AssertEqual(
		services.PredefinedTestCommand1.Hash,
		serviceMock.CalledWith["Get"].(map[string]string)["entity_hash"],
		t,
	)
}

func TestController_GetOneError(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestGet(fmt.Sprintf(
		"%s/entity/test/%s/%s/",
		server.URL, controllers.BadAccountHash, controllers.BadEntityHash,
	))

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.ErrorResponse(nil).GetStatus(), response.Status, t)
	test_utils.AssertEqual(controllers.ErrorResponseCode, response.Data.(map[string]interface{})["code"], t)
	test_utils.AssertEqual(controllers.BadAccountHash, serviceMock.CalledWith["Get"].(map[string]string)["account_hash"], t)
	test_utils.AssertEqual(
		controllers.BadEntityHash,
		serviceMock.CalledWith["Get"].(map[string]string)["entity_hash"],
		t,
	)
}

func TestController_CreateSuccess(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestPost(fmt.Sprintf(
		"%s/entity/test/",
		server.URL,
	), `some-data`)

	var response response_writer.Response
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.DefaultResponse().GetStatus(), response.Status, t)
	test_utils.AssertEqual(`some-data`, serviceMock.CalledWith["Create"].(string), t)
}

func TestController_CreateError(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestPost(fmt.Sprintf(
		"%s/entity/test/",
		server.URL,
	), controllers.BadRequestData)

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.ErrorResponse(nil).GetStatus(), response.Status, t)
	test_utils.AssertEqual(controllers.ErrorResponseCode, response.Data.(map[string]interface{})["code"], t)
	test_utils.AssertEqual(controllers.BadRequestData, serviceMock.CalledWith["Create"].(string), t)
}

func TestController_UpdateSuccess(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestPatch(fmt.Sprintf(
		"%s/entity/test/",
		server.URL,
	), `some-data`)

	var response response_writer.Response
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.DefaultResponse().GetStatus(), response.Status, t)
	test_utils.AssertEqual(`some-data`, serviceMock.CalledWith["Update"].(string), t)
}

func TestController_UpdateError(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestPatch(fmt.Sprintf(
		"%s/entity/test/",
		server.URL,
	), controllers.BadRequestData)

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.ErrorResponse(nil).GetStatus(), response.Status, t)
	test_utils.AssertEqual(controllers.ErrorResponseCode, response.Data.(map[string]interface{})["code"], t)
	test_utils.AssertEqual(controllers.BadRequestData, serviceMock.CalledWith["Update"].(string), t)
}

func TestController_DeleteSuccess(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestDelete(fmt.Sprintf(
		"%s/entity/test/%s/%s/",
		server.URL, services.SomeHash, services.PredefinedTestCommand1.Hash,
	))

	var response response_writer.Response
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.DefaultResponse().GetStatus(), response.Status, t)
	test_utils.AssertEqual(services.SomeHash, serviceMock.CalledWith["Delete"].(map[string]string)["account_hash"], t)
	test_utils.AssertEqual(
		services.PredefinedTestCommand1.Hash,
		serviceMock.CalledWith["Delete"].(map[string]string)["entity_hash"],
		t,
	)
}

func TestController_DeleteError(t *testing.T) {
	server := test_utils.GetTestServer(r)
	defer server.Close()
	defer serviceMock.Reset()

	responseData := test_utils.RequestDelete(fmt.Sprintf(
		"%s/entity/test/%s/%s/",
		server.URL, controllers.BadAccountHash, controllers.BadEntityHash,
	))

	var response response_writer.ResponseWithData
	err := json.NewDecoder(responseData).Decode(&response)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(response_factory.ErrorResponse(nil).GetStatus(), response.Status, t)
	test_utils.AssertEqual(controllers.ErrorResponseCode, response.Data.(map[string]interface{})["code"], t)
	test_utils.AssertEqual(
		controllers.BadAccountHash,
		serviceMock.CalledWith["Delete"].(map[string]string)["account_hash"],
		t,
	)
	test_utils.AssertEqual(
		controllers.BadEntityHash,
		serviceMock.CalledWith["Delete"].(map[string]string)["entity_hash"],
		t,
	)
}
