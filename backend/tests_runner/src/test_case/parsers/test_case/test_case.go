package test_case

import (
	"regexp"
	"strings"
	"test_case/parsers/errors"
)

var (
	testCasePattern = regexp.MustCompile(`(?msi)BEGIN(.+?)END`)
)

func Parse(testCases string) ([]TestCaseTransactionsIterator, error) {
	matches := testCasePattern.FindAllStringSubmatch(testCases, -1)
	if len(matches) == 0 {
		return nil, errors.NoTestCases
	}

	var testCasesIterators []TestCaseTransactionsIterator
	for _, match := range matches {
		testCasesIterators = append(
			testCasesIterators,
			TestCaseTransactionsIterator{
				testCase:     match[1],
				transactions: getMeaningRows(match[1]),
			},
		)
	}

	return testCasesIterators, nil
}

func getMeaningRows(testCase string) []string {
	var meaningRows []string
	for _, row := range strings.Split(testCase, "\n") {
		row = strings.TrimSpace(row)

		if shouldSkipRow(row) {
			continue
		}

		meaningRows = append(meaningRows, row)
	}

	return meaningRows
}

func shouldSkipRow(row string) bool {
	return row == "" || len(row) == 1 || row[:2] == "//"
}
