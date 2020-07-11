package data

import "regexp"

var SimpleTransactionPattern = regexp.MustCompile(
	`(?P<command>[a-zA-Z_]+?) (?P<object>[a-zA-Z_]+?) ?(?P<arguments>{.+})?$`,
)

type SimpleTransactionData struct {
	transactionTextContainer
	command, object, arguments string
}

func (d SimpleTransactionData) GetCommand() string {
	return d.command
}

func (d SimpleTransactionData) GetObject() string {
	return d.object
}

func (d SimpleTransactionData) GetArguments() string {
	return d.arguments
}

func (d *SimpleTransactionData) SetField(name, value string) {
	switch name {
	case commandGroupName:
		d.command = value
	case objectGroupName:
		d.object = value
	case argumentsGroupName:
		d.arguments = value
	}
}

func IsSimple(transaction string) bool {
	return SimpleTransactionPattern.MatchString(transaction)
}
