package iters_test

import (
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/commons/iters"
)

func TestSource(t *testing.T) {
	ss := iters.OfSlice([]string{"a", "b", "c"}).
		Map(func(s string) string { return s + s }).
		ToSlice()
	assert.DeepEqual(t, ss, []string{"aa", "bb", "cc"})

	sum := iters.OfSeq(func(yield func(int64) bool) {
		for i := int64(0); i <= 100; i++ {
			if !yield(i) {
				return
			}
		}
	}).Reduce(0, func(i1, i2 int64) int64 { return i1 + i2 })
	assert.Equal(t, sum, 5050)

	for i, n := range iters.Repeat(1).Seq2() {
		assert.Equal(t, n, 1)
		if i > 10 {
			break
		}
	}

	for i := range iters.Of(10, 11, 12) {
		assert.True(t, i >= 10)
		if i == 11 {
			break
		}
		assert.True(t, i <= 11)
	}

	for i := range iters.Range(10, 20) {
		assert.True(t, i >= 10)
		if i == 15 {
			break
		}
		assert.True(t, i < 15)
	}
	for i := range iters.RangeStep(20, 10, -1) {
		assert.True(t, i >= 15)
		if i == 15 {
			break
		}
		assert.True(t, i <= 20)
	}
	var a, b = uint8(10), uint8(0)
	for i := range iters.Range(a, b) {
		t.Logf("range neg uint: %d(%T)", i, i)
	}
}
