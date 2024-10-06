package iters_test

import (
	"iter"
	"strconv"
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/commons/iters"
)

func TestIters(t *testing.T) {
	assert.DeepEqual(t,
		iters.Maps(iters.Of(1, 2, 3), strconv.Itoa).ToSlice(),
		[]string{"1", "2", "3"},
	)

	assert.DeepEqual(t,
		iters.FlatMaps(iters.Of(1, 2, 3), func(i int) iter.Seq[string] {
			s := strconv.Itoa(i)
			return iters.Repeat(s).Limit(2).Seq()
		}).ToSlice(),
		[]string{"1", "1", "2", "2", "3", "3"},
	)

}
