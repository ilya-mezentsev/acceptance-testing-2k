package test_object

import (
	"api_meta/interfaces"
	"api_meta/models"
	"io"
	"services/errors"
	"services/plugins/hash"
	"services/plugins/request_decoder"
	"services/plugins/response_factory"
	"services/plugins/validation"
)

type service struct {
	repository interfaces.CRUDRepository
}

func New(repository interfaces.CRUDRepository) interfaces.CRUDService {
	return service{repository: repository}
}

func (s service) Create(request io.ReadCloser) interfaces.Response {
	var createTestObjectRequest models.CreateTestObjectRequest
	err := request_decoder.Decode(request, &createTestObjectRequest)
	if err != nil {
		logCreateTestObjectDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestObjectCode,
			Description: decodingRequestError,
		})
	}

	createTestObjectRequest.TestObject.Hash = hash.GetHashWithTimeAsKey(
		createTestObjectRequest.TestObject.Name,
	)
	if !validation.IsValid(&createTestObjectRequest) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestObjectCode,
			Description: invalidRequestError,
		})
	}

	err = s.repository.Create(createTestObjectRequest.AccountHash, map[string]interface{}{
		"name": createTestObjectRequest.TestObject.Name,
		"hash": createTestObjectRequest.TestObject.Hash,
	})
	if err != nil {
		logCreateTestObjectRepositoryError(err, map[string]interface{}{
			"account_hash": createTestObjectRequest.AccountHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToCreateTestObjectCode,
			Description: repositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s service) GetAll(accountHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestObjectsCode,
			Description: invalidRequestError,
		})
	}

	var testObjects []models.TestObject
	err := s.repository.GetAll(accountHash, &testObjects)
	if err != nil {
		logGetAllTestObjectsRepositoryError(err, map[string]interface{}{
			"account_hash": accountHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestObjectsCode,
			Description: repositoryError,
		})
	}

	return response_factory.SuccessResponse(testObjects)
}

func (s service) Get(accountHash, testObjectHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(testObjectHash) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestObjectCode,
			Description: invalidRequestError,
		})
	}

	var testObject models.TestObject
	err := s.repository.Get(accountHash, testObjectHash, &testObject)
	if err != nil {
		logGetTestObjectRepositoryError(err, map[string]interface{}{
			"account_hash":     accountHash,
			"test_object_hash": testObjectHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToFetchTestObjectCode,
			Description: repositoryError,
		})
	}

	return response_factory.SuccessResponse(testObject)
}

func (s service) Update(request io.ReadCloser) interfaces.Response {
	var updateTestObjectRequest models.UpdateTestObjectRequest
	err := request_decoder.Decode(request, &updateTestObjectRequest)
	if err != nil {
		logUpdateTestObjectDecodeError(err)

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToUpdateTestObjectCode,
			Description: decodingRequestError,
		})
	}

	if !validation.IsValid(&updateTestObjectRequest) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToUpdateTestObjectCode,
			Description: invalidRequestError,
		})
	}

	err = s.repository.Update(updateTestObjectRequest.AccountHash, updateTestObjectRequest.UpdatePayload)
	if err != nil {
		logUpdateTestObjectRepositoryError(err, map[string]interface{}{
			"account_hash":   updateTestObjectRequest.AccountHash,
			"update_payload": updateTestObjectRequest.UpdatePayload,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToUpdateTestObjectCode,
			Description: repositoryError,
		})
	}

	return response_factory.DefaultResponse()
}

func (s service) Delete(accountHash, testObjectHash string) interfaces.Response {
	if !validation.IsMd5Hash(accountHash) || !validation.IsMd5Hash(testObjectHash) {
		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToDeleteTestObjectCode,
			Description: invalidRequestError,
		})
	}

	err := s.repository.Delete(accountHash, testObjectHash)
	if err != nil {
		logDeleteTestObjectRepositoryError(err, map[string]interface{}{
			"account_hash":     accountHash,
			"test_object_hash": testObjectHash,
		})

		return response_factory.ErrorResponse(errors.ServiceError{
			Code:        unableToDeleteTestObjectCode,
			Description: repositoryError,
		})
	}

	return response_factory.DefaultResponse()
}
