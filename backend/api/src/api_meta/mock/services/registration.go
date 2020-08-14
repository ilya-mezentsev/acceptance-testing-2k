package services

import (
	"api_meta/models"
	"api_meta/types"
)

type RegistrationRepositoryMock struct {
	AccountHashes      map[string]bool
	AccountCredentials map[string]models.AccountCredentialsRecord
}

func (m *RegistrationRepositoryMock) Reset() {
	m.AccountHashes = map[string]bool{}
	m.AccountCredentials = map[string]models.AccountCredentialsRecord{}
}

func (m *RegistrationRepositoryMock) CreateAccount(accountHash string) error {
	if accountHash == BadAccountHash {
		return someError
	} else if accountHash == ExistsAccountHash {
		return types.AccountHashAlreadyExists{}
	}

	m.AccountHashes[accountHash] = true
	return nil
}

func (m *RegistrationRepositoryMock) CreateAccountCredentials(record models.AccountCredentialsRecord) error {
	if record.Login == BadLogin {
		return someError
	}

	m.AccountCredentials[record.AccountHash] = record
	return nil
}
