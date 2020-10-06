package test_command_meta

import (
	"api_meta/models"
	"db_connector"
	"fmt"
)

const (
	getAllHeadersQuery     = `SELECT key, value, hash, command_hash FROM commands_headers`
	getAllCookiesQuery     = `SELECT key, value, hash, command_hash FROM commands_cookies`
	getCommandHeadersQuery = `
	SELECT key, value, hash, command_hash
	FROM commands_headers WHERE command_hash = ?`
	getCommandCookiesQuery = `
	SELECT key, value, hash, command_hash
	FROM commands_cookies WHERE command_hash = ?`
	addHeaderQuery = `
	INSERT OR REPLACE INTO commands_headers(key, value, hash, command_hash)
	VALUES(:key, :value, :hash, :command_hash)`
	addCookieQuery = `
	INSERT OR REPLACE INTO commands_cookies(key, value, hash, command_hash)
	VALUES(:key, :value, :hash, :command_hash)`
	updateHeaderQuery = `UPDATE commands_headers SET %s = :new_value WHERE hash = :hash`
	updateCookieQuery = `UPDATE commands_cookies SET %s = :new_value WHERE hash = :hash`
	deleteHeaderQuery = `DELETE FROM commands_headers WHERE hash = ?`
	deleteCookieQuery = `DELETE FROM commands_cookies WHERE hash = ?`
)

type Repository struct {
	connector *db_connector.Connector
}

type queryToData map[string]map[string]interface{}

func New(connector *db_connector.Connector) Repository {
	return Repository{connector}
}

func (r Repository) GetAllHeadersAndCookies(accountHash string) (
	headers,
	cookies []models.KeyValueMapping,
	err error,
) {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return nil, nil, err
	}

	err = db.Select(&headers, getAllHeadersQuery)
	if err != nil {
		return nil, nil, err
	}

	err = db.Select(&cookies, getAllCookiesQuery)
	if err != nil {
		return nil, nil, err
	}

	return headers, cookies, nil
}

func (r Repository) GetCommandHeadersAndCookies(accountHash, commandHash string) (
	headers,
	cookies []models.KeyValueMapping,
	err error,
) {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return nil, nil, err
	}

	err = db.Select(&headers, getCommandHeadersQuery, commandHash)
	if err != nil {
		return nil, nil, err
	}

	err = db.Select(&cookies, getCommandCookiesQuery, commandHash)
	if err != nil {
		return nil, nil, err
	}

	return headers, cookies, nil
}

func (r Repository) Create(accountHash string, meta models.CommandMeta) error {
	return r.performTransactions(
		accountHash,
		append(
			r.prepareInsertMap(addHeaderQuery, meta.Headers),
			r.prepareInsertMap(addCookieQuery, meta.Cookies)...,
		),
	)
}

func (r Repository) prepareInsertMap(
	query string,
	mappings []models.KeyValueMapping,
) []queryToData {
	var insert []queryToData
	for _, mapping := range mappings {
		insert = append(insert, queryToData{
			query: map[string]interface{}{
				"key":          mapping.Key,
				"value":        mapping.Value,
				"hash":         mapping.Hash,
				"command_hash": mapping.CommandHash,
			},
		})
	}

	return insert
}

func (r Repository) UpdateHeadersAndCookies(
	accountHash string,
	headers,
	cookies []models.UpdateModel,
) error {
	return r.performTransactions(
		accountHash,
		append(
			r.prepareUpdateMap(updateHeaderQuery, headers),
			r.prepareUpdateMap(updateCookieQuery, cookies)...,
		),
	)
}

func (r Repository) prepareUpdateMap(
	query string,
	updateModels []models.UpdateModel,
) []queryToData {
	var d []queryToData
	for _, updateModel := range updateModels {
		d = append(d, queryToData{
			fmt.Sprintf(query, updateModel.FieldName): {
				"new_value": updateModel.NewValue,
				"hash":      updateModel.Hash,
			},
		})
	}

	return d
}

func (r Repository) performTransactions(accountHash string, data []queryToData) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	for _, queryToData := range data {
		for query, data := range queryToData {
			_, err = tx.NamedExec(query, data)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}

func (r Repository) DeleteCookie(accountHash, cookieHash string) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	_, err = db.Exec(deleteCookieQuery, cookieHash)
	return err
}

func (r Repository) DeleteHeader(accountHash, headerHash string) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	_, err = db.Exec(deleteHeaderQuery, headerHash)
	return err
}
