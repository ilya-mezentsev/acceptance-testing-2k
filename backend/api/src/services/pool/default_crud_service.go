package pool

import (
	"api_meta/interfaces"
	"io"
	"services/errors"
	"services/plugins/response_factory"
)

type defaultCRUDService struct {
	desc string
}

func (s defaultCRUDService) Create(string, io.ReadCloser) interfaces.Response {
	return response_factory.ErrorResponse(errors.ServiceError{
		Code:        noServiceErrorCode,
		Description: s.desc,
	})
}

func (s defaultCRUDService) GetAll(string) interfaces.Response {
	return response_factory.ErrorResponse(errors.ServiceError{
		Code:        noServiceErrorCode,
		Description: s.desc,
	})
}

func (s defaultCRUDService) Get(string, string) interfaces.Response {
	return response_factory.ErrorResponse(errors.ServiceError{
		Code:        noServiceErrorCode,
		Description: s.desc,
	})
}

func (s defaultCRUDService) Update(string, io.ReadCloser) interfaces.Response {
	return response_factory.ErrorResponse(errors.ServiceError{
		Code:        noServiceErrorCode,
		Description: s.desc,
	})
}

func (s defaultCRUDService) Delete(string, string) interfaces.Response {
	return response_factory.ErrorResponse(errors.ServiceError{
		Code:        noServiceErrorCode,
		Description: s.desc,
	})
}
