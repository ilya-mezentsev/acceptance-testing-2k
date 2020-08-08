package response_factory

type errorResponse struct {
	defaultResponse
}

func (r errorResponse) GetStatus() string {
	return statusError
}
