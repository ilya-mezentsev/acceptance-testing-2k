package db_connector

import (
	"containers/expirable"
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
	accountHashToConnection map[string]expirable.Container
}

func New(dbRootPath string) *Connector {
	c := &Connector{
		dbRootPath:              dbRootPath,
		accountHashToConnection: map[string]expirable.Container{},
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

		c.accountHashToConnection[accountHash] = expirable.Init(
			c.configureConnection(connection),
		)
	}

	return c.accountHashToConnection[accountHash].GetValue().(*sqlx.DB), nil
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

	for accountHash, container := range c.accountHashToConnection {
		if container.IsExpired(d) {
			_ = container.GetValue().(*sqlx.DB).Close()

			delete(c.accountHashToConnection, accountHash)
		}
	}
}

func (c *Connector) closeAccountConnection(accountHash string) {
	c.Lock()
	defer c.Unlock()

	container, ok := c.accountHashToConnection[accountHash]
	if !ok {
		return
	}

	_ = container.GetValue().(*sqlx.DB).Close()
	delete(c.accountHashToConnection, accountHash)
}
