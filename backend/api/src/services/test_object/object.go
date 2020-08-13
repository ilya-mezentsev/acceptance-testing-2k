package test_object

import (
	"api_meta/interfaces"
	"api_meta/models"
	"io"
	"services/errors"
	"services/plugins/hash"
	"services/plugins/logger"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

type service struct {
	logger     logger.CRUDEntityErrorsLogger
	repository interfaces.CRUDRepository
}

func New(repository interfaces.CRUDRepository) interfaces.CRUDService {
	return service{
		repository: repository,
		logger:     logger.CRUDEntityErrorsLogger{EntityName: entityName},
	}
}

func (s service) Create(request io.ReadCloser) interfaces.Response {
	var createTestObjectRequest models.CreateTestObjectRequest
	err := request_decoder.Decode(request, &createTestObjectRequest)
	if err != nil {
		s.logger.LogCreateEntityDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestObjectCode,
			Description: errors.DecodingRequestError,
		})
	}

	createTestObjectRequest.TestObject.Hash = hash.GetHashWithTimeAsKey(
		createTestObjectRequest.TestObject.Name,
	)
	if !validation.IsValid(&createTestObjectRequest) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestObjectCode,
			Description: errors.InvalidRequestError,
		})
	}

	err = s.repository.Create(createTestObjectRequest.AccountHash, map[string]interface{}{
		"name": createTestObjectRequest.TestObject.Name,
		"hash": createTestObjectRequest.TestObject.Hash,
	})
	if err != nil {
		s.logger.LogCreateEntityRepositoryError(err, map[string]interface{}{
			"account_hash": createTestObjectRequest.AccountHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestObjectCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s service) GetAll(accountHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestObjectsCode,
			Description: errors.InvalidRequestError,
		})
	}

	var testObjects []models.TestObject
	err := s.repository.GetAll(accountHash, &testObjects)
	if err != nil {
		s.logger.LogGetAllEntitiesRepositoryError(err, map[string]interface{}{
			"account_hash": accountHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestObjectsCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.SuccessResponse(testObjects)
}

func (s service) Get(accountHash, testObjectHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(testObjectHash) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestObjectCode,
			Description: errors.InvalidRequestError,
		})
	}

	var testObject models.TestObject
	err := s.repository.Get(accountHash, testObjectHash, &testObject)
	if err != nil {
		s.logger.LogGetEntityRepositoryError(err, map[string]interface{}{
			"account_hash":     accountHash,
			"test_object_hash": testObjectHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestObjectCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.SuccessResponse(testObject)
}

func (s service) Update(request io.ReadCloser) interfaces.Response {
	var updateTestObjectRequest models.UpdateRequest
	err := request_decoder.Decode(request, &updateTestObjectRequest)
	if err != nil {
		s.logger.LogUpdateEntityDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToUpdateTestObjectCode,
			Description: errors.DecodingRequestError,
		})
	}

	if !validation.IsValid(&updateTestObjectRequest) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToUpdateTestObjectCode,
			Description: errors.InvalidRequestError,
		})
	}

	err = s.repository.Update(updateTestObjectRequest.AccountHash, updateTestObjectRequest.UpdatePayload)
	if err != nil {
		s.logger.LogUpdateEntityRepositoryError(err, map[string]interface{}{
			"account_hash":   updateTestObjectRequest.AccountHash,
			"update_payload": updateTestObjectRequest.UpdatePayload,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToUpdateTestObjectCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s service) Delete(accountHash, testObjectHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(testObjectHash) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToDeleteTestObjectCode,
			Description: errors.InvalidRequestError,
		})
	}

	err := s.repository.Delete(accountHash, testObjectHash)
	if err != nil {
		s.logger.LogDeleteEntityRepositoryError(err, map[string]interface{}{
			"account_hash":     accountHash,
			"test_object_hash": testObjectHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToDeleteTestObjectCode,
			Description: errors.RepositoryError,
		})
	}

	return response_factory.DefaultResponse()
}
