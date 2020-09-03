package query_providers

import "fmt"

type TestObjectQueryProvider struct {
}

func (p TestObjectQueryProvider) CreateQuery() string {
	return `INSERT INTO objects(name, hash) VALUES(:name, :hash)`
}

func (p TestObjectQueryProvider) GetAllQuery() string {
	return `SELECT name, hash FROM objects`
}

func (p TestObjectQueryProvider) GetQuery() string {
	return `SELECT name, hash FROM objects WHERE hash = ?`
}

func (p TestObjectQueryProvider) UpdateQuery(fieldName string) string {
	return fmt.Sprintf(`UPDATE objects SET %s = :new_value WHERE hash = :hash`, fieldName)
}

func (p TestObjectQueryProvider) DeleteQuery() string {
	return `
	PRAGMA foreign_keys=ON;
	DELETE FROM objects WHERE hash = ?;
	PRAGMA foreign_keys=OFF;
	`
}
