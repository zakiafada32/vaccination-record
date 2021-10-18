package utils

import (
	"net/http"

	"github.com/zakiafada32/vaccination-record/business"
)

type ResponseFormat struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ConstructResponse(status string, data interface{}) (int, ResponseFormat) {
	var httpStatus int
	var response ResponseFormat

	switch status {
	// success response
	case business.Success:
		httpStatus = http.StatusOK
		response.Success = true
		response.Message = business.Success
		response.Data = data

	case business.SuccessCreated:
		httpStatus = http.StatusCreated
		response.Success = true
		response.Message = business.SuccessCreated
		response.Data = data

	case business.SuccessDeleted:
		httpStatus = http.StatusOK
		response.Success = true
		response.Message = business.SuccessDeleted
		response.Data = data

	// error response
	case business.BadRequest:
		httpStatus = http.StatusBadRequest
		response.Success = false
		response.Message = business.BadRequest
		response.Data = map[string]interface{}{}

	case business.InternalServerError:
		httpStatus = http.StatusInternalServerError
		response.Success = false
		response.Message = business.InternalServerError
		response.Data = map[string]interface{}{}

	case business.NotFound:
		httpStatus = http.StatusNotFound
		response.Success = false
		response.Message = business.NotFound
		response.Data = map[string]interface{}{}

	case business.EmailAlreadyExist:
		httpStatus = http.StatusBadRequest
		response.Success = false
		response.Message = business.EmailAlreadyExist
		response.Data = map[string]interface{}{}

	case business.PhoneNumberAlreadyExist:
		httpStatus = http.StatusBadRequest
		response.Success = false
		response.Message = business.PhoneNumberAlreadyExist
		response.Data = map[string]interface{}{}

	case business.IdentityCardNumberAlreadyExist:
		httpStatus = http.StatusBadRequest
		response.Success = false
		response.Message = business.IdentityCardNumberAlreadyExist
		response.Data = map[string]interface{}{}

	case business.PasswordIncorrect:
		httpStatus = http.StatusNotFound
		response.Success = false
		response.Message = business.PasswordIncorrect
		response.Data = map[string]interface{}{}
	}

	return httpStatus, response
}
