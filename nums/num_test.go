package nums_test

import (
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/commons/nums"
)

type (
	MyInt  int
	MyUint uint
	Double float64
)

func TestTo(t *testing.T) {
	assert.Equal(t, nums.To[int](int(1)), int(1))
	assert.Equal(t, nums.To[int](int8(1)), int(1))
	assert.Equal(t, nums.To[int](int16(1)), int(1))
	assert.Equal(t, nums.To[int](int32(1)), int(1))
	assert.Equal(t, nums.To[int](int64(1)), int(1))
	assert.Equal(t, nums.To[int](uint(1)), int(1))
	assert.Equal(t, nums.To[int](uint8(1)), int(1))
	assert.Equal(t, nums.To[int](uint16(1)), int(1))
	assert.Equal(t, nums.To[int](uint32(1)), int(1))
	assert.Equal(t, nums.To[int](uint64(1)), int(1))
	assert.Equal(t, nums.To[int](uintptr(1)), int(1))
	assert.Equal(t, nums.To[int](float32(1)), int(1))
	assert.Equal(t, nums.To[int](float64(1)), int(1))
	assert.Equal(t, nums.To[int](MyInt(1)), int(1))
	assert.Equal(t, nums.To[int](MyUint(1)), int(1))
	assert.Equal(t, nums.To[int](Double(1)), int(1))
	assert.Equal(t, nums.To[int]("1"), int(0))
}
