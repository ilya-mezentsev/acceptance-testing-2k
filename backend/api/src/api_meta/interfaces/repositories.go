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

	TestCommandMetaRepository interface {
		Create(accountHash string, meta models.CommandMeta) error
		UpdateHeadersAndCookies(accountHash string, headers, cookies []models.UpdateModel) error
	}

	TestCommandHeadersDeleterRepository interface {
		DeleteHeader(accountHash, headerHash string) error
	}

	TestCommandCookiesDeleterRepository interface {
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
