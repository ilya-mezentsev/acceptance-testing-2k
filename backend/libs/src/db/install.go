package db

import "github.com/jmoiron/sqlx"

const createTablesQuery = `
	CREATE TABLE IF NOT EXISTS account_credentials(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		login VARCHAR(64) NOT NULL UNIQUE,
		password VARCHAR(32) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS objects(
		name TEXT NOT NULL UNIQUE,
		hash VARCHAR(32) NOT NULL PRIMARY KEY
	);
	CREATE TABLE IF NOT EXISTS commands(
		name TEXT NOT NULL,
		hash VARCHAR(32) NOT NULL PRIMARY KEY,
		object_hash VARCHAR(32),
		UNIQUE(object_hash, name),
		FOREIGN KEY(object_hash) REFERENCES objects(hash) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS commands_settings(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		method TEXT NOT NULL,
		base_url TEXT NOT NULL,
		endpoint TEXT DEFAULT '',
		timeout INTEGER NOT NULL,
		pass_arguments_in_url BOOLEAN NOT NULL CHECK (pass_arguments_in_url IN (0,1)),
		command_hash VARCHAR(32),
		FOREIGN KEY(command_hash) REFERENCES commands(hash) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS commands_headers(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT NOT NULL,
		value TEXT NOT NULL,
		hash VARCHAR(32) NOT NULL UNIQUE,
		command_hash VARCHAR(32),
		FOREIGN KEY(command_hash) REFERENCES commands(hash) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS commands_cookies(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT NOT NULL,
		value TEXT NOT NULL,
		hash VARCHAR(32) NOT NULL UNIQUE,
		command_hash VARCHAR(32),
		FOREIGN KEY(command_hash) REFERENCES commands(hash) ON DELETE CASCADE
	);`

func Install(db *sqlx.DB) error {
	_, err := db.Exec(createTablesQuery)
	return err
}
