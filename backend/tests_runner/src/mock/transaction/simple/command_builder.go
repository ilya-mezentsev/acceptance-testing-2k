package simple

import (
	"interfaces"
	"mock/transaction/constant"
)

type MockCommandBuilder struct {
}

func (b MockCommandBuilder) Build(string, string) (interfaces.Command, error) {
	return MockCommand{}, nil
}

type MockCommandBuilderWithError struct {
}

func (b MockCommandBuilderWithError) Build(string, string) (interfaces.Command, error) {
	return nil, constant.BuildCommandError
}

type MockErroredCommandBuilder struct {
}

func (b MockErroredCommandBuilder) Build(string, string) (interfaces.Command, error) {
	return MockCommandWithError{}, nil
}
