package logger

import "logger"

type CRUDEntityErrorsLogger struct {
	EntityName string
}

func (l CRUDEntityErrorsLogger) LogCreateEntityDecodeError(err error) {
	logger.WarningF(
		"Decode request body error while creating entity (%s): %v", err, l.EntityName,
	)
}

func (l CRUDEntityErrorsLogger) LogCreateEntityRepositoryError(
	err error,
	context map[string]interface{},
) {
	context["entity_name"] = l.EntityName
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while creating entity: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}

func (l CRUDEntityErrorsLogger) LogGetAllEntitiesRepositoryError(
	err error,
	context map[string]interface{},
) {
	context["entity_name"] = l.EntityName
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while fetching all entities: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}

func (l CRUDEntityErrorsLogger) LogGetEntityRepositoryError(
	err error,
	context map[string]interface{},
) {
	context["entity_name"] = l.EntityName
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while fetching entity: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}

func (l CRUDEntityErrorsLogger) LogUpdateEntityDecodeError(err error) {
	logger.WarningF(
		"Decode request body error while updating entity (%s): %v", err, l.EntityName,
	)
}

func (l CRUDEntityErrorsLogger) LogUpdateEntityRepositoryError(
	err error,
	context map[string]interface{},
) {
	context["entity_name"] = l.EntityName
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while updating entity: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}

func (l CRUDEntityErrorsLogger) LogDeleteEntityRepositoryError(
	err error,
	context map[string]interface{},
) {
	context["entity_name"] = l.EntityName
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while deleting entity: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}
