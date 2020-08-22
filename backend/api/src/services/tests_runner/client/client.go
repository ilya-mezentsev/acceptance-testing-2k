package client

import (
	"context"
	"google.golang.org/grpc"
	"logger"
	"test_case_runner"
)

type Grpc struct {
	address string
	opts    []grpc.DialOption
}

func New(address string) Grpc {
	return Grpc{
		address: address,
		opts: []grpc.DialOption{
			grpc.WithInsecure(),
		},
	}
}

func (c Grpc) Call(
	accountHash,
	testCasesPath string,
) (*test_case_runner.TestsReport, error) {
	clientConn, err := grpc.Dial(c.address, c.opts...)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to dial GRPC service: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"address": c.address,
			},
		}, logger.Error)

		return nil, err
	}

	serviceClient := test_case_runner.NewTestRunnerServiceClient(clientConn)
	report, err := serviceClient.Run(context.Background(), &test_case_runner.TestCasesRequest{
		AccountHash:   accountHash,
		TestCasesPath: testCasesPath,
	})
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "GRPC service returned error: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"address":         c.address,
				"account_hash":    accountHash,
				"test_cases_path": testCasesPath,
			},
		}, logger.Error)

		return nil, err
	}

	return report, nil
}
