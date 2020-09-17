package controllers

import (
	"api_meta/interfaces"
	"io"
	"io/ioutil"
	"services/plugins/response_factory"
)

type (
	CRUDServiceMock struct {
		CalledWith map[string]interface{}
	}

	Response struct {
		Code string `json:"code"`
	}
)

func (m *CRUDServiceMock) Reset() {
	m.CalledWith = map[string]interface{}{}
}

func (m *CRUDServiceMock) Create(_ string, request io.ReadCloser) interfaces.Response {
	d, _ := ioutil.ReadAll(request)
	data := string(d)
	m.CalledWith["Create"] = data
	if data == BadRequestData {
		return response_factory.ErrorResponse(Response{Code: ErrorResponseCode})
	}

	return response_factory.DefaultResponse()
}

func (m *CRUDServiceMock) GetAll(accountHash string) interfaces.Response {
	m.CalledWith["GetAll"] = accountHash

	if accountHash == BadAccountHash {
		return response_factory.ErrorResponse(Response{Code: ErrorResponseCode})
	}

	return response_factory.DefaultResponse()
}

func (m *CRUDServiceMock) Get(accountHash, entityHash string) interfaces.Response {
	m.CalledWith["Get"] = map[string]string{
		"account_hash": accountHash,
		"entity_hash":  entityHash,
	}

	if accountHash == BadAccountHash || entityHash == BadEntityHash {
		return response_factory.ErrorResponse(Response{Code: ErrorResponseCode})
	}

	return response_factory.DefaultResponse()
}

func (m *CRUDServiceMock) Update(_ string, request io.ReadCloser) interfaces.Response {
	d, _ := ioutil.ReadAll(request)
	data := string(d)
	m.CalledWith["Update"] = data
	if data == BadRequestData {
		return response_factory.ErrorResponse(Response{Code: ErrorResponseCode})
	}

	return response_factory.DefaultResponse()
}

func (m *CRUDServiceMock) Delete(accountHash, entityHash string) interfaces.Response {
	m.CalledWith["Delete"] = map[string]string{
		"account_hash": accountHash,
		"entity_hash":  entityHash,
	}

	if accountHash == BadAccountHash || entityHash == BadEntityHash {
		return response_factory.ErrorResponse(Response{Code: ErrorResponseCode})
	}

	return response_factory.DefaultResponse()
}
