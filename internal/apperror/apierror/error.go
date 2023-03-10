package apierror

import "errors"

type ApiError struct {
	err        error
	msg        string
	statuscode int
	errorcode  string
}

func New(msg string, statuscode int, errorcode string) *ApiError {
	return &ApiError{
		err:        errors.New(msg),
		msg:        msg,
		statuscode: statuscode,
		errorcode:  errorcode,
	}
}

func (e *ApiError) Error() string {
	return e.err.Error()
}

func (e *ApiError) StatusCode() int {
	return e.statuscode
}

func (e *ApiError) Message() string {
	return e.msg
}

func (e *ApiError) ErrorCode() string {
	return e.errorcode
}
