package controller

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"interfaces"
	"math/rand"
	mockController "mock/controller"
	"net"
	"test_case_runner"
	"test_utils"
	"testing"
	"time"
)

var (
	testNameToServerAddress    = map[string]string{}
	beforeServerStartsDuration = 5 * time.Millisecond
)

func addServerAddressForTest(testName string) {
	var port int
	for port < 1000 {
		port = rand.Intn(8000)
	}

	testNameToServerAddress[testName] = fmt.Sprintf("0.0.0.0:%d", port)
}

func initGRPCServer(testName string, client interfaces.TestsRunnerClient) {
	lis, err := net.Listen("tcp", testNameToServerAddress[testName])
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	test_case_runner.RegisterTestRunnerServiceServer(s, New(client))

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func makeRequest(testName, accountHash, testCasesFilename string) *test_case_runner.TestsReport {
	opts := grpc.WithInsecure()
	clientConn, err := grpc.Dial(testNameToServerAddress[testName], opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = clientConn.Close()
	}()

	client := test_case_runner.NewTestRunnerServiceClient(clientConn)
	request := &test_case_runner.TestCasesRequest{
		AccountHash:       accountHash,
		TestCasesFilename: testCasesFilename,
	}

	report, err := client.Run(context.Background(), request)
	if err != nil {
		panic(err)
	}

	return report
}

func TestController_RunSimple(t *testing.T) {
	addServerAddressForTest(t.Name())
	simpleClient := &mockController.SimpleTestRunnerClientMock{}
	serverStarted := time.After(beforeServerStartsDuration)
	go initGRPCServer(t.Name(), simpleClient)
	<-serverStarted

	makeRequest(t.Name(), "hash", "filename")

	test_utils.AssertTrue(simpleClient.CalledWith("hash", "filename"), t)
}

func TestController_RunCheckResponse(t *testing.T) {
	addServerAddressForTest(t.Name())
	respondClient := &mockController.WithReportTestRunnerClientMock{}
	serverStarted := time.After(beforeServerStartsDuration)
	go initGRPCServer(t.Name(), respondClient)
	<-serverStarted

	response := makeRequest(t.Name(), "hash", "filename")

	test_utils.AssertTrue(respondClient.CalledWith("hash", "filename"), t)
	test_utils.AssertEqual(int64(mockController.TestsReport.PassedCount), response.Report.PassedCount, t)
	test_utils.AssertEqual(int64(mockController.TestsReport.FailedCount), response.Report.FailedCount, t)
	test_utils.AssertEqual(len(mockController.TestsReport.Errors), len(response.Report.Errors), t)
}

func TestController_RunCheckApplicationError(t *testing.T) {
	addServerAddressForTest(t.Name())
	applicationErrorClient := &mockController.WithApplicationErrorTestRunnerClientMock{}
	serverStarted := time.After(beforeServerStartsDuration)
	go initGRPCServer(t.Name(), applicationErrorClient)
	<-serverStarted

	response := makeRequest(t.Name(), "hash", "filename")

	test_utils.AssertTrue(applicationErrorClient.CalledWith("hash", "filename"), t)
	test_utils.AssertEqual(
		response.ApplicationError.Code,
		mockController.ApplicationError.Code,
		t,
	)
	test_utils.AssertEqual(
		response.ApplicationError.Description,
		mockController.ApplicationError.Description,
		t,
	)
}
