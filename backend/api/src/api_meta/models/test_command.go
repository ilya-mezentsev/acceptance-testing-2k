package models

import "api_meta/types"

type (
	CommandSettings struct {
		Name               string `db:"name"`
		Hash               string `db:"hash"`
		ObjectName         string `db:"object_name"`
		Method             string `db:"method"`
		BaseURL            string `db:"base_url"`
		Endpoint           string `db:"endpoint"`
		PassArgumentsInURL bool   `db:"pass_arguments_in_url"`
	}

	TestCommandRecord struct {
		CommandSettings
		Headers string `db:"command_headers"`
		Cookies string `db:"command_cookies"`
	}

	TestCommandRequest struct {
		CommandSettings
		Headers types.Mapping
		Cookies types.Mapping
	}
)
