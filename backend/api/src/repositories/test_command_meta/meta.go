package test_command_meta

import (
	"api_meta/models"
	"db_connector"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	addHeaderQuery = `
	INSERT INTO commands_headers(key, value, hash, command_hash)
	VALUES(:key, :value, :hash, :command_hash)`
	addCookieQuery = `
	INSERT INTO commands_cookies(key, value, hash, command_hash)
	VALUES(:key, :value, :hash, :command_hash)`
	updateHeaderQuery = `UPDATE commands_headers SET %s = :new_value WHERE hash = :hash`
	updateCookieQuery = `UPDATE commands_cookies SET %s = :new_value WHERE hash = :hash`
	deleteHeaderQuery = `DELETE FROM commands_headers WHERE hash = ?`
	deleteCookieQuery = `DELETE FROM commands_cookies WHERE hash = ?`
)

type Repository struct {
	connector db_connector.Connector
}

func New(connector db_connector.Connector) Repository {
	return Repository{connector}
}

func (r Repository) Create(accountHash string, keyValues models.CommandKeyValue) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	err = r.applyKeyValueInsert(tx, addHeaderQuery, keyValues.Headers)
	if err != nil {
		return err
	}

	err = r.applyKeyValueInsert(tx, addCookieQuery, keyValues.Cookies)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r Repository) applyKeyValueInsert(
	tx *sqlx.Tx,
	query string,
	mapping []models.KeyValueMapping,
) error {
	for _, keyValue := range mapping {
		_, err := tx.NamedExec(query, keyValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r Repository) UpdateHeaders(accountHash string, entities []models.UpdateModel) error {
	return r.performUpdate(accountHash, updateHeaderQuery, entities)
}

func (r Repository) UpdateCookies(accountHash string, entities []models.UpdateModel) error {
	return r.performUpdate(accountHash, updateCookieQuery, entities)
}

func (r Repository) performUpdate(
	accountHash,
	query string,
	entities []models.UpdateModel,
) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	for _, entity := range entities {
		_, err = tx.NamedExec(fmt.Sprintf(query, entity.FieldName), entity)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r Repository) DeleteHeader(accountHash, headerHash string) error {
	return r.performDelete(accountHash, deleteHeaderQuery, headerHash)
}

func (r Repository) DeleteCookie(accountHash, cookieHash string) error {
	return r.performDelete(accountHash, deleteCookieQuery, cookieHash)
}

func (r Repository) performDelete(accountHash, query, entityHash string) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, entityHash)
	return err
}
