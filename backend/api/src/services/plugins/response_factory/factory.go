package response_factory

import "api_meta/interfaces"

func DefaultResponse() interfaces.Response {
	return defaultResponse{}
}

func SuccessResponse(data interface{}) interfaces.Response {
	return successResponse{defaultResponse{data}}
}

func ErrorResponse(data interface{}) interfaces.Response {
	return errorResponse{defaultResponse{data}}
}
