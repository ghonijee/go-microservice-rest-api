package common

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  bool     `json:"status"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"error"`
}

func (e *ErrorResponse) toJson() string {
	json, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(json)
}

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (s *SuccessResponse) toJson() string {
	json, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(json)
}

func ErrorJson(w http.ResponseWriter, data ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Code)
	w.Write([]byte(data.toJson()))
}

func ErrorDataNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	data := ErrorResponse{
		Status:  false,
		Code:    http.StatusNotFound,
		Message: "Data Not Found",
	}
	w.Write([]byte(data.toJson()))
}

func ErrorBadRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	data := ErrorResponse{
		Status:  false,
		Code:    http.StatusBadRequest,
		Message: message,
	}
	w.Write([]byte(data.toJson()))
}

func ErrorInternalServer(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	data := ErrorResponse{
		Status:  false,
		Code:    http.StatusInternalServerError,
		Message: message,
	}
	w.Write([]byte(data.toJson()))
}

func ErrorUnauthorized(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	data := ErrorResponse{
		Status:  false,
		Code:    http.StatusUnauthorized,
		Message: message,
	}
	w.Write([]byte(data.toJson()))
}

func ErrorForbidden(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	data := ErrorResponse{
		Status:  false,
		Code:    http.StatusForbidden,
		Message: message,
	}
	w.Write([]byte(data.toJson()))
}

func ErrorInvalidData(w http.ResponseWriter, err ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	data := ErrorResponse{
		Status:  false,
		Code:    http.StatusUnprocessableEntity,
		Message: err.Message,
		Errors:  err.Errors,
	}
	w.Write([]byte(data.toJson()))
}

func SuccessJson(w http.ResponseWriter, data SuccessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	success := SuccessResponse{
		Status: true,
	}

	if data.Data != nil {
		success.Data = data.Data
	}

	if data.Code != 0 {
		success.Code = data.Code
	}

	if data.Message != "" {
		success.Message = data.Message
	}
	w.Write([]byte(success.toJson()))
}

func SuccessCreateData(w http.ResponseWriter, data SuccessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	success := SuccessResponse{
		Status: true,
		Code:   http.StatusCreated,
	}

	if data.Data != nil {
		success.Data = data.Data
	}

	if data.Message != "" {
		success.Message = data.Message
	} else {
		success.Message = "Create data success"
	}

	w.Write([]byte(success.toJson()))
}

func SuccessUpdateData(w http.ResponseWriter, data SuccessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	success := SuccessResponse{
		Status: true,
		Code:   http.StatusOK,
	}

	if data.Data != nil {
		success.Data = data.Data
	}

	if data.Message != "" {
		success.Message = data.Message
	} else {
		success.Message = "Update data success"
	}

	w.Write([]byte(success.toJson()))
}
