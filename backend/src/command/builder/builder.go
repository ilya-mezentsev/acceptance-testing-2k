package builder

import (
	"interfaces"
)

type Builder struct {
}

func (b Builder) Build(objectName, commandName string) (interfaces.Command, error) {
	panic("implement me")
}
