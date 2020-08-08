package db_connector

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"plugins/logger"
)

const dbFilename = "db.db"

type Connector struct {
	dbRootPath              string
	accountHashToConnection map[string]*sqlx.DB
}

func New(dbRootPath string) Connector {
	return Connector{dbRootPath: dbRootPath, accountHashToConnection: map[string]*sqlx.DB{}}
}

func (c *Connector) Connect(accountHash string) (*sqlx.DB, error) {
	_, found := c.accountHashToConnection[accountHash]
	if !found {
		connection, err := c.connect(accountHash)
		if err != nil {
			return nil, err
		}

		c.accountHashToConnection[accountHash] = connection
	}

	return c.accountHashToConnection[accountHash], nil
}

func (c Connector) connect(accountHash string) (*sqlx.DB, error) {
	dbFilePath := path.Join(c.dbRootPath, accountHash, dbFilename)
	if _, err := os.Stat(dbFilePath); err == nil {
		return sqlx.Open("sqlite3", dbFilePath)
	} else if os.IsNotExist(err) {
		return nil, DBFileNotFound
	} else {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unexpected error: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"db_root_path": c.dbRootPath,
				"account_hash": accountHash,
			},
		}, logger.Error)

		return nil, UnknownError
	}
}
