package headers

import (
	"api_meta/interfaces"
	"api_meta/models"
	"io"
	"services/errors"
	"services/plugins/hash"
	"services/plugins/logger"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

type Service struct {
	logger     logger.CRUDEntityErrorsLogger
	repository interfaces.TestCommandMetaCreatorRepository
}

func New(repository interfaces.TestCommandMetaCreatorRepository) Service {
	return Service{
		logger:     logger.CRUDEntityErrorsLogger{EntityName: entityName},
		repository: repository,
	}
}

func (s Service) Create(request io.ReadCloser) interfaces.Response {
	var massHeadersCreateRequest models.MassHeadersCreateRequest
	err := request_decoder.Decode(request, &massHeadersCreateRequest)
	if err != nil {
		s.logger.LogCreateEntityDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassAddHeaders,
			Description: errors.DecodingRequestError,
		})
	}

	newCommandMeta := models.CommandMeta{
		Headers: s.prepareHeadersToCreation(massHeadersCreateRequest),
	}
	if !validation.IsMd5Hash(massHeadersCreateRequest.AccountHash) ||
		!validation.IsValid(&newCommandMeta) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassAddHeaders,
			Description: errors.InvalidRequestError,
		})
	}

	err = s.repository.Create(massHeadersCreateRequest.AccountHash, newCommandMeta)
	if err != nil {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"mass_headers_create_request": massHeadersCreateRequest,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassAddHeaders,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) prepareHeadersToCreation(
	massCookiesCreateRequest models.MassHeadersCreateRequest,
) []models.KeyValueMapping {
	var headers []models.KeyValueMapping
	for _, updateTarget := range massCookiesCreateRequest.CommandHashes {
		for _, cookie := range massCookiesCreateRequest.Headers {
			headers = append(headers, models.KeyValueMapping{
				Hash:        hash.Md5WithTimeAsKey(updateTarget.Hash),
				Key:         cookie.Key,
				Value:       cookie.Value,
				CommandHash: updateTarget.Hash,
			})
		}
	}

	return headers
}
