package cookies

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
	var massCookiesCreateRequest models.MassCookiesCreateRequest
	err := request_decoder.Decode(request, &massCookiesCreateRequest)
	if err != nil {
		s.logger.LogCreateEntityDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassAddCookies,
			Description: errors.DecodingRequestError,
		})
	}

	newCommandMeta := models.CommandMeta{
		Cookies: s.prepareCookiesToCreation(massCookiesCreateRequest),
	}
	if !validation.IsMd5Hash(massCookiesCreateRequest.AccountHash) ||
		!validation.IsValid(&newCommandMeta) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassAddCookies,
			Description: errors.InvalidRequestError,
		})
	}

	err = s.repository.Create(massCookiesCreateRequest.AccountHash, newCommandMeta)
	if err != nil {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"mass_cookies_create_request": massCookiesCreateRequest,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassAddCookies,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) prepareCookiesToCreation(
	massCookiesCreateRequest models.MassCookiesCreateRequest,
) []models.KeyValueMapping {
	var cookies []models.KeyValueMapping
	for _, updateTarget := range massCookiesCreateRequest.CommandHashes {
		for _, cookie := range massCookiesCreateRequest.Cookies {
			cookies = append(cookies, models.KeyValueMapping{
				Hash:        hash.Md5WithTimeAsKey(updateTarget.Hash),
				Key:         cookie.Key,
				Value:       cookie.Value,
				CommandHash: updateTarget.Hash,
			})
		}
	}

	return cookies
}
