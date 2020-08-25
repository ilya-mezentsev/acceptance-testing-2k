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
		Headers: entity["command_headers"].(string),
		Cookies: entity["command_cookies"].(string),
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
				case "headers":
					m.Commands[accountHash][commandIndex].Headers = entity.NewValue.(string)
				case "cookies":
					m.Commands[accountHash][commandIndex].Cookies = entity.NewValue.(string)
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
