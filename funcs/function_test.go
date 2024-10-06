package funcs_test

import (
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/commons/funcs"
	"code.gopub.tech/commons/values"
)

func TestIdentidy(t *testing.T) {
	assert.Equal(t, funcs.Identidy(1), 1)
}

func TestNot(t *testing.T) {
	assert.True(t, values.IsZero(0))
	assert.True(t, funcs.Not(values.IsZero[int])(1))
}
