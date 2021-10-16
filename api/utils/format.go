package utils

import (
	"net/http"

	"github.com/zakiafada32/vaccine/business"
)

type ResponseFormat struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ConstructResponse(status string, data map[string]interface{}) (int, ResponseFormat) {
	var httpStatus int
	var response ResponseFormat

	switch status {
	// success response
	case business.Success:
		httpStatus = http.StatusOK
		response.Success = true
		response.Message = "success"
		response.Data = data
	case business.SuccessCreated:
		httpStatus = http.StatusCreated
		response.Success = true
		response.Message = "success created"
		response.Data = data
	// error response
	case business.BadRequest:
		httpStatus = http.StatusBadRequest
		response.Success = false
		response.Message = "bad request"
		response.Data = map[string]interface{}{}
	case business.InternalServerError:
		httpStatus = http.StatusInternalServerError
		response.Success = false
		response.Message = "internal server error"
		response.Data = map[string]interface{}{}
	case business.NotFound:
		httpStatus = http.StatusNotFound
		response.Success = false
		response.Message = "data not found"
		response.Data = map[string]interface{}{}
	}
	return httpStatus, response
}
