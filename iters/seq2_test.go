package iters_test

import (
	"strconv"
	"testing"

	"code.gopub.tech/commons/iters"
	"code.gopub.tech/commons/order"
)

func TestSeq2(t *testing.T) {
	m := iters.Reduce(
		iters.OfSlice2([]int{1, 2, 3}).Entry().Seq(),
		map[string]int{},
		func(sum map[string]int, e iters.Entry[int, int]) map[string]int {
			sum[strconv.Itoa(e.Key)] = e.Val
			return sum
		})
	iters.OfMap(m).Sorted(func(e1, e2 iters.Entry[string, int]) int {
		return order.Reversed(e1.Key, e2.Key)
	}).Entry().ForEach(func(e iters.Entry[string, int]) {
		t.Logf("entry=%v", e)
	})
}
