package assignment

import "test_runner_meta/mock/transaction/constant"

type MockCommandEmptyResult struct {
}

func (m MockCommandEmptyResult) Run(string) (map[string]interface{}, error) {
	return nil, nil
}

type MockCommandSomeResult struct {
}

func (m MockCommandSomeResult) Run(string) (map[string]interface{}, error) {
	return MockCommandResult, nil
}

type MockCommandWithError struct {
}

func (c MockCommandWithError) Run(string) (map[string]interface{}, error) {
	return nil, constant.RunCommandError
}
