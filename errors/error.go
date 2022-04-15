package errors

import (
	"errors"
	"strconv"

	"go.simpleel.com/kit/strings"
)

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
		err:  errors.New(strings.Concat("error code: ", strconv.Itoa(code))),
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
