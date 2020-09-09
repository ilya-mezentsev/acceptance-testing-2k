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

	TestCommandMetaCreatorRepository interface {
		Create(accountHash string, meta models.CommandMeta) error
	}

	TestCommandMetaUpdaterRepository interface {
		UpdateHeadersAndCookies(accountHash string, headers, cookies []models.UpdateModel) error
	}

	TestCommandMetaRepository interface {
		TestCommandMetaCreatorRepository
		TestCommandMetaUpdaterRepository
	}

	TestCommandMetaGetterRepository interface {
		GetAllHeadersAndCookies(accountHash string) (
			headers,
			cookies []models.KeyValueMapping,
			err error,
		)
		GetCommandHeadersAndCookies(accountHash, commandHash string) (
			headers,
			cookies []models.KeyValueMapping,
			err error,
		)
	}

	TestCommandHeadersDeleterRepository interface {
		DeleteHeader(accountHash, headerHash string) error
	}

	TestCommandCookiesDeleterRepository interface {
		DeleteCookie(accountHash, cookieHash string) error
	}

	CreateRepository interface {
		Create(accountHash string, entity map[string]interface{}) error
	}

	GetRepository interface {
		GetAll(accountHash string, dest interface{}) error
		Get(accountHash, entityHash string, dest interface{}) error
	}

	UpdateRepository interface {
		Update(accountHash string, entities []models.UpdateModel) error
	}

	DeleteRepository interface {
		Delete(accountHash, entityHash string) error
	}

	CRUDRepository interface {
		CreateRepository
		GetRepository
		UpdateRepository
		DeleteRepository
	}

	QueryProvider interface {
		CreateQuery() string
		GetAllQuery() string
		GetQuery() string
		UpdateQuery(updateFieldName string) string
		DeleteQuery() string
	}
)
