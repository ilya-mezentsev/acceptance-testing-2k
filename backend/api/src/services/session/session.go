package session

import (
	"api_meta/interfaces"
	"api_meta/models"
	"containers/expirable"
	"db_connector"
	"events/listener"
	"net/http"
	"services/errors"
	"services/plugins/account_credentials"
	"services/plugins/logger"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
	"sync"
	"time"
)

type Service struct {
	sync.Mutex
	logger               logger.CRUDEntityErrorsLogger
	repository           interfaces.SessionRepository
	deletedAccountHashes []expirable.Container
}

func New(repository interfaces.SessionRepository) *Service {
	s := Service{
		repository: repository,
		logger:     logger.CRUDEntityErrorsLogger{EntityName: entityName},
	}

	listener.Get().Subscribe.Admin.DeleteAccount(s.addDeletedAccountHash)
	listener.Get().Subscribe.System.CleanExpiredAccountHashes(s.cleanExpiredDeletedAccounts)

	return &s
}

func (s *Service) CreateSession(w http.ResponseWriter, r *http.Request) interfaces.Response {
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
			AccountHash: accountHash,
		})
	} else {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateSessionCode,
			Description: accountDoesNotExistsError,
		})
	}
}

func (s *Service) GetSession(w http.ResponseWriter, r *http.Request) interfaces.Response {
	cookie, err := getSessionCookie(r)
	if err != nil {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToGetSessionCode,
			Description: sessionCookieNotFoundError,
		})
	} else if s.isAccountHashDeleted(cookie.Value) {
		unsetSessionCookie(w)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToGetSessionCode,
			Description: accountIsDeleted,
		})
	} else {
		return response_factory.SuccessResponse(models.SessionResponse{
			AccountHash: cookie.Value,
		})
	}
}

func (s *Service) isAccountHashDeleted(accountHash string) bool {
	s.Lock()
	defer s.Unlock()

	for _, container := range s.deletedAccountHashes {
		if container.GetValue().(string) == accountHash {
			return true
		}
	}

	return false
}

func (s *Service) DeleteSession(w http.ResponseWriter) interfaces.Response {
	unsetSessionCookie(w)
	return response_factory.DefaultResponse()
}

func (s *Service) addDeletedAccountHash(accountHash string) {
	s.Lock()
	defer s.Unlock()

	s.deletedAccountHashes = append(s.deletedAccountHashes, expirable.Init(accountHash))
}

func (s *Service) cleanExpiredDeletedAccounts(d time.Duration) {
	s.Lock()
	defer s.Unlock()

	for i, hashContainer := range s.deletedAccountHashes {
		if hashContainer.IsExpired(d) {
			s.deletedAccountHashes = append(
				s.deletedAccountHashes[:i],
				s.deletedAccountHashes[i+1:]...,
			)
		}
	}
}
