package test_command

import (
	"api_meta/interfaces"
	"api_meta/models"
	"api_meta/types"
	"io"
	"services/errors"
	"services/plugins/hash"
	"services/plugins/logger"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
	"strings"
)

type service struct {
	logger     logger.CRUDEntityErrorsLogger
	repository interfaces.CRUDRepository
}

func New(repository interfaces.CRUDRepository) interfaces.CRUDService {
	return service{
		repository: repository,
		logger:     logger.CRUDEntityErrorsLogger{EntityName: entityName},
	}
}

func (s service) Create(request io.ReadCloser) interfaces.Response {
	var createTestCommandRequest models.CreateTestCommandRequest
	err := request_decoder.Decode(request, &createTestCommandRequest)
	if err != nil {
		s.logger.LogCreateEntityDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestCommandCode,
			Description: errors.DecodingRequestError,
		})
	}

	testCommandRecord := models.TestCommandRecord{
		CommandSettings: createTestCommandRequest.TestCommand.CommandSettings,
		Headers:         createTestCommandRequest.TestCommand.Headers.ReduceToRecordable(),
		Cookies:         createTestCommandRequest.TestCommand.Cookies.ReduceToRecordable(),
	}
	testCommandRecord.Hash = hash.GetHashWithTimeAsKey(testCommandRecord.Name)

	if !validation.IsValid(&testCommandRecord) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestCommandCode,
			Description: errors.InvalidRequestError,
		})
	}

	err = s.repository.Create(createTestCommandRequest.AccountHash, map[string]interface{}{
		"name":                  testCommandRecord.Name,
		"hash":                  testCommandRecord.Hash,
		"object_name":           testCommandRecord.ObjectName,
		"method":                testCommandRecord.Method,
		"base_url":              testCommandRecord.BaseURL,
		"endpoint":              testCommandRecord.Endpoint,
		"pass_arguments_in_url": testCommandRecord.PassArgumentsInURL,
		"command_headers":       testCommandRecord.Headers,
		"command_cookies":       testCommandRecord.Cookies,
	})
	if err != nil {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"create_test_command_request": createTestCommandRequest,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestCommandCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s service) GetAll(accountHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestCommandsCode,
			Description: errors.InvalidRequestError,
		})
	}

	var testCommands []models.TestCommandRecord
	err := s.repository.GetAll(accountHash, &testCommands)
	if err != nil {
		s.logger.LogGetAllEntitiesRepositoryError(err, map[string]interface{}{
			"account_hash": accountHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestCommandsCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.SuccessResponse(s.testCommandsRecordToTestCommandRequest(testCommands))
}

func (s service) testCommandsRecordToTestCommandRequest(
	testCommands []models.TestCommandRecord,
) []models.TestCommandRequest {
	var commands []models.TestCommandRequest
	for _, command := range testCommands {
		commands = append(commands, models.TestCommandRequest{
			CommandSettings: command.CommandSettings,
			Headers:         s.keyValueToMapping(command.Headers),
			Cookies:         s.keyValueToMapping(command.Cookies),
		})
	}

	return commands
}

func (s service) keyValueToMapping(keyValues string) types.Mapping {
	var mapping types.Mapping = map[string]string{}
	for _, keyValue := range strings.Split(keyValues, ";") {
		keyValue := strings.Split(keyValue, "=")
		mapping[keyValue[0]] = keyValue[1]
	}

	return mapping
}

func (s service) Get(accountHash, testCommandHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(testCommandHash) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestCommandCode,
			Description: errors.InvalidRequestError,
		})
	}

	var testCommand models.TestCommandRecord
	err := s.repository.Get(accountHash, testCommandHash, &testCommand)
	if err != nil {
		s.logger.LogGetEntityRepositoryError(err, map[string]interface{}{
			"account_hash":      accountHash,
			"test_command_hash": testCommandHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestCommandCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.SuccessResponse(models.TestCommandRequest{
		CommandSettings: testCommand.CommandSettings,
		Headers:         s.keyValueToMapping(testCommand.Headers),
		Cookies:         s.keyValueToMapping(testCommand.Cookies),
	})
}

func (s service) Update(request io.ReadCloser) interfaces.Response {
	var updateTestCommandRequest models.UpdateRequest
	err := request_decoder.Decode(request, &updateTestCommandRequest)
	if err != nil {
		s.logger.LogUpdateEntityDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToUpdateTestCommandCode,
			Description: errors.DecodingRequestError,
		})
	}

	if !validation.IsValid(&updateTestCommandRequest) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToUpdateTestCommandCode,
			Description: errors.InvalidRequestError,
		})
	}

	err = s.repository.Update(updateTestCommandRequest.AccountHash, updateTestCommandRequest.UpdatePayload)
	if err != nil {
		s.logger.LogUpdateEntityRepositoryError(err, map[string]interface{}{
			"account_hash":   updateTestCommandRequest.AccountHash,
			"update_payload": updateTestCommandRequest.UpdatePayload,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToUpdateTestCommandCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s service) Delete(accountHash, testCommandHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(testCommandHash) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToDeleteTestCommandCode,
			Description: errors.InvalidRequestError,
		})
	}

	err := s.repository.Delete(accountHash, testCommandHash)
	if err != nil {
		s.logger.LogDeleteEntityRepositoryError(err, map[string]interface{}{
			"account_hash":      accountHash,
			"test_command_hash": testCommandHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToDeleteTestCommandCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}
