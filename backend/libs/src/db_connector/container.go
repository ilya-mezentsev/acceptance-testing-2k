package db_connector

import (
	"github.com/jmoiron/sqlx"
	"time"
)

var connectionCacheLifetime = connectionLifetime / 2

type ConnectionContainer struct {
	db      *sqlx.DB
	created time.Time
}

func (cc ConnectionContainer) GetConnection() *sqlx.DB {
	return cc.db
}

func (cc ConnectionContainer) Close() {
	_ = cc.db.Close()
}

func (cc ConnectionContainer) IsExpired() bool {
	return cc.created.Add(connectionCacheLifetime).Before(time.Now())
}
