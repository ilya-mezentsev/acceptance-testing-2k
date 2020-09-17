package interfaces

import (
	"io"
)

type (
	Response interface {
		GetStatus() string
		HasData() bool
		GetData() interface{}
	}

	CreateService interface {
		Create(accountHash string, request io.ReadCloser) Response
	}

	ReadService interface {
		GetAll(accountHash string) Response
		Get(accountHash, entityHash string) Response
	}

	UpdateService interface {
		Update(accountHash string, request io.ReadCloser) Response
	}

	DeleteService interface {
		Delete(accountHash, entityHash string) Response
	}

	CRUDService interface {
		CreateService
		ReadService
		UpdateService
		DeleteService
	}

	CRUDServicesPool interface {
		GetCreateService(entityType string) CreateService
		GetReadService(entityType string) ReadService
		GetUpdateService(entityType string) UpdateService
		GetDeleteService(entityType string) DeleteService
	}
)
