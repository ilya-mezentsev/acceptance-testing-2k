package services

import "errors"

type DefaultAccountRepositoryMock struct {
	Hashes []string
}

func (d DefaultAccountRepositoryMock) GetNonVerifiedAccountsCreatedDayAgo() ([]string, error) {
	return d.Hashes, nil
}

func (d DefaultAccountRepositoryMock) DeleteAccounts(hashes []string) error {
	return nil
}

type BadGetHashesAccountRepositoryMock struct {
}

func (b BadGetHashesAccountRepositoryMock) GetNonVerifiedAccountsCreatedDayAgo() ([]string, error) {
	return nil, errors.New("some-error")
}

func (b BadGetHashesAccountRepositoryMock) DeleteAccounts(hashes []string) error {
	panic("implement me")
}

type BadDeleteAccountRepositoryMock struct {
}

func (b BadDeleteAccountRepositoryMock) GetNonVerifiedAccountsCreatedDayAgo() ([]string, error) {
	return []string{"1"}, nil
}

func (b BadDeleteAccountRepositoryMock) DeleteAccounts(hashes []string) error {
	return errors.New("some-error")
}
