package main

import (
	"api_meta/interfaces"
	crudController "controllers/crud"
	registrationController "controllers/registration"
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
	"services/pool"
	"services/registration"
	"services/session"
	"services/test_command"
	"services/test_object"
	"services/tests_runner/client"
	"services/tests_runner/runner"
	"services/tests_runner/tests_file"
	"time"
	"utils"
)

var (
	r                *mux.Router
	db               *sqlx.DB
	connector        db_connector.Connector
	crudServicesPool pool.CRUDServicesPool

	registrationService registration.Service
	sessionService      session.Service
	testCommandService  interfaces.CRUDService
	testObjectService   interfaces.CRUDService
	testsRunnerService  runner.Service
	testsRunnerClient   client.Grpc
	testsFileManager    tests_file.Manager

	registrationRepository interfaces.RegistrationRepository
	sessionRepository      interfaces.SessionRepository
	testCommandRepository  interfaces.CRUDRepository
	testObjectRepository   interfaces.CRUDRepository

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

	// should be called after reading env
	initRepositories()

	// should be called after initializing repositories
	initServices()

	for entityType, crudService := range map[string]interfaces.CRUDService{
		"test-object":  testObjectService,
		"test-command": testCommandService,
	} {
		crudServicesPool.AddService(entityType, crudService)
	}

	initMiddleware()

	// should be called after adding crud services to pool and initializing middleware
	initControllers()
}

func initRepositories() {
	registrationRepository = registrationRepo.New(db, connector)
	sessionRepository = sessionRepo.New(connector)

	testCommandRepository = crud.New(connector, query_providers.TestCommandQueryProvider{})
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

	testCommandService = test_command.New(testCommandRepository)
	testObjectService = test_object.New(testObjectRepository)

	testsRunnerClient = client.New(testsRunnerAddress)
	testsFileManager = tests_file.New(filesRootPath)
	testsRunnerService = runner.New(testsFileManager, testsRunnerClient)
}

func initMiddleware() {
	csrfMiddleware = csrf.Middleware{PrivateKey: csrfPrivateKey}
}

func initControllers() {
	r.Use(csrfMiddleware.CheckCSRFToken)

	crudController.Init(r, crudServicesPool)
	sessionController.Init(r, sessionService)
	tests_runner.Init(r, testsRunnerService)
	registrationController.Init(r, registrationService)
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
