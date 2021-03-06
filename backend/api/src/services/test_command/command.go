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
	logger                       logger.CRUDEntityErrorsLogger
	crudRepository               interfaces.CRUDRepository
	commandsMetaGetterRepository interfaces.TestCommandMetaGetterRepository
}

func New(
	crudRepository interfaces.CRUDRepository,
	commandsMetaGetterRepository interfaces.TestCommandMetaGetterRepository,
) interfaces.CRUDService {
	return service{
		crudRepository:               crudRepository,
		commandsMetaGetterRepository: commandsMetaGetterRepository,
		logger:                       logger.CRUDEntityErrorsLogger{EntityName: entityName},
	}
}

func (s service) Create(accountHash string, request io.ReadCloser) interfaces.Response {
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
	if !validation.IsMd5Hash(accountHash) || !validation.IsValid(&commandSettings) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToCreateTestCommandCode,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	err = s.crudRepository.Create(accountHash, map[string]interface{}{
		"name":                  commandSettings.Name,
		"hash":                  commandSettings.Hash,
		"object_hash":           commandSettings.ObjectHash,
		"method":                commandSettings.Method,
		"base_url":              commandSettings.BaseURL,
		"endpoint":              commandSettings.Endpoint,
		"timeout":               commandSettings.Timeout,
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

	var commandsSettings []models.CommandSettings
	getCommandsSettingsError := s.crudRepository.GetAll(accountHash, &commandsSettings)
	headers, cookies, getCommandMetaError :=
		s.commandsMetaGetterRepository.GetAllHeadersAndCookies(accountHash)

	if getCommandsSettingsError != nil || getCommandMetaError != nil {
		s.logger.LogGetAllEntitiesRepositoryError(getCommandsSettingsError, map[string]interface{}{
			"account_hash":              accountHash,
			"get_command_setting_error": getCommandsSettingsError,
			"get_command_meta_error":    getCommandMetaError,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToFetchTestCommandsCode,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.SuccessResponse(makeGetCommandsResponse(
		commandsSettings,
		headers,
		cookies,
	))
}

func makeGetCommandsResponse(
	commandsSettings []models.CommandSettings,
	headers,
	cookies []models.KeyValueMapping,
) []models.GetCommandResponse {
	var response []models.GetCommandResponse
	for _, commandsSettings := range commandsSettings {
		r := models.GetCommandResponse{
			CommandSettings: commandsSettings,
			CommandMeta:     models.CommandMeta{},
		}

		for _, header := range headers {
			if header.CommandHash == commandsSettings.Hash {
				r.Headers = append(r.Headers, header)
			}
		}

		for _, cookie := range cookies {
			if cookie.CommandHash == commandsSettings.Hash {
				r.Cookies = append(r.Cookies, cookie)
			}
		}

		response = append(response, r)
	}

	return response
}

func (s service) Get(accountHash, testCommandHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(testCommandHash) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToFetchTestCommandCode,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	var commandSettings models.CommandSettings
	getCommandSettingsError := s.crudRepository.Get(accountHash, testCommandHash, &commandSettings)
	headers, cookies, getCommandMetaError :=
		s.commandsMetaGetterRepository.GetCommandHeadersAndCookies(accountHash, testCommandHash)

	if getCommandSettingsError != nil || getCommandMetaError != nil {
		s.logger.LogGetEntityRepositoryError(getCommandSettingsError, map[string]interface{}{
			"account_hash":               accountHash,
			"test_command_hash":          testCommandHash,
			"get_command_settings_error": getCommandSettingsError,
			"get_command_meta_error":     getCommandMetaError,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToFetchTestCommandCode,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.SuccessResponse(models.GetCommandResponse{
		CommandSettings: commandSettings,
		CommandMeta: models.CommandMeta{
			Headers: headers,
			Cookies: cookies,
		},
	})
}

func (s service) Update(accountHash string, request io.ReadCloser) interfaces.Response {
	var updateTestCommandRequest models.UpdateTestCommandRequest
	err := request_decoder.Decode(request, &updateTestCommandRequest)
	if err != nil {
		s.logger.LogUpdateEntityDecodeError(err)

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToUpdateTestCommandCode,
			Description: servicesErrors.DecodingRequestError,
		})
	}

	if !validation.IsMd5Hash(accountHash) || !validation.IsValid(&updateTestCommandRequest) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToUpdateTestCommandCode,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	err = s.crudRepository.Update(
		accountHash,
		s.getUpdatePayload(updateTestCommandRequest),
	)
	if err != nil {
		s.logger.LogUpdateEntityRepositoryError(err, map[string]interface{}{
			"account_hash":        accountHash,
			"update_test_command": updateTestCommandRequest,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToUpdateTestCommandCode,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s service) getUpdatePayload(
	updateTestCommandRequest models.UpdateTestCommandRequest,
) []models.UpdateModel {
	var entities []models.UpdateModel

	nameChanged :=
		updateTestCommandRequest.UpdatedCommand.Name != updateTestCommandRequest.ExistsCommand.Name

	methodChanged :=
		updateTestCommandRequest.UpdatedCommand.Method != updateTestCommandRequest.ExistsCommand.Method

	baseURLChanged :=
		updateTestCommandRequest.UpdatedCommand.BaseURL != updateTestCommandRequest.ExistsCommand.BaseURL

	endpointChanged :=
		updateTestCommandRequest.UpdatedCommand.Endpoint != updateTestCommandRequest.ExistsCommand.Endpoint

	timeoutChanged :=
		updateTestCommandRequest.UpdatedCommand.Timeout != updateTestCommandRequest.ExistsCommand.Timeout

	passArgumentsFlagChanged :=
		updateTestCommandRequest.UpdatedCommand.PassArgumentsInURL !=
			updateTestCommandRequest.ExistsCommand.PassArgumentsInURL

	if nameChanged {
		entities = append(entities, models.UpdateModel{
			Hash:      updateTestCommandRequest.ExistsCommand.Hash,
			FieldName: "command:name",
			NewValue:  updateTestCommandRequest.UpdatedCommand.Name,
		})
	}

	if methodChanged {
		entities = append(entities, models.UpdateModel{
			Hash:      updateTestCommandRequest.ExistsCommand.Hash,
			FieldName: "command_setting:method",
			NewValue:  updateTestCommandRequest.UpdatedCommand.Method,
		})
	}

	if baseURLChanged {
		entities = append(entities, models.UpdateModel{
			Hash:      updateTestCommandRequest.ExistsCommand.Hash,
			FieldName: "command_setting:base_url",
			NewValue:  updateTestCommandRequest.UpdatedCommand.BaseURL,
		})
	}

	if endpointChanged {
		entities = append(entities, models.UpdateModel{
			Hash:      updateTestCommandRequest.ExistsCommand.Hash,
			FieldName: "command_setting:endpoint",
			NewValue:  updateTestCommandRequest.UpdatedCommand.Endpoint,
		})
	}

	if timeoutChanged {
		entities = append(entities, models.UpdateModel{
			Hash:      updateTestCommandRequest.ExistsCommand.Hash,
			FieldName: "command_setting:timeout",
			NewValue:  updateTestCommandRequest.UpdatedCommand.Timeout,
		})
	}

	if passArgumentsFlagChanged {
		entities = append(entities, models.UpdateModel{
			Hash:      updateTestCommandRequest.ExistsCommand.Hash,
			FieldName: "command_setting:pass_arguments_in_url",
			NewValue:  updateTestCommandRequest.UpdatedCommand.PassArgumentsInURL,
		})
	}

	return entities
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
