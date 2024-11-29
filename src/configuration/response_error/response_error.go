package response_error

import "net/http"

type ResponseError struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ResponseError) Error() string {
	return r.Message
}

// constructor
func NewResponseErr(message, err string, code int, causes []Cause) *ResponseError {
	return &ResponseError{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func BadRequestErr(message string) *ResponseError {
	return &ResponseError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NotFoundErr(message string) *ResponseError {
	return &ResponseError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func ValidationErr(message string, causes []Cause) *ResponseError {
	return &ResponseError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func InternalServerErr(message string) *ResponseError {
	return &ResponseError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}
