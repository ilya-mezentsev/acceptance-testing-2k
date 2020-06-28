package interfaces

type (
	TransactionData interface {
		SetField(name, value string)
	}

	variableContainer interface {
		GetVariableName() string
	}

	SimpleTransactionData interface {
		TransactionData
		GetCommand() string
		GetObject() string
		GetArguments() string
	}

	AssignmentTransactionData interface {
		SimpleTransactionData
		variableContainer
	}

	AssertionTransactionData interface {
		TransactionData
		variableContainer
		GetDataPath() string
		GetNewValue() string
	}
)
