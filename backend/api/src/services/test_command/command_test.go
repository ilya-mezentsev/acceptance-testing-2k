package test_command

import (
	"api_meta/mock/services"
	"api_meta/models"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"services/errors"
	"services/plugins/response_factory"
	"test_utils"
	"testing"
)

var (
	repository            = services.TestCommandsRepositoryMock{}
	s                     = New(&repository)
	expectedSuccessStatus = response_factory.DefaultResponse().GetStatus()
	expectedErrorStatus   = response_factory.ErrorResponse(nil).GetStatus()
)

func init() {
	repository.Reset()
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestService_CreateSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(
		fmt.Sprintf(`{"account_hash": "%s", "command_settings": {
			"name": "CREATE",
			"object_name": "FOO",
			"method": "POST",
			"base_url": "https://link.com/api/v1",
			"endpoint": "user/settings",
			"pass_arguments_in_url": true
		}}`, services.SomeHash),
	))

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertNotNil(response.GetData(), t)
	test_utils.AssertEqual(
		"CREATE",
		repository.Commands[services.SomeHash][0].Name,
		t,
	)
	test_utils.AssertEqual(
		"FOO",
		repository.Commands[services.SomeHash][0].ObjectName,
		t,
	)
	test_utils.AssertEqual(
		"POST",
		repository.Commands[services.SomeHash][0].Method,
		t,
	)
	test_utils.AssertEqual(
		repository.Commands[services.SomeHash][0].BaseURL,
		"https://link.com/api/v1",
		t,
	)
	test_utils.AssertEqual(
		repository.Commands[services.SomeHash][0].Endpoint,
		"user/settings",
		t,
	)
	test_utils.AssertEqual(
		repository.Commands[services.SomeHash][0].PassArgumentsInURL,
		true,
		t,
	)
}

