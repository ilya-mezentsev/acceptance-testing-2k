package test_runner_client

import (
	"io/ioutil"
	"strings"
)

// Constants were extracted from mock_server.go
const data = `
BEGIN
  userResponse = GET USER hash-1

	ASSERT userResponse.data.name EQUALS John
END
BEGIN
	CREATE USER {"hash": "hash-3", "name": "Joe"}

	userResponse = GET USER hash-3

	ASSERT userResponse.data.name EQUALS Joe
END
BEGIN
	userResponse = GET USER hash-1

	ASSERT userResponse.data.name EQUALS Dude
END
`

const (
	PassedCount = 2
	FailedCount = 1
)

func FillTestCasesFile(filePath string) {
	err := ioutil.WriteFile(filePath, []byte(strings.TrimSpace(data)), 0644)
	if err != nil {
		panic(err)
	}
}

func FillBadTestCasesData(filePath string) {
	err := ioutil.WriteFile(filePath, []byte(``), 0644)
	if err != nil {
		panic(err)
	}
}
