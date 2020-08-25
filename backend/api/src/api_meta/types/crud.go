package types

type UniqueEntityAlreadyExists struct {
}

func (u UniqueEntityAlreadyExists) Error() string {
	return "unique entity already exists"
}
