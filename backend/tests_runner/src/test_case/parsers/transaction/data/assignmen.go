package data

import "regexp"

var AssignmentTransactionPattern = regexp.MustCompile(
	`(?P<variableName>[a-zA-Z0-9_]+?) ?= ?(?P<command>[a-zA-Z_]+?) (?P<object>[a-zA-Z_]+?)(?: (?P<arguments>.+))?$`,
)

type AssignmentTransactionData struct {
	transactionTextContainer
	command, object, variableName, arguments string
}

func (d AssignmentTransactionData) GetVariableName() string {
	return d.variableName
}

func (d AssignmentTransactionData) GetCommand() string {
	return d.command
}

func (d AssignmentTransactionData) GetObject() string {
	return d.object
}

func (d AssignmentTransactionData) GetArguments() string {
	return d.arguments
}

func (d *AssignmentTransactionData) SetField(name, value string) {
	switch name {
	case variableNameGroupName:
		d.variableName = value
	case commandGroupName:
		d.command = value
	case objectGroupName:
		d.object = value
	case argumentsGroupName:
		d.arguments = value
	}
}

func IsAssignment(transaction string) bool {
	return AssignmentTransactionPattern.MatchString(transaction)
}
