package pool

import (
	"api_meta/interfaces"
)

type MetaServicesPool struct {
	serviceTypeToImplementation map[string]interfaces.HTTPMetaService
}

func NewMeta() MetaServicesPool {
	return MetaServicesPool{map[string]interfaces.HTTPMetaService{}}
}

func (p MetaServicesPool) AddService(serviceType string, implementation interfaces.HTTPMetaService) {
	p.serviceTypeToImplementation[serviceType] = implementation
}

func (p MetaServicesPool) Get(serviceType string) interfaces.HTTPMetaService {
	return p.serviceTypeToImplementation[serviceType]
}
