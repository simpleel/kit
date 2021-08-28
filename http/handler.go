package http

import (
	"net/http"

	"code.simpleel.com/simpleel/kit/errors"
	"code.simpleel.com/simpleel/kit/http/errcode"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// Unauthorized 输出未通过接口验证时的错误状态
func Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, errors.NewCodeMsg(errcode.ErrAccessDenyCode, http.StatusText(http.StatusUnauthorized)).Data())
}

// Write 输出 res 对象或错误消息
func Write(w http.ResponseWriter, res interface{}, err error) {
	if err != nil {
		e, ok := err.(errors.CodeError)
		if !ok {
			e = errors.NewCodeMsg(errcode.FailedCode, err.Error())
		}
		httpx.OkJson(w, e.Data())
	} else {
		httpx.OkJson(w, res)
	}
}
