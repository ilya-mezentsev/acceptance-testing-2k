package services

import (
	"api_meta/interfaces"
	"io"
	"net/http"
)

type CRUDServiceMock struct {
}

func (m CRUDServiceMock) Create(io.ReadCloser) interfaces.Response {
	panic("implement me")
}

func (m CRUDServiceMock) GetAll(string) interfaces.Response {
	panic("implement me")
}

func (m CRUDServiceMock) Get(string, string) interfaces.Response {
	panic("implement me")
}

func (m CRUDServiceMock) Update(io.ReadCloser) interfaces.Response {
	panic("implement me")
}

func (m CRUDServiceMock) Delete(string, string) interfaces.Response {
	panic("implement me")
}

type MetaServiceMock struct {
}

func (m MetaServiceMock) PostProcess(*http.Request, interfaces.Response) error {
	panic("implement me")
}
