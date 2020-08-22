package controller

import (
	"context"
	"google.golang.org/grpc"
	"test_case_runner"
	mockController "test_runner_meta/mock/controller"
	"test_utils"
	"testing"
	"time"
)

func makeRequest(testName, accountHash, testCasesPath string) *test_case_runner.TestsReport {
	opts := grpc.WithInsecure()
	clientConn, err := grpc.Dial(test_utils.TestNameToServerAddress[testName], opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = clientConn.Close()
	}()

	client := test_case_runner.NewTestRunnerServiceClient(clientConn)
	request := &test_case_runner.TestCasesRequest{
		AccountHash:   accountHash,
		TestCasesPath: testCasesPath,
	}

	report, err := client.Run(context.Background(), request)
	if err != nil {
		panic(err)
	}

	return report
}

func TestController_RunSimple(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	simpleClient := &mockController.SimpleTestRunnerClientMock{}
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), New(simpleClient))
	<-serverStarted

	makeRequest(t.Name(), "hash", "filename")

	test_utils.AssertTrue(simpleClient.CalledWith("hash", "filename"), t)
}

func TestController_RunCheckResponse(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	respondClient := &mockController.WithReportTestRunnerClientMock{}
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), New(respondClient))
	<-serverStarted

	response := makeRequest(t.Name(), "hash", "filename")

	test_utils.AssertTrue(respondClient.CalledWith("hash", "filename"), t)
	test_utils.AssertEqual(
		int64(mockController.TestsReport.PassedCount),
		response.Report.PassedCount,
		t,
	)
	test_utils.AssertEqual(
		int64(mockController.TestsReport.FailedCount),
		response.Report.FailedCount,
		t,
	)
	test_utils.AssertEqual(
		len(mockController.TestsReport.Errors),
		len(response.Report.Errors),
		t,
	)
}

func TestController_RunCheckApplicationError(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	applicationErrorClient := &mockController.WithApplicationErrorTestRunnerClientMock{}
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), New(applicationErrorClient))
	<-serverStarted

	response := makeRequest(t.Name(), "hash", "filename")

	test_utils.AssertTrue(
		applicationErrorClient.CalledWith("hash", "filename"),
		t,
	)
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
