package errors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacktrace(t *testing.T) {
	err := methodC()
	stack := err.Stacktrace().ToJson()

	// Выделение из goapp-boilerplate/pkg/errors.methodA названия метода methodA
	shortFunctionName := func(name string) string {
		return strings.Split(name, ".")[1]
	}

	assert.Equal(t, "methodA", shortFunctionName(stack[0].Function))
	assert.Equal(t, "methodB", shortFunctionName(stack[1].Function))
	assert.Equal(t, "methodC", shortFunctionName(stack[2].Function))
}

func methodC() *Error {
	return methodB()
}

func methodB() *Error {
	return methodA()
}

func methodA() *Error {
	return New("methodA error")
}
