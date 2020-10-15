package assertion

import "test_case/parsers/transaction/data"

var (
	MockDataScore10      = data.AssertionTransactionData{}
	MockDataArray        = data.AssertionTransactionData{}
	MockDataArrayWithMap = data.AssertionTransactionData{}
	MockDataTemplateNewValue = data.AssertionTransactionData{}
)

func init() {
	MockDataScore10.SetField("variableName", "response")
	MockDataScore10.SetField("dataPath", "data.score")
	MockDataScore10.SetField("newValue", "10")
	MockDataScore10.SetTestCaseText(`some test case text`)

	MockDataTemplateNewValue.SetField("variableName", "response")
	MockDataTemplateNewValue.SetField("dataPath", "data.score")
	MockDataTemplateNewValue.SetField("newValue", "${foo.bar}")
	MockDataTemplateNewValue.SetTestCaseText(`some test case text`)

	MockDataArray.SetField("variableName", "response")
	MockDataArray.SetField("dataPath", "data.1")
	MockDataArray.SetField("newValue", "2")
	MockDataArray.SetTestCaseText(`some test case text`)

	MockDataArrayWithMap.SetField("variableName", "response")
	MockDataArrayWithMap.SetField("dataPath", "data.0.y")
	MockDataArrayWithMap.SetField("newValue", "2")
	MockDataArrayWithMap.SetTestCaseText(`some test case text`)
}
