package http

import (
	"net/http"

	"code.simpleel.com/simpleel/kit/errors"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// Unauthorized 输出未通过接口验证时的错误状态
func Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, errors.NewCodeMsg(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)).Data())
}

// Write 输出 res 对象或错误消息
func Write(w http.ResponseWriter, res interface{}, err error) {
	if err != nil {
		if e, ok := err.(errors.CodeError); ok {
			httpx.OkJson(w, e.Data())
		} else {
			httpx.Error(w, err)
		}
	} else {
		httpx.OkJson(w, res)
	}
}
