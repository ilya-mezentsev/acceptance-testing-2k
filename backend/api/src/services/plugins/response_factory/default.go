package response_factory

type defaultResponse struct {
	data interface{}
}

func (r defaultResponse) GetStatus() string {
	return statusOk
}

func (r defaultResponse) HasData() bool {
	return r.data != nil
}

func (r defaultResponse) GetData() interface{} {
	return r.data
}
