package arguments_builder

type emptyArguments struct {
}

func (a emptyArguments) Value() string {
	return ""
}

func (a emptyArguments) AmpersandSeparated() (string, error) {
	return "", nil
}
