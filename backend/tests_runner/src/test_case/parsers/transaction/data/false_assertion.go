package data

import "regexp"

var FalseAssertionTransactionPattern = regexp.MustCompile(
	`(?i)^ASSERT_FALSE (?P<variableName>[a-zA-Z0-9_]+)(?:\.(?P<dataPath>[a-zA-Z0-9_.]+?))?$`,
)

type FalseAssertionTransactionData struct {
	ConstAssertionTransactionData
}

func (d FalseAssertionTransactionData) GetNewValue() string {
	return "false"
}

func IsFalseAssertion(transaction string) bool {
	return FalseAssertionTransactionPattern.MatchString(transaction)
}
