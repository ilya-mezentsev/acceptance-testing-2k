package data

type transactionTextContainer struct {
	transactionText string
}

func (h transactionTextContainer) GetTransactionText() string {
	return h.transactionText
}

func (h *transactionTextContainer) SetTransactionText(text string) {
	h.transactionText = text
}
