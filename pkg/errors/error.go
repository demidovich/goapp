package errors

import (
	"fmt"
	"io"
	"strings"
)

func New(message string) Error {
	return Error{
		message:    message,
		stacktrace: newStacktrace(3),
	}
}

func Wrap(err error) Error {
	return Error{
		message:    err.Error(),
		stacktrace: newStacktrace(3),
	}
}

type Stacktracer interface {
	Stacktrace() stacktrace
}

type Error struct {
	message    string
	stacktrace stacktrace
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Stacktrace() stacktrace {
	return e.stacktrace
}

func (e Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			b := strings.Builder{}
			b.WriteString("\n\nStack Trace:")
			for _, frame := range e.Stacktrace().Frames() {
				b.WriteString(
					fmt.Sprintf("\n-> %s:%d (%s)", frame.File, frame.Line, frame.Function),
				)
			}
			io.WriteString(s, e.message)
			io.WriteString(s, b.String())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.message)
	case 'q':
		fmt.Fprintf(s, "%q", e.message)
	}
}
