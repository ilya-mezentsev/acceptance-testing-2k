package services

type SessionRepositoryMock struct {
	Accounts map[string]bool
}

func (m SessionRepositoryMock) AccountExists(accountHash string) (bool, error) {
	if accountHash == BadAccountHash {
		return false, someError
	}

	_, ok := m.Accounts[accountHash]
	return ok, nil
}
