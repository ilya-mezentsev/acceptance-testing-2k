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

	CRUDService interface {
		Create(request io.ReadCloser) Response
		GetAll(accountHash string) Response
		Get(accountHash, entityHash string) Response
		Update(request io.ReadCloser) Response
		Delete(accountHash, entityHash string) Response
	}

	CRUDServicesPool interface {
		Get(serviceType string) CRUDService
	}
)
