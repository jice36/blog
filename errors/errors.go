package errors

type httpError struct{
	err error
	errS string
}

func (e *httpError) idNotTransferred(){}

