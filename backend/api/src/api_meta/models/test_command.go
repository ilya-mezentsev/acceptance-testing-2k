package models

import "api_meta/types"

type (
	CommandSettings struct {
		Name               string `json:"name" db:"name" validation:"regular-name"`
		Hash               string `json:"hash" db:"hash" validation:"md5-hash"`
		ObjectName         string `json:"object_name" db:"object_name" validation:"regular-name"`
		Method             string `json:"method" db:"method" validation:"meaning-http-method"`
		BaseURL            string `json:"base_url" db:"base_url" validation:"base-url"`
		Endpoint           string `json:"endpoint" db:"endpoint" validation:"endpoint"`
		PassArgumentsInURL bool   `json:"pass_arguments_in_url" db:"pass_arguments_in_url"`
	}

	TestCommandRecord struct {
		CommandSettings
		Headers string `db:"command_headers" validation:"key-value-mapping; no-validate-empty-string"`
		Cookies string `db:"command_cookies" validation:"key-value-mapping; no-validate-empty-string"`
	}

	TestCommandRequest struct {
		CommandSettings
		Headers types.Mapping `json:"headers"`
		Cookies types.Mapping `json:"cookies"`
	}

	CreateTestCommandRequest struct {
		AccountHash string             `json:"account_hash"`
		TestCommand TestCommandRequest `json:"test_command"`
	}
)
