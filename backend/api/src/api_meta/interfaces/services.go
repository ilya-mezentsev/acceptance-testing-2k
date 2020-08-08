package interfaces

import (
	"io"
	"net/http"
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

	// Should be used for cookies and headers
	HTTPMetaService interface {
		PostProcess(r *http.Request, serviceResponse Response) error
	}

	CRUDServicesPool interface {
		Get(serviceType string) (CRUDService, Response)
	}

	MetaServicesPool interface {
		Get(serviceType string) HTTPMetaService
	}
)
