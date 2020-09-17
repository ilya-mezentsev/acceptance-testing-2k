package services

import (
	"api_meta/interfaces"
	"io"
)

type CRUDServiceMock struct {
}

func (m CRUDServiceMock) Create(string, io.ReadCloser) interfaces.Response {
	panic("implement me")
}

func (m CRUDServiceMock) GetAll(string) interfaces.Response {
	panic("implement me")
}

func (m CRUDServiceMock) Get(string, string) interfaces.Response {
	panic("implement me")
}

func (m CRUDServiceMock) Update(string, io.ReadCloser) interfaces.Response {
	panic("implement me")
}

func (m CRUDServiceMock) Delete(string, string) interfaces.Response {
	panic("implement me")
}
