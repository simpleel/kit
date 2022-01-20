package errors

import "errors"

type CodeError struct {
	code int
	err  error
}

type response struct {
	ErrCode int    `json:"code"`
	ErrMsg  string `json:"message,omitempty"`
}

func New(err error) CodeError {
	return CodeError{
		err: err,
	}
}

func NewError(msg string) CodeError {
	return CodeError{
		err: errors.New(msg),
	}
}

func NewCodeError(code int, msg string) CodeError {
	return CodeError{
		code: code,
		err:  errors.New(msg),
	}
}

func (e CodeError) Error() string {
	return e.err.Error()
}

func (e CodeError) Raw() error {
	return e.err
}

func (e CodeError) Data() response {
	return response{
		ErrCode: e.code,
		ErrMsg:  e.err.Error(),
	}
}
