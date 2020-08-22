package main

import (
	"controller"
	"google.golang.org/grpc"
	"logger"
	"net"
	"os"
	"test_case_runner"
	"tests_runner_client/client"
	"utils"
)

var (
	address     = utils.MustGetEnv("TESTS_RUNNER_ADDRESS")
	dbFilesRoot = utils.MustGetEnv("DB_FILES_ROOT_PATH")
)

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to start listen address: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"address": address,
			},
		}, logger.Error)
		os.Exit(1)
	}

	s := grpc.NewServer()
	test_case_runner.RegisterTestRunnerServiceServer(
		s,
		controller.New(client.New(dbFilesRoot)),
	)

	logger.Info("Starting GRPC server on address: " + address)

	err = s.Serve(listener)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to start server: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"address": address,
			},
		}, logger.Error)
		os.Exit(1)
	}
}
