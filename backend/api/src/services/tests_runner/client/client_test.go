package client

import (
	"api_meta/mock/services"
	"io/ioutil"
	"log"
	"os"
	"test_utils"
	"testing"
	"time"
)

var (
	testHash          = "some-hash"
	testPath          = "/some/path"
	serviceServerMock = services.GRPCServiceServerMock{}
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestGrpc_CallSuccess(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	serviceAddress := test_utils.TestNameToServerAddress[t.Name()]
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), serviceServerMock)
	<-serverStarted

	report, err := New(serviceAddress).Call(testHash, testPath)

	test_utils.AssertNil(err, t)
	test_utils.AssertEqual(
		services.MockTestCasesReport.Report.PassedCount,
		report.Report.PassedCount,
		t,
	)
	test_utils.AssertEqual(
		services.MockTestCasesReport.Report.FailedCount,
		report.Report.FailedCount,
		t,
	)
}

func TestGrpc_CallServiceError(t *testing.T) {
	test_utils.AddServerAddressForTest(t.Name())
	serviceAddress := test_utils.TestNameToServerAddress[t.Name()]
	serverStarted := time.After(test_utils.BeforeServerStartsDuration)
	go test_utils.InitGRPCServer(t.Name(), serviceServerMock)
	<-serverStarted

	_, err := New(serviceAddress).Call(services.BadAccountHash, testPath)

	test_utils.AssertNotNil(err, t)
}

// for coverage only)
func TestGrpc_CallBadOpts(t *testing.T) {
	c := New("blah")
	c.opts = nil

	_, err := c.Call(testHash, testPath)

	test_utils.AssertNotNil(err, t)
}
