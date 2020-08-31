package meta

import (
	"api_meta/interfaces"
	"api_meta/models"
	"io"
	servicesErrors "services/errors"
	"services/plugins/hash"
	"services/plugins/logger"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

// Service for create headers/cookies
type Service struct {
	logger     logger.CRUDEntityErrorsLogger
	repository interfaces.TestCommandMetaRepository
}

func New(repository interfaces.TestCommandMetaRepository) Service {
	return Service{
		repository: repository,
		logger:     logger.CRUDEntityErrorsLogger{EntityName: entityName},
	}
}

func (s Service) Create(request io.ReadCloser) interfaces.Response {
	var createMetaRequest models.CreateMetaRequest
	err := request_decoder.Decode(request, &createMetaRequest)
	if err != nil {
		s.logger.LogCreateEntityDecodeError(err)

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToCreateCommandMeta,
			Description: servicesErrors.DecodingRequestError,
		})
	}

	s.addHashAndCommandHashToKeyValue(
		createMetaRequest.CommandHash,
		createMetaRequest.CommandMeta.Headers,
	)
	s.addHashAndCommandHashToKeyValue(
		createMetaRequest.CommandHash,
		createMetaRequest.CommandMeta.Cookies,
	)
	if !validation.IsValid(&createMetaRequest) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToCreateCommandMeta,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	err = s.repository.Create(createMetaRequest.AccountHash, createMetaRequest.CommandMeta)
	if err != nil {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"create_meta_request": createMetaRequest,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToCreateCommandMeta,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) addHashAndCommandHashToKeyValue(
	commandHash string,
	mapping []models.KeyValueMapping,
) {
	for index := range mapping {
		mapping[index].CommandHash = commandHash
		mapping[index].Hash = hash.Md5WithTimeAsKey(commandHash)
	}
}

func (s Service) Update(request io.ReadCloser) interfaces.Response {
	var updateMetaRequest models.UpdateMetaRequest
	err := request_decoder.Decode(request, &updateMetaRequest)
	if err != nil {
		s.logger.LogUpdateEntityDecodeError(err)

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToUpdateCommandMeta,
			Description: servicesErrors.DecodingRequestError,
		})
	}

	if !validation.IsValid(&updateMetaRequest) ||
		!s.isInvalidFieldName(append(updateMetaRequest.Headers, updateMetaRequest.Cookies...)) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToUpdateCommandMeta,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	err = s.repository.UpdateHeadersAndCookies(
		updateMetaRequest.AccountHash,
		updateMetaRequest.Headers,
		updateMetaRequest.Cookies,
	)
	if err != nil {
		s.logger.LogUpdateEntityRepositoryError(err, map[string]interface{}{
			"update_meta_request": updateMetaRequest,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToUpdateCommandMeta,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) isInvalidFieldName(updatePayload []models.UpdateModel) bool {
	for _, payload := range updatePayload {
		if !validation.IsKeyOrValue(payload.FieldName) {
			return false
		}
	}

	return true
}
