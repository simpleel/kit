package errors

import "net/http"

type CodeError struct {
	code int
	err  string
}

type response struct {
	ResultCode int    `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
}

func HTTPErrorHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case CodeError:
		return http.StatusOK, e.Data()
	default:
		return http.StatusInternalServerError, nil
	}
}
func New(err string) CodeError {
	return CodeError{
		code: http.StatusInternalServerError,
		err:  err,
	}
}

func NewCodeError(code int, err string) CodeError {
	return CodeError{
		code: code,
		err:  err,
	}
}

func (e CodeError) Error() string {
	return e.err
}

func (e CodeError) Data() response {
	return response{
		ResultCode: e.code,
		ResultMsg:  e.err,
	}
}
