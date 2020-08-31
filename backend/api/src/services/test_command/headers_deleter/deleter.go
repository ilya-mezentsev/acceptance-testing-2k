package headers_deleter

import (
	"api_meta/interfaces"
	servicesErrors "services/errors"
	"services/plugins/logger"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

type Service struct {
	logger     logger.CRUDEntityErrorsLogger
	repository interfaces.TestCommandHeadersDeleterRepository
}

func New(repository interfaces.TestCommandHeadersDeleterRepository) Service {
	return Service{
		logger:     logger.CRUDEntityErrorsLogger{EntityName: entityName},
		repository: repository,
	}
}

func (s Service) Delete(accountHash, headerHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(headerHash) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToDeleteCommandHeader,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	err := s.repository.DeleteHeader(accountHash, headerHash)
	if err != nil {
		s.logger.LogDeleteEntityRepositoryError(err, map[string]interface{}{
			"account_hash": accountHash,
			"header_hash":  headerHash,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToDeleteCommandHeader,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}
