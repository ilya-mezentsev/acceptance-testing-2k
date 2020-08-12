package query_providers

import "fmt"

type AccountCredentialsQueryProvider struct {
}

func (p AccountCredentialsQueryProvider) CreateQuery() string {
	return `
	INSERT INTO account_credentials(login, password, hash)
	VALUES(:login, :password, :hash)`
}

func (p AccountCredentialsQueryProvider) GetAllQuery() string {
	panic("should not used here")
}

func (p AccountCredentialsQueryProvider) GetQuery() string {
	return `SELECT hash, login, verified FROM account_credentials WHERE hash = ?`
}

func (p AccountCredentialsQueryProvider) UpdateQuery(updateFieldName string) string {
	return fmt.Sprintf(`
	UPDATE account_credentials
	SET %s = :new_value
	WHERE hash = :hash AND verified = 1`, updateFieldName)
}

func (p AccountCredentialsQueryProvider) DeleteQuery() string {
	return `DELETE FROM account_credentials WHERE hash = ?`
}
