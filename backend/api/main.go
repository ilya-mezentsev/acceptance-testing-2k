package main

import (
	"api_meta/interfaces"
	crudController "controllers/crud"
	sessionController "controllers/session"
	"controllers/tests_runner"
	"db_connector"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"logger"
	"middlewares/csrf"
	"net/http"
	"repositories/crud"
	"repositories/crud/query_providers"
	registrationRepo "repositories/registration"
	sessionRepo "repositories/session"
	"repositories/test_command_meta"
	"services/pool"
	"services/registration"
	"services/session"
	"services/test_command"
	"services/test_command/cookies_deleter"
	"services/test_command/headers_deleter"
	"services/test_command/mass_update/base_url"
	"services/test_command/mass_update/cookies"
	"services/test_command/mass_update/headers"
	"services/test_command/mass_update/timeout"
	"services/test_command/meta"
	"services/test_object"
	"services/tests_runner/client"
	"services/tests_runner/file_creator"
	"services/tests_runner/runner"
	"time"
	"utils"
)

var (
	r                *mux.Router
	db               *sqlx.DB
	connector        *db_connector.Connector
	crudServicesPool pool.CRUDServicesPool

	registrationService              interfaces.CreateService
	sessionService                   session.Service
	testCommandService               interfaces.CRUDService
	massBaseUrlsUpdateService        base_url.Service
	massTimeoutsUpdaterService       timeout.Service
	massCookiesCreatorService        cookies.Service
	massHeadersCreatorService        headers.Service
	testCommandMetaCreatorService    meta.Service
	testCommandHeadersDeleterService headers_deleter.Service
	testCommandCookiesDeleterService cookies_deleter.Service
	testObjectService                interfaces.CRUDService
	testsFileCreatorService          file_creator.Service
	testsRunnerService               runner.Service
	testsRunnerClient                client.Grpc

	registrationRepository    interfaces.RegistrationRepository
	sessionRepository         interfaces.SessionRepository
	testCommandRepository     interfaces.CRUDRepository
	testCommandMetaRepository test_command_meta.Repository
	testObjectRepository      interfaces.CRUDRepository

	csrfMiddleware csrf.Middleware

	projectDBFilePath  string
	filesRootPath      string
	testsRunnerAddress string
	csrfPrivateKey     string
	apiAddress         string
)

func init() {
	r = mux.NewRouter()
	crudServicesPool = pool.New()

	readEnv()
	connector = db_connector.New(filesRootPath)
	var err error
	db, err = sqlx.Open("sqlite3", projectDBFilePath)
	if err != nil {
		panic(err)
	}
	initProjectDB()

	// should be called after reading env
	initRepositories()

	// should be called after initializing repositories
	initServices()

	crudServicesPool.AddCRUDService("test-object", testObjectService)
	crudServicesPool.AddCRUDService("test-command", testCommandService)
	crudServicesPool.AddService(
		"test-command-meta",
		[]string{pool.CreateServiceOperationType, pool.UpdateServiceOperationType},
		testCommandMetaCreatorService,
	)
	crudServicesPool.AddService(
		"test-command-headers",
		[]string{pool.DeleteServiceOperationType},
		testCommandHeadersDeleterService,
	)
	crudServicesPool.AddService(
		"test-command-cookies",
		[]string{pool.DeleteServiceOperationType},
		testCommandCookiesDeleterService,
	)
	crudServicesPool.AddService(
		"mass-base-urls",
		[]string{pool.UpdateServiceOperationType},
		massBaseUrlsUpdateService,
	)
	crudServicesPool.AddService(
		"mass-timeouts",
		[]string{pool.UpdateServiceOperationType},
		massTimeoutsUpdaterService,
	)
	crudServicesPool.AddService(
		"mass-cookies",
		[]string{pool.CreateServiceOperationType},
		massCookiesCreatorService,
	)
	crudServicesPool.AddService(
		"mass-headers",
		[]string{pool.CreateServiceOperationType},
		massHeadersCreatorService,
	)
	crudServicesPool.AddService(
		"registration",
		[]string{pool.CreateServiceOperationType},
		registrationService,
	)

	initMiddleware()

	// should be called after adding crud services to pool and initializing middleware
	initControllers()
}

func initProjectDB() {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS accounts(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		hash VARCHAR(32) NOT NULL UNIQUE,
		verified BOOLEAN NOT NULL DEFAULT 0 CHECK (verified IN (0,1)),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		panic(err)
	}
}

func initRepositories() {
	registrationRepository = registrationRepo.New(db, connector)
	sessionRepository = sessionRepo.New(connector)

	testCommandRepository = crud.New(connector, query_providers.TestCommandQueryProvider{})
	testCommandMetaRepository = test_command_meta.New(connector)
	testObjectRepository = crud.New(connector, query_providers.TestObjectQueryProvider{})
}

func readEnv() {
	testsRunnerAddress = utils.MustGetEnv("TESTS_RUNNER_ADDRESS")
	projectDBFilePath = utils.MustGetEnv("PROJECT_DB_FILE_PATH")
	filesRootPath = utils.MustGetEnv("FILES_ROOT_PATH")
	csrfPrivateKey = utils.MustGetEnv("CSRF_PRIVATE_KEY")
	apiAddress = utils.MustGetEnv("API_ADDRESS")
}

func initServices() {
	registrationService = registration.New(registrationRepository, filesRootPath)
	sessionService = session.New(sessionRepository)

	testCommandService = test_command.New(testCommandRepository, testCommandMetaRepository)
	massBaseUrlsUpdateService = base_url.New(testCommandRepository)
	massTimeoutsUpdaterService = timeout.New(testCommandRepository)
	massCookiesCreatorService = cookies.New(testCommandMetaRepository)
	massHeadersCreatorService = headers.New(testCommandMetaRepository)
	testCommandMetaCreatorService = meta.New(testCommandMetaRepository)
	testCommandHeadersDeleterService = headers_deleter.New(testCommandMetaRepository)
	testCommandCookiesDeleterService = cookies_deleter.New(testCommandMetaRepository)
	testObjectService = test_object.New(testObjectRepository)

	testsRunnerClient = client.New(testsRunnerAddress)
	testsRunnerService = runner.New(testsRunnerClient)
	testsFileCreatorService = file_creator.New()
}

func initMiddleware() {
	csrfMiddleware = csrf.Middleware{PrivateKey: csrfPrivateKey}
}

func initControllers() {
	r.Use(csrfMiddleware.CheckCSRFToken)

	sessionController.Init(r, sessionService)
	tests_runner.Init(r, testsFileCreatorService, testsRunnerService)
	crudController.Init(r, crudServicesPool)
}

func main() {
	logger.Info("Starting application on address " + apiAddress)

	log.Fatal((&http.Server{
		Handler:      r,
		Addr:         apiAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe())
}
