package errors

type CodeError struct {
	code int
	err  error
}

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, err error) CodeError {
	return CodeError{
		code: code,
		err:  err,
	}
}

func (e CodeError) Error() error {
	return e.err
}

func (e CodeError) Data() response {
	return response{
		Code: e.code,
		Msg:  e.err.Error(),
	}
}
