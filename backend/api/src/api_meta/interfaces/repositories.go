package interfaces

import "api_meta/models"

type (
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
