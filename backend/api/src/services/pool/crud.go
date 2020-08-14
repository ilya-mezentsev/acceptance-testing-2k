package pool

import (
	"api_meta/interfaces"
	"fmt"
	"logger"
)

type CRUDServicesPool struct {
	serviceTypeToImplementation map[string]interfaces.CRUDService
}

func New() CRUDServicesPool {
	return CRUDServicesPool{map[string]interfaces.CRUDService{}}
}

func (p CRUDServicesPool) AddService(serviceType string, implementation interfaces.CRUDService) {
	p.serviceTypeToImplementation[serviceType] = implementation
}

func (p CRUDServicesPool) Get(serviceType string) interfaces.CRUDService {
	service, found := p.serviceTypeToImplementation[serviceType]
	if !found {
		desc := p.getNoServiceErrorDescription(serviceType)
		logger.Warning(desc)

		return defaultCRUDService{desc}
	}

	return service
}

func (p CRUDServicesPool) getNoServiceErrorDescription(serviceType string) string {
	return fmt.Sprintf("Unable to find service for operation type: %s", serviceType)
}
