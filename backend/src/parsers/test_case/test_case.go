package test_case

import (
	"parsers/errors"
	"regexp"
	"strings"
)

var (
	testCasePattern = regexp.MustCompile(`(?ms){(.+?)}`)
)

type Parser struct {
	currentTestCase int
	testCases       [][]string
}

func (p *Parser) Init(testCases string) error {
	matches := testCasePattern.FindAllStringSubmatch(testCases, -1)
	if len(matches) == 0 {
		return errors.NoTestCases
	}

	p.testCases = nil
	for _, match := range matches {
		p.testCases = append(p.testCases, p.getMeaningRows(match[1]))
	}

	p.currentTestCase = 0
	return nil
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

func (p Parser) Done() bool {
	return p.currentTestCase == len(p.testCases)
}

func (p *Parser) NextTransactions() []string {
	testCaseTransactions := p.testCases[p.currentTestCase]
	p.currentTestCase++

	return testCaseTransactions
}
