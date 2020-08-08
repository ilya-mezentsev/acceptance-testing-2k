package simple

import "test_runner_meta/mock/transaction/constant"

type MockCommand struct {
}

func (c MockCommand) Run(string) (map[string]interface{}, error) {
	return nil, nil
}

type MockCommandWithError struct {
}

func (c MockCommandWithError) Run(string) (map[string]interface{}, error) {
	return nil, constant.RunCommandError
}
