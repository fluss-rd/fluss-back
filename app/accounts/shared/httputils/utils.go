package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var internalServerError = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: "internal server error",
}

// ErrorResponse error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

// NewBadRequestError returns a bad request error response
func NewBadRequestError(msg string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "bad request: " + msg,
	}
}

// NewNotFoundError returns an ErrorResponse with the not found values
func NewNotFoundError(resourceName string) ErrorResponse {
	message := "not found"
	if resourceName != "" {
		message = fmt.Sprintf("%s %s", resourceName, message)
	}

	return ErrorResponse{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// RespondJSON responds with a json with the given status code and data
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

// RespondWithError responds with a json with the given status code and message
func RespondWithError(w http.ResponseWriter, err error) {
	errorResponse, ok := err.(ErrorResponse)
	if !ok {
		RespondInternalServerError(w)
		return
	}

	RespondJSON(w, errorResponse.Code, errorResponse)
}

// RespondInternalServerError responds with an internal server error response
func RespondInternalServerError(w http.ResponseWriter) {
	RespondJSON(w, internalServerError.Code, internalServerError)
}
