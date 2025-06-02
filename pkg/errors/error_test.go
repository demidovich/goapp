package errors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacktraceCaller(t *testing.T) {
	err := methodC()
	caller := err.Stacktrace().Caller()

	// Выделение из goapp-boilerplate/pkg/errors.methodA названия метода methodA
	shortFunctionName := func(name string) string {
		return strings.Split(name, ".")[1]
	}

	assert.Equal(t, "methodA", shortFunctionName(caller.Function))
}

func methodC() Error {
	return methodB()
}

func methodB() Error {
	return methodA()
}

func methodA() Error {
	return New("methodA error")
}
