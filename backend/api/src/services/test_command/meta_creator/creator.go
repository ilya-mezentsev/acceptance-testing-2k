package meta_creator

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
	repository interfaces.TestCommandKeyValueCreatorRepository
}

func New(repository interfaces.TestCommandKeyValueCreatorRepository) Service {
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
