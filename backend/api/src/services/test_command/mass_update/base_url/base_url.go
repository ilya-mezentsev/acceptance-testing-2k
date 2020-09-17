package base_url

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

func (s Service) Update(accountHash string, request io.ReadCloser) interfaces.Response {
	var massBaseURLsUpdateRequest models.MassBaseURLsUpdateRequest
	err := request_decoder.Decode(request, &massBaseURLsUpdateRequest)
	if err != nil {
		s.logger.LogUpdateEntityDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassUpdateBaseURL,
			Description: errors.DecodingRequestError,
		})
	}

	if !validation.IsMd5Hash(accountHash) || !validation.IsValid(&massBaseURLsUpdateRequest) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassUpdateBaseURL,
			Description: errors.InvalidRequestError,
		})
	}

	err = s.repository.Update(
		accountHash,
		s.getUpdateModels(massBaseURLsUpdateRequest),
	)
	if err != nil {
		s.logger.LogUpdateEntityRepositoryError(err, map[string]interface{}{
			"mass_base_urls_update_request": massBaseURLsUpdateRequest,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToMassUpdateBaseURL,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) getUpdateModels(
	massBaseURLsUpdateRequest models.MassBaseURLsUpdateRequest,
) []models.UpdateModel {
	var updateModels []models.UpdateModel
	for _, updateTarget := range massBaseURLsUpdateRequest.CommandHashes {
		updateModels = append(updateModels, models.UpdateModel{
			Hash:      updateTarget.Hash,
			FieldName: "command_setting:base_url",
			NewValue:  massBaseURLsUpdateRequest.BaseURL,
		})
	}

	return updateModels
}
