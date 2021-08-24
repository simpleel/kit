package http

import (
	"strings"

	"code.simpleel.com/simpleel/kit/errors"
	"code.simpleel.com/simpleel/kit/http/errcode"
)

// NewInternalServerError 返回一个服务器内部错误
func NewInternalServerError(msg ...string) error {
	s := combineMsg(msg...)
	return errors.NewCodeMsg(errcode.ErrServerCode, s)
}

// NewOprationFaildError 返回一个操作失败的错误
func NewOprationFaildError(msg ...string) error {
	s := combineMsg(msg...)
	return errors.NewCodeMsg(errcode.FailedCode, s)
}

// NewAccessDenyError 返回一个越权错误
func NewAccessDenyError(msg ...string) error {
	s := combineMsg(msg...)
	return errors.NewCodeMsg(errcode.ErrAccessDenyCode, s)
}

func combineMsg(msg ...string) string {
	if len(msg) > 1 {
		return strings.Join(msg, " ")
	} else if len(msg) == 1 {
		return msg[0]
	}
	return ""
}
