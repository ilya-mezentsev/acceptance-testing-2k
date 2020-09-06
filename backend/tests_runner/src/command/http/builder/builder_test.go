package builder

import (
	"command/http/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"test_runner_meta/interfaces"
	"test_utils"
	"testing"
	"utils"
)

var (
	db      *sqlx.DB
	builder interfaces.CommandBuilder
)

func init() {
	dbFile := utils.MustGetEnv("TEST_DB_FILE")

	var err error
	db, err = sqlx.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}

	builder = New(db)
}

func TestBuilder_BuildSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	command, err := builder.Build(test_utils.ObjectName, test_utils.CreateCommandName)

	test_utils.AssertNil(err, t)
	test_utils.AssertNotNil(command, t)
}

func TestBuilder_BuildNoCommand(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	_, errNoObject := builder.Build("blah-blah", test_utils.CreateCommandName)
	_, errNoCommand := builder.Build(test_utils.ObjectName, "blah-blah")

	test_utils.AssertErrorsEqual(errors.CommandNotFound, errNoObject, t)
	test_utils.AssertErrorsEqual(errors.CommandNotFound, errNoCommand, t)
}

func TestBuilder_BuildNoDB(t *testing.T) {
	test_utils.DropTables(db)

	command, err := builder.Build(test_utils.ObjectName, test_utils.CreateCommandName)

	test_utils.AssertNotNil(err, t)
	test_utils.AssertNil(command, t)
}

func TestBuilder_BuildNoTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsSettings(db)
	defer test_utils.DropTables(db)

	command, err := builder.Build(test_utils.ObjectName, test_utils.CreateCommandName)

	test_utils.AssertNotNil(err, t)
	test_utils.AssertNil(command, t)
}

func TestBuilder_GetCommandSettingsSuccess(t *testing.T) {
	test_utils.InitTables(db)
	defer test_utils.DropTables(db)

	settings, err := builder.(Builder).getCommandSettings(test_utils.CreateCommandHash)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(test_utils.Settings[0]["method"], settings.GetMethod(), t)
	test_utils.AssertEqual(test_utils.Settings[0]["base_url"], settings.GetBaseURL(), t)
	test_utils.AssertEqual(test_utils.Settings[0]["endpoint"], settings.GetEndpoint(), t)
	test_utils.AssertEqual(
		test_utils.Settings[0]["pass_arguments_in_url"],
		settings.ShouldPassArgumentsInURL(),
		t,
	)
	for _, header := range test_utils.Headers {
		test_utils.AssertEqual(header["value"], settings.GetHeaders()[header["key"].(string)], t)
	}
	for index, expectedCookie := range test_utils.Cookies {
		test_utils.AssertEqual(expectedCookie["key"].(string), settings.GetCookies()[index].Name, t)
		test_utils.AssertEqual(expectedCookie["value"].(string), settings.GetCookies()[index].Value, t)
	}
}

func TestBuilder_GetCommandSettingsNoHeadersTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsHeaders(db)
	defer test_utils.DropTables(db)

	_, err := builder.(Builder).getCommandSettings(test_utils.CreateCommandHash)

	test_utils.AssertNotNil(err, t)
}

func TestBuilder_GetCommandSettingsNoCookiesTable(t *testing.T) {
	test_utils.InitTables(db)
	test_utils.DropCommandsCookies(db)
	defer test_utils.DropTables(db)

	_, err := builder.(Builder).getCommandSettings(test_utils.CreateCommandHash)

	test_utils.AssertNotNil(err, t)
}
