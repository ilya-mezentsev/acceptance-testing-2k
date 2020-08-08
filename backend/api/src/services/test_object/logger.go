package test_object

import "plugins/logger"

func logCreateTestObjectDecodeError(err error) {
	logger.WarningF("Decode request body error while creating test object: %v", err)
}

func logCreateTestObjectRepositoryError(err error, context map[string]interface{}) {
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while creating test object: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}

func logGetAllTestObjectsRepositoryError(err error, context map[string]interface{}) {
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while fetching all test objects: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}

func logGetTestObjectRepositoryError(err error, context map[string]interface{}) {
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while fetching test object: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}

func logUpdateTestObjectDecodeError(err error) {
	logger.WarningF("Decode request body error while updating test object: %v", err)
}

func logUpdateTestObjectRepositoryError(err error, context map[string]interface{}) {
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while updating test object: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}

func logDeleteTestObjectRepositoryError(err error, context map[string]interface{}) {
	logger.WithFields(logger.Fields{
		MessageTemplate: "Repository error while deleting test object: %v",
		Args: []interface{}{
			err,
		},
		Optional: context,
	}, logger.Warning)
}
