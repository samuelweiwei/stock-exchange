package errorx

type CodeError struct {
	Code    int
	Message string
	Args    []interface{}
}

func (e CodeError) Error() string {
	return e.Message
}

func NewWithCode(code int, args ...interface{}) CodeError {
	return CodeError{Code: code, Args: args}
}

func New(code int, msg string) CodeError {
	return CodeError{Code: code, Message: msg}
}
