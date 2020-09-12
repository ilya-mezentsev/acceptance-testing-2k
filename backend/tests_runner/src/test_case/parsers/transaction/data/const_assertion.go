package data

type ConstAssertionTransactionData struct {
	transactionTextContainer
	variableName, dataPath string
}

func (d ConstAssertionTransactionData) GetVariableName() string {
	return d.variableName
}

func (d ConstAssertionTransactionData) GetDataPath() string {
	return d.dataPath
}

func (d ConstAssertionTransactionData) GetNewValue() string {
	panic("not implemented")
}

func (d *ConstAssertionTransactionData) SetField(name, value string) {
	switch name {
	case variableNameGroupName:
		d.variableName = value
	case dataPathGroupName:
		d.dataPath = value
	}
}
