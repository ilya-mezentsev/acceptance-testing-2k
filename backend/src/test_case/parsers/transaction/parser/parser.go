package parser

import (
	"interfaces"
	"regexp"
	"test_case/parsers/errors"
)

func Parse(
	reg *regexp.Regexp,
	transaction string,
	transactionData interfaces.TransactionData,
) error {
	match := reg.FindStringSubmatch(transaction)
	if len(match) == 0 {
		return errors.InvalidTransactionFormat
	}

	for i, name := range reg.SubexpNames() {
		if i == 0 || name == "" {
			continue
		}

		transactionData.SetField(name, match[i])
	}

	return nil
}
