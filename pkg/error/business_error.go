package error

type BusinessError struct {
	msg string
}

func NewBusinessError(errorMsg string) *BusinessError {
	return &BusinessError{
		msg: errorMsg,
	}
}

func (e *BusinessError) Error() string {
	return e.msg
}
