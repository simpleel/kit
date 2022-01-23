package http

import (
	"net/http"

	"simpleel/kit/errors"
	"simpleel/kit/http/errcode"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// Unauthorized 输出未通过接口验证时的错误状态
func Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, errors.NewCodeError(errcode.ErrAccessDenyCode, http.StatusText(http.StatusUnauthorized)).Data())
}

// 错误处理
func ErrorHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case *errors.CodeError:
		return http.StatusOK, e.Data()
	default:
		codeErr := errors.NewCodeError(errcode.FailedCode, err.Error())
		return http.StatusOK, codeErr.Data()
	}
}
