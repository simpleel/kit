package http

import (
	"code.simpleel.com/simpleel/kit/errors"
	"code.simpleel.com/simpleel/kit/http/errcode"
)

// NewInternalServerError 返回一个服务器内部错误
func NewInternalServerError(msg string) error {
	return errors.NewCodeMsg(errcode.ErrServerCode, msg)
}

// NewOprationFaildError 返回一个操作失败的错误
func NewOprationFaildError(msg string) error {
	return errors.NewCodeMsg(errcode.FailedCode, msg)
}
