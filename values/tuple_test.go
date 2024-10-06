package values_test

import (
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/commons/values"
)

func TestTuple(t *testing.T) {
	pair := values.Make2("age", 18)
	assert.Equal(t, pair.Val1, "age")
	assert.Equal(t, pair.Val2, 18)
}
