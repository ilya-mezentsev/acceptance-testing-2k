package query_providers

import (
	"fmt"
	"strings"
)

var updateTargetToTableNameAndHashColumnName = map[string][2]string{
	"command":         {"commands", "hash"},
	"command_setting": {"commands_settings", "command_hash"},
	"command_header":  {"commands_headers", "command_hash"},
	"command_cookie":  {"commands_cookies", "command_hash"},
}

type TestCommandQueryProvider struct {
}

func (p TestCommandQueryProvider) CreateQuery() string {
	return `
	INSERT INTO commands(name, hash, object_name)
	VALUES(:name, :hash, :object_name);

	INSERT INTO commands_settings(method, base_url, endpoint, pass_arguments_in_url, command_hash)
	VALUES(:method, :base_url, :endpoint, :pass_arguments_in_url, :hash);

	WITH RECURSIVE split(key, value, str) AS (
    SELECT '', '', :command_headers||';'
    UNION
  	SELECT
      substr(str, 0, instr(str, '=')),
      substr(str, instr(str, '=')+1),
      substr(str, instr(str, ';')+1)
    FROM split WHERE str != ''
	)
		INSERT INTO commands_headers(key, value, command_hash)
		SELECT key, substr(value, 0, instr(value, ';')), :hash
		FROM split WHERE key != '' and value != '';

	WITH RECURSIVE split(key, value, str) AS (
    SELECT '', '', :command_cookies||';'
    UNION
  	SELECT
      substr(str, 0, instr(str, '=')),
      substr(str, instr(str, '=')+1),
      substr(str, instr(str, ';')+1)
    FROM split WHERE str != ''
	)
		INSERT INTO commands_cookies(key, value, command_hash)
		SELECT key, substr(value, 0, instr(value, ';')), :hash
		FROM split WHERE key != '' and value != '';`
}

func (p TestCommandQueryProvider) GetAllQuery() string {
	return `
	SELECT c.name, c.object_name, c.hash, cs.method, cs.base_url, cs.endpoint, cs.pass_arguments_in_url,
	replace(
		group_concat(ifnull(ch.key, ' ') || '=' || ifnull(ch.value, ' '), ';'), ' = ', ''
	) as command_headers,
	replace(
		group_concat(ifnull(cc.key, ' ') || '=' || ifnull(cc.value, ' '), ';'), ' = ', ''
	) as command_cookies
	FROM commands c
	LEFT JOIN commands_settings cs ON cs.command_hash = c.hash
	LEFT JOIN commands_headers ch ON ch.command_hash = c.hash
	LEFT JOIN commands_cookies cc ON cc.command_hash = c.hash
	GROUP BY c.id`
}

func (p TestCommandQueryProvider) GetQuery() string {
	return `
	SELECT c.name, c.object_name, c.hash, cs.method, cs.base_url, cs.endpoint, cs.pass_arguments_in_url,
	replace(
		group_concat(ifnull(ch.key, ' ') || '=' || ifnull(ch.value, ' '), ';'), ' = ', ''
	) as command_headers,
	replace(
		group_concat(ifnull(cc.key, ' ') || '=' || ifnull(cc.value, ' '), ';'), ' = ', ''
	) as command_cookies
	FROM commands c
	LEFT JOIN commands_settings cs ON cs.command_hash = c.hash
	LEFT JOIN commands_headers ch ON ch.command_hash = c.hash
	LEFT JOIN commands_cookies cc ON cc.command_hash = c.hash
	WHERE c.hash = $1
	GROUP BY c.id`
}

func (p TestCommandQueryProvider) UpdateQuery(updateFieldName string) string {
	components := strings.Split(updateFieldName, ":")
	if len(components) < 2 {
		panic("invalid update field name format")
	}
	updateTarget, fieldName := components[0], components[1]
	varSlice := updateTargetToTableNameAndHashColumnName[updateTarget]
	tableName, hashColumnName := varSlice[0], varSlice[1]

	return fmt.Sprintf(
		"UPDATE %s SET %s = :new_value WHERE %s = :hash AND id = (SELECT MIN(id) FROM %s)",
		tableName, fieldName, hashColumnName, tableName,
	)
}

func (p TestCommandQueryProvider) DeleteQuery() string {
	return `DELETE FROM commands WHERE hash = ?`
}
