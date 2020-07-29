package interfaces

type (
	transactionTextGetter interface {
		GetTransactionText() string
	}

	variableContainer interface {
		GetVariableName() string
	}

	TransactionDataSetter interface {
		SetTransactionText(text string)
		SetField(name, value string)
	}

	SimpleTransactionData interface {
		transactionTextGetter
		GetCommand() string
		GetObject() string
		GetArguments() string
	}

	AssignmentTransactionData interface {
		SimpleTransactionData
		variableContainer
	}

	AssertionTransactionData interface {
		transactionTextGetter
		variableContainer
		GetDataPath() string
		GetNewValue() string
	}
)
