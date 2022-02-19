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

func NewError(code int, msg string) *CodeError {
	return &CodeError{
		code: code,
		err:  errors.New(msg),
	}
}

func NewCodeError(code int) *CodeError {
	return &CodeError{
		code: code,
	}
}

func (e *CodeError) Error() string {
	return e.err.Error()
}

func (e *CodeError) Data() *response {
	msg := ""
	if e.err != nil {
		msg = e.err.Error()
	}

	return &response{
		ErrCode: e.code,
		ErrMsg:  msg,
	}
}
