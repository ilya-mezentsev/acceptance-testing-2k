package test_case

import (
	"interfaces"
	"regexp"
	"strings"
	"test_case/parsers/errors"
)

var (
	testCasePattern = regexp.MustCompile(`(?msi)BEGIN(.+?)END`)
)

type Parser struct {
	currentTestCase int
	testCases       [][]string
}

func (p Parser) Parse(testCases string) ([]interfaces.TestCaseTransactionsIterator, error) {
	matches := testCasePattern.FindAllStringSubmatch(testCases, -1)
	if len(matches) == 0 {
		return nil, errors.NoTestCases
	}

	var testCasesIterators []interfaces.TestCaseTransactionsIterator
	for _, match := range matches {
		testCasesIterators = append(
			testCasesIterators,
			&TestCaseTransactionsIterator{transactions: p.getMeaningRows(match[1])},
		)
	}

	return testCasesIterators, nil
}

func (p Parser) getMeaningRows(testCase string) []string {
	var meaningRows []string
	for _, row := range strings.Split(testCase, "\n") {
		row = strings.TrimSpace(row)

		if p.shouldSkipRow(row) {
			continue
		}

		meaningRows = append(meaningRows, row)
	}

	return meaningRows
}

func (p Parser) shouldSkipRow(row string) bool {
	return row == "" || len(row) == 1 || row[:2] == "//"
}
