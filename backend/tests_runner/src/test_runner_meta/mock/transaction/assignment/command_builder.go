package assignment

import (
	"test_runner_meta/interfaces"
	"test_runner_meta/mock/transaction/constant"
)

type MockNilResultCommandBuilder struct {
}

func (b MockNilResultCommandBuilder) Build(string, string) (interfaces.Command, error) {
	return MockCommandEmptyResult{}, nil
}

type MockNotNilResultCommandBuilder struct {
}

func (b MockNotNilResultCommandBuilder) Build(string, string) (interfaces.Command, error) {
	return MockCommandSomeResult{}, nil
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
