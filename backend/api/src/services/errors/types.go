package errors

type (
	ServiceError struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	}
)
