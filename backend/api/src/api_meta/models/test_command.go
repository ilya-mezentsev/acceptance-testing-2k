package models

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
		Headers string `db:"command_headers" json:"headers"`
		Cookies string `db:"command_cookies" json:"cookies"`
	}

	CreateTestCommandRequest struct {
		AccountHash     string          `json:"account_hash"`
		CommandSettings CommandSettings `json:"command_settings"`
	}

	CreatedTestCommandResponse struct {
		CommandHash string `json:"command_hash"`
	}

	KeyValueMapping struct {
		Hash        string `db:"hash" validation:"md5-hash"`
		Key         string `db:"key" validation:"regular-name"`
		Value       string `db:"value"`
		CommandHash string `db:"command_hash" validation:"md5-hash"`
	}

	CommandKeyValue struct {
		Headers []KeyValueMapping `json:"headers"`
		Cookies []KeyValueMapping `json:"cookies"`
	}

	CreateMetaRequest struct {
		AccountHash string          `json:"account_hash" validation:"md5-hash"`
		CommandHash string          `json:"command_hash" validation:"md5-hash"`
		CommandMeta CommandKeyValue `json:"command_meta"`
	}
)
