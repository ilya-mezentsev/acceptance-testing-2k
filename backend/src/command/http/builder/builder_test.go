package builder

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"interfaces"
	mockCommand "mock/command"
	"os"
	"testing"
	"utils"
)

var (
	db      *sqlx.DB
	builder interfaces.CommandBuilder
)

func init() {
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		panic("DB_FILE is not provided")
	}

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

	command, err := builder.Build(mockCommand.CreatedObjectName, mockCommand.CreatedCommandName)

	utils.AssertNil(err, t)
	utils.AssertNotNil(command, t)
}

func TestBuilder_BuildNoDB(t *testing.T) {
	mockCommand.DropTables(db)

	command, err := builder.Build(mockCommand.CreatedObjectName, mockCommand.CreatedCommandName)

	utils.AssertNotNil(err, t)
	utils.AssertNil(command, t)
}

func TestBuilder_BuildNoTable(t *testing.T) {
	mockCommand.InitTables(db)
	mockCommand.DropCommandsSettings(db)
	defer mockCommand.DropTables(db)

	command, err := builder.Build(mockCommand.CreatedObjectName, mockCommand.CreatedCommandName)

	utils.AssertNotNil(err, t)
	utils.AssertNil(command, t)
}

func TestBuilder_GetCommandSettingsSuccess(t *testing.T) {
	mockCommand.InitTables(db)
	defer mockCommand.DropTables(db)

	settings, err := builder.(Builder).getCommandSettings(mockCommand.CreatedCommandHash)

	utils.AssertNil(err, t)
	utils.AssertEqual(mockCommand.Settings[0]["method"], settings.GetMethod(), t)
	utils.AssertEqual(mockCommand.Settings[0]["base_url"], settings.GetBaseURL(), t)
	utils.AssertEqual(mockCommand.Settings[0]["endpoint"], settings.GetEndpoint(), t)
	utils.AssertEqual(mockCommand.Settings[0]["pass_arguments_in_url"], settings.ShouldPassArgumentsInURL(), t)
	for _, header := range mockCommand.Headers {
		utils.AssertEqual(header["value"], settings.GetHeaders()[header["key"].(string)], t)
	}
	for index, expectedCookie := range mockCommand.Cookies {
		utils.AssertEqual(expectedCookie["key"].(string), settings.GetCookies()[index].Name, t)
		utils.AssertEqual(expectedCookie["value"].(string), settings.GetCookies()[index].Value, t)
	}
}

func TestBuilder_GetCommandSettingsNoHeadersTable(t *testing.T) {
	mockCommand.InitTables(db)
	mockCommand.DropCommandsHeaders(db)
	defer mockCommand.DropTables(db)

	_, err := builder.(Builder).getCommandSettings(mockCommand.CreatedCommandHash)

	utils.AssertNotNil(err, t)
}

func TestBuilder_GetCommandSettingsNoCookiesTable(t *testing.T) {
	mockCommand.InitTables(db)
	mockCommand.DropCommandsCookies(db)
	defer mockCommand.DropTables(db)

	_, err := builder.(Builder).getCommandSettings(mockCommand.CreatedCommandHash)

	utils.AssertNotNil(err, t)
}
