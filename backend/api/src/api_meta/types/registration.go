package types

type AccountHashAlreadyExists struct {
}

func (a AccountHashAlreadyExists) Error() string {
	return "account hash already exists"
}
