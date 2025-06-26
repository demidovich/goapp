package testfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Testgen(t *testing.T) {
	e := UniqueEmail()

	assert.NotEmpty(t, e)
}
