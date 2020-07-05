package simple

type MockCommand struct {
}

func (c MockCommand) Run(string) (map[string]string, error) {
	return nil, nil
}

type MockCommandWithError struct {
}

func (c MockCommandWithError) Run(string) (map[string]string, error) {
	return nil, RunCommandError
}
