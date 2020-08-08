package pool

import (
	"api_meta/interfaces"
	"fmt"
	"plugins/logger"
	"services/errors"
	"services/plugins/response_factory"
)

const (
	noServiceErrorCode = "no-service-for-operation-type"
)

type CRUDServicesPool struct {
	serviceTypeToImplementation map[string]interfaces.CRUDService
}

func NewCRUD() CRUDServicesPool {
	return CRUDServicesPool{map[string]interfaces.CRUDService{}}
}

func (p CRUDServicesPool) AddService(serviceType string, implementation interfaces.CRUDService) {
	p.serviceTypeToImplementation[serviceType] = implementation
}

func (p CRUDServicesPool) Get(serviceType string) (interfaces.CRUDService, interfaces.Response) {
	service, found := p.serviceTypeToImplementation[serviceType]
	if !found {
		desc := p.getNoServiceErrorDescription(serviceType)
		logger.Warning(desc)

		return nil, response_factory.ErrorResponse(errors.ServiceError{
			Code:        noServiceErrorCode,
			Description: desc,
		})
	}

	return service, nil
}

func (p CRUDServicesPool) getNoServiceErrorDescription(serviceType string) string {
	return fmt.Sprintf("Unable to find service for operation type: %s", serviceType)
}
