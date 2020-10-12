package db_connector

import (
	"github.com/jmoiron/sqlx"
	"time"
)

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

func (cc ConnectionContainer) IsExpired(d time.Duration) bool {
	return cc.created.Add(d).Before(time.Now())
}
