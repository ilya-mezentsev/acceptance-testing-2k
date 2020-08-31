package cookies_deleter

import (
	"api_meta/interfaces"
	servicesErrors "services/errors"
	"services/plugins/logger"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

type Service struct {
	logger     logger.CRUDEntityErrorsLogger
	repository interfaces.TestCommandCookiesDeleterRepository
}

func New(repository interfaces.TestCommandCookiesDeleterRepository) Service {
	return Service{
		logger:     logger.CRUDEntityErrorsLogger{EntityName: entityName},
		repository: repository,
	}
}

func (s Service) Delete(accountHash, cookieHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(cookieHash) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToDeleteCommandCookie,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	err := s.repository.DeleteCookie(accountHash, cookieHash)
	if err != nil {
		s.logger.LogDeleteEntityRepositoryError(err, map[string]interface{}{
			"account_hash": accountHash,
			"header_hash":  cookieHash,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToDeleteCommandCookie,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}
