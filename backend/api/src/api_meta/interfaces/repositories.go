package interfaces

import "api_meta/models"

type (
	SessionRepository interface {
		CredentialsExists(accountHash, login, password string) (bool, error)
	}

	RegistrationRepository interface {
		CreateAccount(accountHash string) error
		CreateAccountCredentials(record models.AccountCredentialsRecord) error
	}

	TestCommandKeyValueCreatorRepository interface {
		Create(accountHash string, keyValues models.CommandKeyValue) error
	}

	TestCommandHeadersEditorRepository interface {
		UpdateHeaders(accountHash string, entities []models.UpdateModel) error
		DeleteHeader(accountHash, headerHash string) error
	}

	TestCommandCookiesEditorRepository interface {
		UpdateCookies(accountHash string, entities []models.UpdateModel) error
		DeleteCookie(accountHash, cookieHash string) error
	}

	CRUDRepository interface {
		Create(accountHash string, entity map[string]interface{}) error
		GetAll(accountHash string, dest interface{}) error
		Get(accountHash, entityHash string, dest interface{}) error
		Update(accountHash string, entities []models.UpdateModel) error
		Delete(accountHash, entityHash string) error
	}

	QueryProvider interface {
		CreateQuery() string
		GetAllQuery() string
		GetQuery() string
		UpdateQuery(updateFieldName string) string
		DeleteQuery() string
	}
)
