package session

import (
	"api_meta/interfaces"
	"api_meta/models"
	"db_connector"
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

	accountHash := account_credentials.GenerateAccountHash(createSessionRequest.Login)
	credentialsExists, err := s.repository.CredentialsExists(
		accountHash,
		createSessionRequest.Login,
		account_credentials.GenerateAccountPassword(
			createSessionRequest.Login,
			createSessionRequest.Password,
		),
	)
	if err != nil && err != db_connector.DBFileNotFound {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"login": createSessionRequest.Login,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateSessionCode,
			Description: errors.RepositoryError,
		})
	}

	// if err == db_connector.DBFileNotFound then account with
	// particular accountHash does not exits
	accountExists := err == nil

	if accountExists && credentialsExists {
		setSessionCookie(w, accountHash)
		return response_factory.SuccessResponse(models.SessionResponse{
			Login:       createSessionRequest.Login,
			AccountHash: accountHash,
		})
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
		return response_factory.SuccessResponse(models.SessionResponse{
			AccountHash: cookie.Value,
		})
	}
}

func (s Service) DeleteSession(w http.ResponseWriter) interfaces.Response {
	unsetSessionCookie(w)
	return response_factory.DefaultResponse()
}
