package iters_test

import (
	"iter"
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/commons/funcs"
	"code.gopub.tech/commons/iters"
	"code.gopub.tech/commons/nums"
	"code.gopub.tech/commons/order"
)

func TestSeq(t *testing.T) {
	iters.Range(0, 10).
		Peek(func(i int) {
			t.Logf("from source: %d", i)
		}).
		Filter(func(i int) bool { return i%2 == 0 }).
		Peek(func(i int) {
			t.Logf("after filter: %d", i)
		}).
		Map(func(i int) int { return 2 * i }).
		Peek(func(i int) {
			t.Logf("map to *2: %d", i)
		}).
		FlatMap(func(i int) iter.Seq[int] {
			return iters.Repeat(i).Limit(2).Seq()
		}).
		Peek(func(i int) {
			t.Logf("after flatten: %d", i)
		}).
		Distinct(funcs.Identidy).
		Sorted(order.Reversed[int]).
		Skip(1).
		ForEach(func(i int) {
			t.Logf("got: %v", i)
		})

	assert.DeepEqual(t,
		iters.Range(0, 10).ToSlice(),
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	)
	assert.Equal(t, iters.Range(0, 10).Count(), 10)
	assert.True(t, iters.Range(0, 10).
		AllMatch(func(i int) bool { return i >= 0 }))
	assert.True(t, iters.Range(0, 10).
		AnyMatch(func(i int) bool { return i >= 0 }))
	assert.True(t, iters.Range(0, 10).
		NoneMatch(func(i int) bool { return i < 0 }))
	assert.Equal(t, iters.Range(0, 10).Reduce(0, nums.Add), 45)

	for i, n := range iters.Range(0, 10).Seq2() {
		t.Logf("i=%v, n=%v", i, n)
		if i == 5 {
			break
		}
	}

}
