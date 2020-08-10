package crud

import (
	"api_meta/interfaces"
	"api_meta/models"
	"db_connector"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	connector     db_connector.Connector
	queryProvider interfaces.QueryProvider
}

func New(
	connector db_connector.Connector,
	queryProvider interfaces.QueryProvider,
) interfaces.CRUDRepository {
	return Repository{connector: connector, queryProvider: queryProvider}
}

func (r Repository) Create(accountHash string, entity map[string]interface{}) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	_, err = db.NamedExec(r.queryProvider.CreateQuery(), entity)
	return err
}

func (r Repository) GetAll(accountHash string, dest interface{}) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	return db.Select(dest, r.queryProvider.GetAllQuery())
}

func (r Repository) Get(accountHash, entityHash string, dest interface{}) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	return db.Get(dest, r.queryProvider.GetQuery(), entityHash)
}

func (r Repository) Update(accountHash string, entities []models.UpdateModel) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	for _, entity := range entities {
		err = r.execTransaction(tx, entity)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r Repository) execTransaction(tx *sqlx.Tx, entity models.UpdateModel) error {
	_, err := tx.NamedExec(r.queryProvider.UpdateQuery(entity.FieldName), entity)

	return err
}

func (r Repository) Delete(accountHash, entityHash string) error {
	db, err := r.connector.Connect(accountHash)
	if err != nil {
		return err
	}

	_, err = db.Exec(r.queryProvider.DeleteQuery(), entityHash)
	return err
}
