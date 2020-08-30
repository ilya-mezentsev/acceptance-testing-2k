package test_command

import (
	"api_meta/interfaces"
	"api_meta/models"
	"api_meta/types"
	"errors"
	"io"
	servicesErrors "services/errors"
	"services/plugins/hash"
	"services/plugins/logger"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

type service struct {
	logger         logger.CRUDEntityErrorsLogger
	crudRepository interfaces.CRUDRepository
}

func New(crudRepository interfaces.CRUDRepository) interfaces.CRUDService {
	return service{
		crudRepository: crudRepository,
		logger:         logger.CRUDEntityErrorsLogger{EntityName: entityName},
	}
}

func (s service) Create(request io.ReadCloser) interfaces.Response {
	var createTestCommandRequest models.CreateTestCommandRequest
	err := request_decoder.Decode(request, &createTestCommandRequest)
	if err != nil {
		s.logger.LogCreateEntityDecodeError(err)

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToCreateTestCommandCode,
			Description: servicesErrors.DecodingRequestError,
		})
	}

	commandSettings := createTestCommandRequest.CommandSettings
	commandHash := hash.Md5WithTimeAsKey(commandSettings.Name)
	commandSettings.Hash = commandHash
	if !validation.IsValid(&commandSettings) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToCreateTestCommandCode,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	err = s.crudRepository.Create(createTestCommandRequest.AccountHash, map[string]interface{}{
		"name":                  commandSettings.Name,
		"hash":                  commandSettings.Hash,
		"object_name":           commandSettings.ObjectName,
		"method":                commandSettings.Method,
		"base_url":              commandSettings.BaseURL,
		"endpoint":              commandSettings.Endpoint,
		"pass_arguments_in_url": commandSettings.PassArgumentsInURL,
	})
	if errors.As(err, &types.UniqueEntityAlreadyExists{}) {
		s.logger.LogCreateEntityUniqueConstraintError(err, map[string]interface{}{
			"create_test_command_request": createTestCommandRequest,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToCreateTestCommandCode,
			Description: servicesErrors.UniqueEntityExistsError,
		})
	} else if err != nil {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"create_test_command_request": createTestCommandRequest,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToCreateTestCommandCode,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.SuccessResponse(models.CreatedTestCommandResponse{
		CommandHash: commandHash,
	})
}

func (s service) GetAll(accountHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToFetchTestCommandsCode,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	var testCommands []models.TestCommandRecord
	err := s.crudRepository.GetAll(accountHash, &testCommands)
	if err != nil {
		s.logger.LogGetAllEntitiesRepositoryError(err, map[string]interface{}{
			"account_hash": accountHash,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToFetchTestCommandsCode,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.SuccessResponse(testCommands)
}

func (s service) Get(accountHash, testCommandHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(testCommandHash) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToFetchTestCommandCode,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	var testCommand models.TestCommandRecord
	err := s.crudRepository.Get(accountHash, testCommandHash, &testCommand)
	if err != nil {
		s.logger.LogGetEntityRepositoryError(err, map[string]interface{}{
			"account_hash":      accountHash,
			"test_command_hash": testCommandHash,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToFetchTestCommandCode,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.SuccessResponse(testCommand)
}

func (s service) Update(request io.ReadCloser) interfaces.Response {
	var updateTestCommandRequest models.UpdateRequest
	err := request_decoder.Decode(request, &updateTestCommandRequest)
	if err != nil {
		s.logger.LogUpdateEntityDecodeError(err)

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToUpdateTestCommandCode,
			Description: servicesErrors.DecodingRequestError,
		})
	}

	if !validation.IsValid(&updateTestCommandRequest) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToUpdateTestCommandCode,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	err = s.crudRepository.Update(updateTestCommandRequest.AccountHash, updateTestCommandRequest.UpdatePayload)
	if err != nil {
		s.logger.LogUpdateEntityRepositoryError(err, map[string]interface{}{
			"account_hash":   updateTestCommandRequest.AccountHash,
			"update_payload": updateTestCommandRequest.UpdatePayload,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToUpdateTestCommandCode,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s service) Delete(accountHash, testCommandHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(testCommandHash) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToDeleteTestCommandCode,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	err := s.crudRepository.Delete(accountHash, testCommandHash)
	if err != nil {
		s.logger.LogDeleteEntityRepositoryError(err, map[string]interface{}{
			"account_hash":      accountHash,
			"test_command_hash": testCommandHash,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToDeleteTestCommandCode,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}
