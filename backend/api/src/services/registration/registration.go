package registration

import (
	"api_meta/interfaces"
	"api_meta/models"
	"api_meta/types"
	db2 "db"
	"db_connector"
	"env"
	"errors"
	"io"
	"os"
	"path"
	servicesErrors "services/errors"
	"services/plugins/account_credentials"
	"services/plugins/logger"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

type Service struct {
	filesRootPath string
	logger        logger.CRUDEntityErrorsLogger
	repository    interfaces.RegistrationRepository
}

func New(repository interfaces.RegistrationRepository, filesRootPath string) Service {
	return Service{
		filesRootPath: filesRootPath,
		repository:    repository,
		logger:        logger.CRUDEntityErrorsLogger{EntityName: entityName},
	}
}

func (s Service) Register(request io.ReadCloser) interfaces.Response {
	var registrationRequest models.RegistrationRequest
	err := request_decoder.Decode(request, &registrationRequest)
	if err != nil {
		s.logger.LogCreateEntityDecodeError(err)

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToRegisterAccountCode,
			Description: servicesErrors.DecodingRequestError,
		})
	}

	if !validation.IsRegularName(registrationRequest.Login) {
		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToRegisterAccountCode,
			Description: servicesErrors.InvalidRequestError,
		})
	}

	accountHash := account_credentials.GenerateAccountHash(
		registrationRequest.Login,
		registrationRequest.Password,
	)
	err = s.repository.CreateAccount(accountHash)
	if err != nil {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"login":        registrationRequest.Login,
			"account_hash": accountHash,
		})
		var description string
		if errors.As(err, &types.AccountHashAlreadyExists{}) {
			description = loginAlreadyExistsError
		} else {
			description = servicesErrors.RepositoryError
		}

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToRegisterAccountCode,
			Description: description,
		})
	}

	err = s.installDB(accountHash)
	if err != nil {
		s.logger.LogCreateEntityRepositoryOrOSError(err, map[string]interface{}{
			"account_hash": accountHash,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToRegisterAccountCode,
			Description: servicesErrors.UnknownError,
		})
	}

	err = s.repository.CreateAccountCredentials(models.AccountCredentialsRecord{
		AccountHash: accountHash,
		Login:       registrationRequest.Login,
		Password: account_credentials.GenerateAccountPassword(
			registrationRequest.Login,
			registrationRequest.Password,
		),
	})
	if err != nil {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"login":        registrationRequest.Login,
			"account_hash": accountHash,
		})

		return response_factory.ErrorResponse(servicesErrors.ServiceError{
			Code:        unableToRegisterAccountCode,
			Description: servicesErrors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) installDB(accountHash string) error {
	err := os.MkdirAll(path.Join(s.filesRootPath, accountHash), 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(s.filesRootPath, accountHash, env.DBFilename))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	connector := db_connector.New(s.filesRootPath)
	db, err := connector.Connect(accountHash)
	if err != nil {
		return err
	}

	return db2.Install(db)
}