func TestService_CreateDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(`1`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.DecodingRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateInvalidRequestError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(
		fmt.Sprintf(`{"account_hash": "%s", "command_settings": {
			"name": "@#$!@#4",
			"object_name": "",
			"method": "HEAD",
			"base_url": "bad-url",
			"endpoint": ""
		}}`, services.SomeHash),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateCommandExistsError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(
		fmt.Sprintf(`{"account_hash": "%s", "command_settings": {
			"name": "%s",
			"object_name": "%s",
			"method": "POST",
			"base_url": "https://link.com/api/v1",
			"endpoint": "user/settings",
			"pass_arguments_in_url": true
		}}`,
			services.PredefinedAccountHash,
			services.PredefinedTestCommand1.Name,
			services.PredefinedTestCommand1.ObjectName,
		),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.UniqueEntityExistsError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_CreateRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.Create(test_utils.GetReadCloser(
		fmt.Sprintf(`{"account_hash": "%s", "command_settings": {
			"name": "CREATE",
			"object_name": "FOO",
			"method": "POST",
			"base_url": "https://link.com/api/v1",
			"endpoint": "user/settings",
			"pass_arguments_in_url": true
		}}`, services.BadAccountHash),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToCreateTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_GetAllSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.GetAll(services.PredefinedAccountHash)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	for expectedCommandIndex, expectedCommand := range repository.Commands[services.PredefinedAccountHash] {
		currentCommand := response.GetData().([]models.TestCommandRecord)[expectedCommandIndex]

		test_utils.AssertEqual(currentCommand.Name, expectedCommand.Name, t)
		test_utils.AssertEqual(currentCommand.ObjectName, expectedCommand.ObjectName, t)
		test_utils.AssertEqual(currentCommand.Method, expectedCommand.Method, t)
		test_utils.AssertEqual(expectedCommand.BaseURL, currentCommand.BaseURL, t)
		test_utils.AssertEqual(expectedCommand.Endpoint, currentCommand.Endpoint, t)
		test_utils.AssertEqual(expectedCommand.PassArgumentsInURL, currentCommand.PassArgumentsInURL, t)
		test_utils.AssertEqual(
			expectedCommand.Headers,
			currentCommand.Headers,
			t,
		)
		test_utils.AssertEqual(
			expectedCommand.Cookies,
			currentCommand.Cookies,
			t,
		)
	}
}

func TestService_GetAllInvalidRequestError(t *testing.T) {
	defer repository.Reset()

	response := s.GetAll("some-hash")

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertEqual(
		unableToFetchTestCommandsCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_GetAllRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.GetAll(services.BadAccountHash)
	test_utils.AssertEqual(
		unableToFetchTestCommandsCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_GetSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Get(services.PredefinedAccountHash, services.PredefinedTestCommand1.Hash)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	expectedCommand, currentCommand :=
		services.PredefinedTestCommand1, response.GetData().(models.TestCommandRecord)
	test_utils.AssertEqual(currentCommand.Name, expectedCommand.Name, t)
	test_utils.AssertEqual(currentCommand.ObjectName, expectedCommand.ObjectName, t)
	test_utils.AssertEqual(currentCommand.Method, expectedCommand.Method, t)
	test_utils.AssertEqual(expectedCommand.BaseURL, currentCommand.BaseURL, t)
	test_utils.AssertEqual(expectedCommand.Endpoint, currentCommand.Endpoint, t)
	test_utils.AssertEqual(expectedCommand.PassArgumentsInURL, currentCommand.PassArgumentsInURL, t)
	test_utils.AssertEqual(
		expectedCommand.Headers,
		currentCommand.Headers,
		t,
	)
	test_utils.AssertEqual(
		expectedCommand.Cookies,
		currentCommand.Cookies,
		t,
	)
}

func TestService_GetInvalidRequestError(t *testing.T) {
	defer repository.Reset()

	response := s.Get("some-hash", "some-hash")

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertEqual(
		unableToFetchTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_GetRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.Get(services.BadAccountHash, services.PredefinedTestCommand1.Hash)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertEqual(
		unableToFetchTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_UpdateSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Update(test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"account_hash": "%s",
			"update_payload": [
				{"hash": "%s", "field_name": "command_setting:name", "new_value": "FOO"},
				{"hash": "%s", "field_name": "command:object_name", "new_value": "BAR"}
			]
		}`,
			services.PredefinedAccountHash,
			services.PredefinedTestCommand1.Hash,
			services.PredefinedTestCommand1.Hash),
	))

	var updatedCommand models.TestCommandRecord
	_ = repository.Get(
		services.PredefinedAccountHash,
		services.PredefinedTestCommand1.Hash,
		&updatedCommand,
	)
	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)
	test_utils.AssertEqual("FOO", updatedCommand.Name, t)
	test_utils.AssertEqual("BAR", updatedCommand.ObjectName, t)
}

func TestService_UpdateDecodeBodyError(t *testing.T) {
	defer repository.Reset()

	response := s.Update(test_utils.GetReadCloser(`1`))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertTrue(response.HasData(), t)
	test_utils.AssertEqual(
		unableToUpdateTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.DecodingRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_UpdateInvalidRequestError(t *testing.T) {
	defer repository.Reset()

	response := s.Update(test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"account_hash": "some-hash",
			"update_payload": [
				{"hash": "%s", "field_name": "command_setting:name", "new_value": "FOO"},
				{"hash": "%s", "field_name": "command:object_name", "new_value": "BAR"}
			]
		}`,
			services.PredefinedTestCommand1.Hash,
			services.PredefinedTestCommand1.Hash),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertEqual(
		unableToUpdateTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_UpdateRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.Update(test_utils.GetReadCloser(
		fmt.Sprintf(`{
			"account_hash": "%s",
			"update_payload": [
				{"hash": "%s", "field_name": "command_setting:name", "new_value": "FOO"},
				{"hash": "%s", "field_name": "command:object_name", "new_value": "BAR"}
			]
		}`,
			services.BadAccountHash,
			services.PredefinedTestCommand1.Hash,
			services.PredefinedTestCommand1.Hash),
	))

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertEqual(
		unableToUpdateTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_DeleteSuccess(t *testing.T) {
	defer repository.Reset()

	response := s.Delete(services.PredefinedAccountHash, services.PredefinedTestCommand1.Hash)

	test_utils.AssertEqual(expectedSuccessStatus, response.GetStatus(), t)
	test_utils.AssertFalse(response.HasData(), t)
	test_utils.AssertNil(response.GetData(), t)
	for _, command := range repository.Commands[services.PredefinedAccountHash] {
		test_utils.AssertNotEqual(services.PredefinedTestCommand1.Hash, command.Hash, t)
	}
}

func TestService_DeleteInvalidRequestError(t *testing.T) {
	defer repository.Reset()

	response := s.Delete("some-hash", "some-hash")

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertEqual(
		unableToDeleteTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.InvalidRequestError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}

func TestService_DeleteRepositoryError(t *testing.T) {
	defer repository.Reset()

	response := s.Delete(services.BadAccountHash, services.PredefinedTestCommand1.Hash)

	test_utils.AssertEqual(expectedErrorStatus, response.GetStatus(), t)
	test_utils.AssertEqual(
		unableToDeleteTestCommandCode,
		response.GetData().(errors.ServiceError).Code,
		t,
	)
	test_utils.AssertEqual(
		errors.RepositoryError,
		response.GetData().(errors.ServiceError).Description,
		t,
	)
}
