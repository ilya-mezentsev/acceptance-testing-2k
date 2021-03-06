package models

type (
	CommandSettings struct {
		Name               string `json:"name" db:"name" validation:"regular-name"`
		Hash               string `json:"hash" db:"hash" validation:"md5-hash"`
		ObjectHash         string `json:"object_hash" db:"object_hash" validation:"md5-hash"`
		Method             string `json:"method" db:"method" validation:"meaning-http-method"`
		BaseURL            string `json:"base_url" db:"base_url" validation:"base-url"`
		Endpoint           string `json:"endpoint" db:"endpoint" validation:"endpoint"`
		Timeout            int    `json:"timeout" db:"timeout" range:"1,15"`
		PassArgumentsInURL bool   `json:"pass_arguments_in_url" db:"pass_arguments_in_url"`
	}

	GetCommandResponse struct {
		CommandSettings
		CommandMeta
	}

	CreateTestCommandRequest struct {
		CommandSettings CommandSettings `json:"command_settings"`
	}

	CreatedTestCommandResponse struct {
		CommandHash string `json:"command_hash"`
	}

	UpdateTestCommandRequest struct {
		ExistsCommand  CommandSettings `json:"exists_command"`
		UpdatedCommand CommandSettings `json:"updated_command"`
	}

	UpdateTarget struct {
		Hash string `json:"hash" validation:"md5-hash"`
	}

	massUpdateRequest struct {
		CommandHashes []UpdateTarget `json:"command_hashes"`
	}

	MassBaseURLsUpdateRequest struct {
		massUpdateRequest
		BaseURL string `json:"base_url" validation:"base-url"`
	}

	MassTimeoutsUpdateRequest struct {
		massUpdateRequest
		Timeout int `json:"timeout" range:"1,15"`
	}

	MassCookiesCreateRequest struct {
		massUpdateRequest
		Cookies []KeyValueMapping `json:"cookies"`
	}

	MassHeadersCreateRequest struct {
		massUpdateRequest
		Headers []KeyValueMapping `json:"headers"`
	}

	KeyValueMapping struct {
		Hash        string `db:"hash" json:"hash" validation:"md5-hash"`
		Key         string `db:"key" json:"key" validation:"regular-name"`
		Value       string `db:"value" json:"value"`
		CommandHash string `db:"command_hash" json:"command_hash" validation:"md5-hash"`
	}

	CommandMeta struct {
		Headers []KeyValueMapping `json:"headers"`
		Cookies []KeyValueMapping `json:"cookies"`
	}

	CreateMetaRequest struct {
		CommandHash string      `json:"command_hash" validation:"md5-hash"`
		CommandMeta CommandMeta `json:"command_meta"`
	}
)
