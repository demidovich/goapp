package errors

type Stacktracer interface {
	Stacktrace() *stacktrace
}

type Error struct {
	message    string
	stacktrace *stacktrace
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Stacktrace() *stacktrace {
	return e.stacktrace
}

func New(message string) *Error {
	return &Error{
		message:    message,
		stacktrace: newStacktrace(),
	}
}

func Wrap(err error) *Error {
	return &Error{
		message:    err.Error(),
		stacktrace: newStacktrace(),
	}
}
