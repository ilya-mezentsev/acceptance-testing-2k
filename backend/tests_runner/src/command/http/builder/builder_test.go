package builder

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"interfaces"
	mockCommand "mock/command"
	"test_utils"
	"testing"
)

var (
	db      *sqlx.DB
	builder interfaces.CommandBuilder
)

func init() {
	dbFile := test_utils.MustGetEnv("TEST_RUNNER_DB_FILE")

	var err error
	db, err = sqlx.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}

	builder = New(db)
}

func TestBuilder_BuildSuccess(t *testing.T) {
	mockCommand.InitTables(db)
	defer mockCommand.DropTables(db)

	command, err := builder.Build(mockCommand.ObjectName, mockCommand.CreateCommandName)

	test_utils.AssertNil(err, t)
	test_utils.AssertNotNil(command, t)
}

func TestBuilder_BuildNoDB(t *testing.T) {
	mockCommand.DropTables(db)

	command, err := builder.Build(mockCommand.ObjectName, mockCommand.CreateCommandName)

	test_utils.AssertNotNil(err, t)
	test_utils.AssertNil(command, t)
}

func TestBuilder_BuildNoTable(t *testing.T) {
	mockCommand.InitTables(db)
	mockCommand.DropCommandsSettings(db)
	defer mockCommand.DropTables(db)

	command, err := builder.Build(mockCommand.ObjectName, mockCommand.CreateCommandName)

	test_utils.AssertNotNil(err, t)
	test_utils.AssertNil(command, t)
}

func TestBuilder_GetCommandSettingsSuccess(t *testing.T) {
	mockCommand.InitTables(db)
	defer mockCommand.DropTables(db)

	settings, err := builder.(Builder).getCommandSettings(mockCommand.CreateCommandHash)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(mockCommand.Settings[0]["method"], settings.GetMethod(), t)
	test_utils.AssertEqual(mockCommand.Settings[0]["base_url"], settings.GetBaseURL(), t)
	test_utils.AssertEqual(mockCommand.Settings[0]["endpoint"], settings.GetEndpoint(), t)
	test_utils.AssertEqual(mockCommand.Settings[0]["pass_arguments_in_url"], settings.ShouldPassArgumentsInURL(), t)
	for _, header := range mockCommand.Headers {
		test_utils.AssertEqual(header["value"], settings.GetHeaders()[header["key"].(string)], t)
	}
	for index, expectedCookie := range mockCommand.Cookies {
		test_utils.AssertEqual(expectedCookie["key"].(string), settings.GetCookies()[index].Name, t)
		test_utils.AssertEqual(expectedCookie["value"].(string), settings.GetCookies()[index].Value, t)
	}
}

func TestBuilder_GetCommandSettingsNoHeadersTable(t *testing.T) {
	mockCommand.InitTables(db)
	mockCommand.DropCommandsHeaders(db)
	defer mockCommand.DropTables(db)

	_, err := builder.(Builder).getCommandSettings(mockCommand.CreateCommandHash)

	test_utils.AssertNotNil(err, t)
}

func TestBuilder_GetCommandSettingsNoCookiesTable(t *testing.T) {
	mockCommand.InitTables(db)
	mockCommand.DropCommandsCookies(db)
	defer mockCommand.DropTables(db)

	_, err := builder.(Builder).getCommandSettings(mockCommand.CreateCommandHash)

	test_utils.AssertNotNil(err, t)
}
