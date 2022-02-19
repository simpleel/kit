package http

import (
	"net/http"

	"go.simpleel.com/kit/errors"
	"go.simpleel.com/kit/http/errcode"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Unauthorized 输出未通过接口验证时的错误状态
func Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, errors.NewError(errcode.ErrAccessDenyCode, http.StatusText(http.StatusUnauthorized)).Data())
}

// 错误处理
func ErrorHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case *errors.CodeError:
		return http.StatusOK, e.Data()
	default:
		codeErr := errors.NewError(errcode.FailedCode, err.Error())
		return http.StatusOK, codeErr.Data()
	}
}
