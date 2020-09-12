package data

import "regexp"

var TrueAssertionTransactionPattern = regexp.MustCompile(
	`(?i)^ASSERT_TRUE (?P<variableName>[a-zA-Z0-9_]+)(?:\.(?P<dataPath>[a-zA-Z0-9_.]+?))?$`,
)

type TrueAssertionTransactionData struct {
	ConstAssertionTransactionData
}

func (d TrueAssertionTransactionData) GetNewValue() string {
	return "true"
}

func IsTrueAssertion(transaction string) bool {
	return TrueAssertionTransactionPattern.MatchString(transaction)
}
