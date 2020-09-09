package timeout

import (
	"api_meta/interfaces"
	"api_meta/models"
	"io"
	"services/errors"
	"services/plugins/logger"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

type Service struct {
	logger     logger.CRUDEntityErrorsLogger
	repository interfaces.UpdateRepository
}

func New(repository interfaces.UpdateRepository) Service {
	return Service{
		logger:     logger.CRUDEntityErrorsLogger{EntityName: entityName},
		repository: repository,
	}
}

func (s Service) Update(request io.ReadCloser) interfaces.Response {
	var massTimeoutsUpdateRequest models.MassTimeoutsUpdateRequest
	err := request_decoder.Decode(request, &massTimeoutsUpdateRequest)
	if err != nil {
		s.logger.LogUpdateEntityDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassUpdateTimeout,
			Description: errors.DecodingRequestError,
		})
	}

	if !validation.IsValid(&massTimeoutsUpdateRequest) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassUpdateTimeout,
			Description: errors.InvalidRequestError,
		})
	}

	err = s.repository.Update(
		massTimeoutsUpdateRequest.AccountHash,
		s.getUpdateModels(massTimeoutsUpdateRequest),
	)
	if err != nil {
		s.logger.LogUpdateEntityRepositoryError(err, map[string]interface{}{
			"mass_base_urls_update_request": massTimeoutsUpdateRequest,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassUpdateTimeout,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) getUpdateModels(
	massTimeoutsUpdateRequest models.MassTimeoutsUpdateRequest,
) []models.UpdateModel {
	var updateModels []models.UpdateModel
	for _, updateTarget := range massTimeoutsUpdateRequest.CommandHashes {
		updateModels = append(updateModels, models.UpdateModel{
			Hash:      updateTarget.Hash,
			FieldName: "command_setting:timeout",
			NewValue:  massTimeoutsUpdateRequest.Timeout,
		})
	}

	return updateModels
}
