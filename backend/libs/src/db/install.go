package db

import "github.com/jmoiron/sqlx"

const createTablesQuery = `
	CREATE TABLE IF NOT EXISTS account_credentials(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		login VARCHAR(64) NOT NULL UNIQUE,
		password VARCHAR(32) NOT NULL,
		verified BOOLEAN NOT NULL DEFAULT 0 CHECK (verified IN (0,1))
	);
	CREATE TABLE IF NOT EXISTS objects(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		hash VARCHAR(32) NOT NULL UNIQUE
	);
	CREATE TABLE IF NOT EXISTS commands(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		object_name TEXT REFERENCES objects(name),
		name TEXT NOT NULL,
		hash VARCHAR(32) NOT NULL UNIQUE,
		UNIQUE(object_name, name)
	);
	CREATE TABLE IF NOT EXISTS commands_settings(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		method TEXT NOT NULL,
		base_url TEXT NOT NULL,
		endpoint TEXT DEFAULT '',
		pass_arguments_in_url BOOLEAN NOT NULL CHECK (pass_arguments_in_url IN (0,1)),
		command_hash VARCHAR(32) REFERENCES commands(hash)
	);
	CREATE TABLE IF NOT EXISTS commands_headers(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT NOT NULL,
		value TEXT NOT NULL,
		command_hash VARCHAR(32) REFERENCES commands(hash)
	);
	CREATE TABLE IF NOT EXISTS commands_cookies(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT NOT NULL,
		value TEXT NOT NULL,
		command_hash VARCHAR(32) REFERENCES commands(hash)
	);`

func Install(db *sqlx.DB) error {
	_, err := db.Exec(createTablesQuery)
	return err
}
