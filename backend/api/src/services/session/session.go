package session

import (
	"api_meta/interfaces"
	"api_meta/models"
	"net/http"
	"services/errors"
	"services/plugins/account_credentials"
	"services/plugins/logger"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

type Service struct {
	logger     logger.CRUDEntityErrorsLogger
	repository interfaces.SessionRepository
}

func New(repository interfaces.SessionRepository) Service {
	return Service{
		repository: repository,
		logger:     logger.CRUDEntityErrorsLogger{EntityName: entityName},
	}
}

func (s Service) CreateSession(w http.ResponseWriter, r *http.Request) interfaces.Response {
	var createSessionRequest models.CreateSessionRequest
	err := request_decoder.Decode(r.Body, &createSessionRequest)
	if err != nil {
		s.logger.LogCreateEntityDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateSessionCode,
			Description: errors.DecodingRequestError,
		})
	}

	if !validation.IsRegularName(createSessionRequest.Login) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateSessionCode,
			Description: errors.InvalidRequestError,
		})
	}

	accountHash := account_credentials.GenerateAccountHash(
		createSessionRequest.Login,
		createSessionRequest.Password,
	)
	accountExists, err := s.repository.AccountExists(accountHash)
	if err != nil {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"account_hash": accountHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateSessionCode,
			Description: errors.RepositoryError,
		})
	}

	if accountExists {
		setSessionCookie(w, accountHash)
		return response_factory.DefaultResponse()
	} else {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateSessionCode,
			Description: accountDoesNotExistsError,
		})
	}
}

func (s Service) GetSession(r *http.Request) interfaces.Response {
	cookie, err := getSessionCookie(r)
	if err != nil {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToGetSessionCode,
			Description: sessionCookieNotFoundError,
		})
	} else {
		return response_factory.SuccessResponse(models.SessionResponse{AccountHash: cookie.Value})
	}
}

func (s Service) DeleteSession(w http.ResponseWriter) interfaces.Response {
	unsetSessionCookie(w)
	return response_factory.DefaultResponse()
}
