package data

import "regexp"

var AssignmentTransactionPattern = regexp.MustCompile(
	`(?P<variableName>[a-zA-Z0-9_]+?) ?= ?(?P<command>[a-zA-Z_]+?) (?P<object>[a-zA-Z_]+?)(?: (?P<arguments>.+))?$`,
)

type AssignmentTransactionData struct {
	transactionTextContainer
	SimpleTransactionData
	variableName string
}

func (d AssignmentTransactionData) GetVariableName() string {
	return d.variableName
}

func (d *AssignmentTransactionData) SetField(name, value string) {
	switch name {
	case variableNameGroupName:
		d.variableName = value
	default:
		d.SimpleTransactionData.SetField(name, value)
	}
}

func IsAssignment(transaction string) bool {
	return AssignmentTransactionPattern.MatchString(transaction)
}
