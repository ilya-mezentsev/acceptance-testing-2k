package data

type transactionTextContainer struct {
	transactionText string
	testCaseText    string
}

func (h transactionTextContainer) GetTransactionText() string {
	return h.transactionText
}

func (h *transactionTextContainer) SetTransactionText(text string) {
	h.transactionText = text
}

func (h *transactionTextContainer) SetTestCaseText(text string) {
	h.testCaseText = text
}

func (h transactionTextContainer) GetTestCaseText() string {
	return h.testCaseText
}
