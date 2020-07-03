package data

import "regexp"

var AssertionTransactionPattern = regexp.MustCompile(
	`(?i)^ASSERT (?P<variableName>[a-zA-Z0-9_]+)(?P<dataPath>[a-zA-Z0-9_.]+?)? EQUALS (?P<newValue>.+)$`,
)

type AssertionTransactionData struct {
	variableName, dataPath, newValue string
}

func (d AssertionTransactionData) GetVariableName() string {
	return d.variableName
}

func (d AssertionTransactionData) GetDataPath() string {
	return d.dataPath
}

func (d AssertionTransactionData) GetNewValue() string {
	return d.newValue
}

func (d *AssertionTransactionData) SetField(name, value string) {
	switch name {
	case variableNameGroupName:
		d.variableName = value
	case dataPathGroupName:
		d.dataPath = value
	case newValueGroupName:
		d.newValue = value
	}
}

func IsAssertion(transaction string) bool {
	return AssertionTransactionPattern.MatchString(transaction)
}
