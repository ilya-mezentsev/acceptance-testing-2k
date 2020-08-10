package test_object

import "fmt"

type QueryProvider struct {
}

func (p QueryProvider) CreateQuery() string {
	return `INSERT INTO objects(name, hash) VALUES(:name, :hash)`
}

func (p QueryProvider) GetAllQuery() string {
	return `SELECT name, hash FROM objects`
}

func (p QueryProvider) GetQuery() string {
	return `SELECT name, hash FROM objects WHERE hash = ?`
}

func (p QueryProvider) UpdateQuery(fieldName string) string {
	return fmt.Sprintf(`UPDATE objects SET %s = :new_value WHERE hash = :hash`, fieldName)
}

func (p QueryProvider) DeleteQuery() string {
	return `DELETE FROM objects WHERE hash = ?`
}
