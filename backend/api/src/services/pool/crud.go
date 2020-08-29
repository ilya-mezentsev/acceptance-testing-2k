package pool

import (
	"api_meta/interfaces"
	"fmt"
	"logger"
)

const (
	CreateServiceOperationType = "create"
	ReadServiceOperationType   = "read"
	UpdateServiceOperationType = "update"
	DeleteServiceOperationType = "delete"
)

type CRUDServicesPool struct {
	entityTypeToCreateService map[string]interfaces.CreateService
	entityTypeToReadService   map[string]interfaces.ReadService
	entityTypeToUpdateService map[string]interfaces.UpdateService
	entityTypeToDeleteService map[string]interfaces.DeleteService
}

func New() CRUDServicesPool {
	return CRUDServicesPool{
		entityTypeToCreateService: map[string]interfaces.CreateService{},
		entityTypeToReadService:   map[string]interfaces.ReadService{},
		entityTypeToUpdateService: map[string]interfaces.UpdateService{},
		entityTypeToDeleteService: map[string]interfaces.DeleteService{},
	}
}

func (p CRUDServicesPool) AddService(
	entityType string,
	operationTypes []string,
	service interface{},
) {
	for _, operationType := range operationTypes {
		switch operationType {
		case CreateServiceOperationType:
			p.AddCreateService(entityType, service.(interfaces.CreateService))
		case ReadServiceOperationType:
			p.AddReadService(entityType, service.(interfaces.ReadService))
		case UpdateServiceOperationType:
			p.AddUpdateService(entityType, service.(interfaces.UpdateService))
		case DeleteServiceOperationType:
			p.AddDeleteService(entityType, service.(interfaces.DeleteService))
		default:
			panic(fmt.Sprintf("Unexpected operation type: %s", operationType))
		}
	}
}

func (p CRUDServicesPool) AddCRUDService(entityType string, service interfaces.CRUDService) {
	p.AddCreateService(entityType, service)
	p.AddReadService(entityType, service)
	p.AddUpdateService(entityType, service)
	p.AddDeleteService(entityType, service)
}

func (p CRUDServicesPool) AddCreateService(entityType string, service interfaces.CreateService) {
	p.entityTypeToCreateService[entityType] = service
}

func (p CRUDServicesPool) GetCreateService(entityType string) interfaces.CreateService {
	src := make(map[string]interface{})
	for t, service := range p.entityTypeToCreateService {
		src[t] = service
	}

	return p.getService(entityType, src).(interfaces.CreateService)
}

func (p CRUDServicesPool) AddReadService(entityType string, service interfaces.ReadService) {
	p.entityTypeToReadService[entityType] = service
}

func (p CRUDServicesPool) GetReadService(entityType string) interfaces.ReadService {
	src := make(map[string]interface{})
	for t, service := range p.entityTypeToReadService {
		src[t] = service
	}

	return p.getService(entityType, src).(interfaces.ReadService)
}

func (p CRUDServicesPool) AddUpdateService(entityType string, service interfaces.UpdateService) {
	p.entityTypeToUpdateService[entityType] = service
}

func (p CRUDServicesPool) GetUpdateService(entityType string) interfaces.UpdateService {
	src := make(map[string]interface{})
	for t, service := range p.entityTypeToUpdateService {
		src[t] = service
	}

	return p.getService(entityType, src).(interfaces.UpdateService)
}

func (p CRUDServicesPool) AddDeleteService(entityType string, service interfaces.DeleteService) {
	p.entityTypeToDeleteService[entityType] = service
}

func (p CRUDServicesPool) GetDeleteService(entityType string) interfaces.DeleteService {
	src := make(map[string]interface{})
	for t, service := range p.entityTypeToDeleteService {
		src[t] = service
	}

	return p.getService(entityType, src).(interfaces.DeleteService)
}

func (p CRUDServicesPool) getService(entityType string, src map[string]interface{}) interface{} {
	service, found := src[entityType]
	if !found {
		desc := p.getNoServiceErrorDescription(entityType)
		logger.Warning(desc)

		return defaultCRUDService{desc}
	}

	return service
}

func (p CRUDServicesPool) getNoServiceErrorDescription(serviceType string) string {
	return fmt.Sprintf("Unable to find service for operation type: %s", serviceType)
}
