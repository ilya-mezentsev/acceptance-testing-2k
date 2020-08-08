package interfaces

type (
	CommandBuilder interface {
		Build(objectName, commandName string) (Command, error)
	}

	Command interface {
		Run(arguments string) (map[string]interface{}, error)
	}
)
