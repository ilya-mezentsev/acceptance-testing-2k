package db_connector

import (
	"env"
	"events/listener"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"logger"
	"os"
	"path"
	"sync"
	"time"
)

type Connector struct {
	sync.Mutex
	dbRootPath              string
	accountHashToConnection map[string]ConnectionContainer
}

func New(dbRootPath string) *Connector {
	c := &Connector{
		dbRootPath:              dbRootPath,
		accountHashToConnection: map[string]ConnectionContainer{},
	}

	listener.Get().Subscribe.System.CleanExpiredDBConnections(
		c.cleanExpiredConnections,
	)
	listener.Get().Subscribe.Admin.DeleteAccount(
		c.closeAccountConnection,
	)

	return c
}

func (c *Connector) Connect(accountHash string) (*sqlx.DB, error) {
	c.Lock()
	defer c.Unlock()

	_, found := c.accountHashToConnection[accountHash]
	if !found {
		connection, err := c.connect(accountHash)
		if err != nil {
			return nil, err
		}

		c.accountHashToConnection[accountHash] = ConnectionContainer{
			db:      c.configureConnection(connection),
			created: time.Now(),
		}
	}

	return c.accountHashToConnection[accountHash].GetConnection(), nil
}

func (c *Connector) connect(accountHash string) (*sqlx.DB, error) {
	dbFilePath := path.Join(c.dbRootPath, accountHash, env.DBFilename)
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

func (c *Connector) configureConnection(connection *sqlx.DB) *sqlx.DB {
	connection.SetMaxOpenConns(50)
	connection.SetMaxIdleConns(50)
	connection.SetConnMaxLifetime(time.Hour)

	return connection
}

func (c *Connector) cleanExpiredConnections(d time.Duration) {
	c.Lock()
	defer c.Unlock()

	for accountHash, connection := range c.accountHashToConnection {
		if connection.IsExpired(d) {
			connection.Close()

			delete(c.accountHashToConnection, accountHash)
		}
	}
}

func (c *Connector) closeAccountConnection(accountHash string) {
	c.Lock()
	defer c.Unlock()

	connection, ok := c.accountHashToConnection[accountHash]
	if !ok {
		return
	}

	connection.Close()
	delete(c.accountHashToConnection, accountHash)
}
