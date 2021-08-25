package errors

import "errors"

type CodeError struct {
	code int
	err  error
}

type response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func New(err error) CodeError {
	return CodeError{
		err: err,
	}
}

func NewMsg(msg string) CodeError {
	return CodeError{
		err: errors.New(msg),
	}
}

func NewCodeMsg(code int, msg string) CodeError {
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
