package services

import (
	"api_meta/models"
	"api_meta/types"
)

type TestCommandsRepositoryMock struct {
	Commands map[string][]models.TestCommandRecord
}

func (m *TestCommandsRepositoryMock) Reset() {
	m.Commands = map[string][]models.TestCommandRecord{
		PredefinedAccountHash: {PredefinedTestCommand1, PredefinedTestCommand2},
	}
}

func (m *TestCommandsRepositoryMock) Create(accountHash string, entity map[string]interface{}) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for _, command := range m.Commands[accountHash] {
		if command.Name == entity["name"].(string) {
			return types.UniqueEntityAlreadyExists{}
		}
	}

	m.Commands[accountHash] = append(m.Commands[accountHash], models.TestCommandRecord{
		CommandSettings: models.CommandSettings{
			Name:               entity["name"].(string),
			Hash:               entity["hash"].(string),
			ObjectName:         entity["object_name"].(string),
			Method:             entity["method"].(string),
			BaseURL:            entity["base_url"].(string),
			Endpoint:           entity["endpoint"].(string),
			PassArgumentsInURL: entity["pass_arguments_in_url"].(bool),
		},
	})

	return nil
}

func (m *TestCommandsRepositoryMock) GetAll(accountHash string, dest interface{}) error {
	if accountHash == BadAccountHash {
		return someError
	}

	destPtr := dest.(*[]models.TestCommandRecord)
	*destPtr = append(*destPtr, m.Commands[accountHash]...)
	return nil
}

func (m *TestCommandsRepositoryMock) Get(accountHash, entityHash string, dest interface{}) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for _, command := range m.Commands[accountHash] {
		if command.Hash == entityHash {
			dest.(*models.TestCommandRecord).Name = command.Name
			dest.(*models.TestCommandRecord).Hash = command.Hash
			dest.(*models.TestCommandRecord).ObjectName = command.ObjectName
			dest.(*models.TestCommandRecord).Method = command.Method
			dest.(*models.TestCommandRecord).BaseURL = command.BaseURL
			dest.(*models.TestCommandRecord).Endpoint = command.Endpoint
			dest.(*models.TestCommandRecord).PassArgumentsInURL = command.PassArgumentsInURL
			dest.(*models.TestCommandRecord).Headers = command.Headers
			dest.(*models.TestCommandRecord).Cookies = command.Cookies
			break
		}
	}

	return nil
}

func (m *TestCommandsRepositoryMock) Update(accountHash string, entities []models.UpdateModel) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for commandIndex, command := range m.Commands[accountHash] {
		for _, entity := range entities {
			if command.Hash == entity.Hash {
				switch entity.FieldName {
				case "command_setting:name":
					m.Commands[accountHash][commandIndex].Name = entity.NewValue.(string)
				case "hash":
					m.Commands[accountHash][commandIndex].Hash = entity.NewValue.(string)
				case "command:object_name":
					m.Commands[accountHash][commandIndex].ObjectName = entity.NewValue.(string)
				case "method":
					m.Commands[accountHash][commandIndex].Method = entity.NewValue.(string)
				case "base_url":
					m.Commands[accountHash][commandIndex].BaseURL = entity.NewValue.(string)
				case "endpoint":
					m.Commands[accountHash][commandIndex].Endpoint = entity.NewValue.(string)
				case "pass_arguments_in_url":
					m.Commands[accountHash][commandIndex].PassArgumentsInURL = entity.NewValue.(bool)
				}
			}
		}
	}

	return nil
}

func (m *TestCommandsRepositoryMock) Delete(accountHash, entityHash string) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for commandIndex, command := range m.Commands[accountHash] {
		if command.Hash == entityHash {
			m.Commands[accountHash] = append(
				m.Commands[accountHash][:commandIndex],
				m.Commands[accountHash][commandIndex+1:]...,
			)
		}
	}

	return nil
}

type TestCommandMetaRepositoryMock struct {
	Meta map[string][]models.CommandMeta
}

func (m *TestCommandMetaRepositoryMock) Reset() {
	m.Meta = map[string][]models.CommandMeta{
		PredefinedAccountHash: {
			{
				Headers: []models.KeyValueMapping{PredefinedHeader1, PredefinedHeader2},
				Cookies: []models.KeyValueMapping{PredefinedCookie1, PredefinedCookie2},
			},
		},
	}
}

func (m *TestCommandMetaRepositoryMock) Create(
	accountHash string,
	meta models.CommandMeta,
) error {
	if accountHash == BadAccountHash {
		return someError
	}

	m.Meta[accountHash] = append(m.Meta[accountHash], meta)
	return nil
}

func (m *TestCommandMetaRepositoryMock) UpdateHeadersAndCookies(
	accountHash string,
	headers,
	cookies []models.UpdateModel,
) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for _, entity := range append(headers, cookies...) {
		for metaIndex, meta := range m.Meta[accountHash] {
			for headerIndex, header := range meta.Headers {
				if header.Hash != entity.Hash {
					continue
				}

				switch entity.FieldName {
				case "key":
					m.Meta[accountHash][metaIndex].Headers[headerIndex].Key = entity.NewValue.(string)
				case "value":
					m.Meta[accountHash][metaIndex].Headers[headerIndex].Value = entity.NewValue.(string)
				}
			}

			for cookieIndex, cookie := range meta.Cookies {
				if cookie.Hash != entity.Hash {
					continue
				}

				switch entity.FieldName {
				case "key":
					m.Meta[accountHash][metaIndex].Cookies[cookieIndex].Key = entity.NewValue.(string)
				case "value":
					m.Meta[accountHash][metaIndex].Cookies[cookieIndex].Value = entity.NewValue.(string)
				}
			}
		}
	}

	return nil
}

type TestCommandHeadersDeleterRepositoryMock struct {
	Headers map[string][]models.KeyValueMapping
}

func (m *TestCommandHeadersDeleterRepositoryMock) Reset() {
	m.Headers = map[string][]models.KeyValueMapping{
		PredefinedAccountHash: {PredefinedHeader1, PredefinedHeader2},
	}
}

func (m *TestCommandHeadersDeleterRepositoryMock) DeleteHeader(
	accountHash,
	headerHash string,
) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for i, header := range m.Headers[accountHash] {
		if header.Hash == headerHash {
			m.Headers[accountHash] = append(
				m.Headers[accountHash][:i],
				m.Headers[accountHash][i+1:]...,
			)
			break
		}
	}

	return nil
}

type TestCommandCookiesDeleterRepositoryMock struct {
	Cookies map[string][]models.KeyValueMapping
}

func (m *TestCommandCookiesDeleterRepositoryMock) Reset() {
	m.Cookies = map[string][]models.KeyValueMapping{
		PredefinedAccountHash: {PredefinedCookie1, PredefinedCookie2},
	}
}

func (m *TestCommandCookiesDeleterRepositoryMock) DeleteCookie(
	accountHash,
	cookieHash string,
) error {
	if accountHash == BadAccountHash {
		return someError
	}

	for i, cookie := range m.Cookies[accountHash] {
		if cookie.Hash == cookieHash {
			m.Cookies[accountHash] = append(
				m.Cookies[accountHash][:i],
				m.Cookies[accountHash][i+1:]...,
			)
			break
		}
	}

	return nil
}
