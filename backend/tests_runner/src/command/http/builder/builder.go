package builder

import (
	"command/http/command"
	"command/http/errors"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"test_runner_meta/interfaces"
)

const (
	getCommandInfoQuery = `
	SELECT name, hash FROM commands
	WHERE name = $1 AND object_hash = (SELECT hash FROM objects WHERE name = $2)`
	getCommandSettingsQuery = `
	SELECT method, base_url, endpoint, pass_arguments_in_url
	FROM commands_settings
	WHERE command_hash = ?`
	getCommandHeadersQuery = `SELECT key, value FROM commands_headers WHERE command_hash = ?`
	getCommandCookiesQuery = `SELECT key, value FROM commands_cookies WHERE command_hash = ?`
)

type Builder struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) interfaces.CommandBuilder {
	return Builder{db}
}

func (b Builder) Build(objectName, commandName string) (interfaces.Command, error) {
	commandInfo, err := b.getCommandInfo(objectName, commandName)
	if err != nil {
		return nil, err
	}

	commandSettings, err := b.getCommandSettings(commandInfo.Hash)
	if err != nil {
		return nil, err
	}

	return command.New(commandSettings), nil
}

func (b Builder) getCommandInfo(objectName, commandName string) (CommandInfo, error) {
	var info CommandInfo
	err := b.db.Get(&info, getCommandInfoQuery, commandName, objectName)
	if err != nil && err == sql.ErrNoRows {
		err = errors.CommandNotFound
	}

	return info, err
}

func (b Builder) getCommandSettings(commandHash string) (settings, error) {
	var s settings
	err := b.db.Get(&s, getCommandSettingsQuery, commandHash)
	if err != nil {
		return settings{}, err
	}

	err = b.db.Select(&s.Headers, getCommandHeadersQuery, commandHash)
	if err != nil {
		return settings{}, err
	}

	err = b.db.Select(&s.Cookies, getCommandCookiesQuery, commandHash)
	if err != nil {
		return settings{}, err
	}

	return s, nil
}
