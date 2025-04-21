package httper

import "net/http"

type HttpCodeError struct {
	Code    int
	Message string
}

func (this HttpCodeError) Error() string {
	return this.Message
}

func NewHttpError(code int, msg string) HttpCodeError {
	return HttpCodeError{
		Code:    code,
		Message: msg,
	}
}

func NewHttpErrorBadRequest(msg string) HttpCodeError {
	return NewHttpError(http.StatusBadRequest, msg)
}
