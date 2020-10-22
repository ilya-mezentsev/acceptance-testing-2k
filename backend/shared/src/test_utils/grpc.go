package test_utils

import (
	"fmt"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"test_case_runner"
	"time"
)

var (
	TestNameToServerAddress    = map[string]string{}
	BeforeServerStartsDuration = 5 * time.Millisecond
)

func AddServerAddressForTest(testName string) {
	var port int
	for port < 1000 {
		port = rand.Intn(8000)
	}

	TestNameToServerAddress[testName] = fmt.Sprintf("0.0.0.0:%d", port)
}

func InitGRPCServer(testName string, service test_case_runner.TestRunnerServiceServer) {
	lis, err := net.Listen("tcp", TestNameToServerAddress[testName])
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	test_case_runner.RegisterTestRunnerServiceServer(s, service)

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
