package funcs_test

import (
	"strconv"
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/commons/funcs"
)

func add(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10 int) int {
	return t1 + t2 + t3 + t4 + t5 + t6 + t7 + t8 + t9 + t10
}

func TestPartial(t *testing.T) {
	fn := funcs.Of10(add)
	t.Run("of10", func(t *testing.T) {
		assert.Equal(t,
			fn.Partial1(1)(2, 3, 4, 5, 6, 7, 8, 9, 10),
			fn.PartialR(10)(1, 2, 3, 4, 5, 6, 7, 8, 9),
		)
		assert.Equal(t,
			fn.Partial2(2)(1, 3, 4, 5, 6, 7, 8, 9, 10),
			fn.Partial3(3)(1, 2, 4, 5, 6, 7, 8, 9, 10),
		)
		assert.Equal(t,
			fn.Partial4(4)(1, 2, 3, 5, 6, 7, 8, 9, 10),
			fn.Partial5(5)(1, 2, 3, 4, 6, 7, 8, 9, 10),
		)
		assert.Equal(t,
			fn.Partial6(6)(1, 2, 3, 4, 5, 7, 8, 9, 10),
			fn.Partial7(7)(1, 2, 3, 4, 5, 6, 8, 9, 10),
		)
		assert.Equal(t,
			fn.Partial8(8)(1, 2, 3, 4, 5, 6, 7, 9, 10),
			fn.Partial9(9)(1, 2, 3, 4, 5, 6, 7, 8, 10),
		)
		assert.Equal(t,
			fn.Partial1(1)(2, 3, 4, 5, 6, 7, 8, 9, 10),
			fn.Partial10(10)(1, 2, 3, 4, 5, 6, 7, 8, 9),
		)
		assert.Equal(t,
			fn.Call(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			55,
		)
	})
	t.Run("of9", func(t *testing.T) {
		f := fn.PartialR(10)
		assert.Equal(t, funcs.Of9(f).Partial(1)(2, 3, 4, 5, 6, 7, 8, 9),
			f.PartialR(9)(1, 2, 3, 4, 5, 6, 7, 8))
		assert.Equal(t, f.Partial1(1)(2, 3, 4, 5, 6, 7, 8, 9), f.Partial9(9)(1, 2, 3, 4, 5, 6, 7, 8))
		assert.Equal(t, f.Partial2(2)(1, 3, 4, 5, 6, 7, 8, 9), f.Partial8(8)(1, 2, 3, 4, 5, 6, 7, 9))
		assert.Equal(t, f.Partial3(3)(1, 2, 4, 5, 6, 7, 8, 9), f.Partial7(7)(1, 2, 3, 4, 5, 6, 8, 9))
		assert.Equal(t, f.Partial4(4)(1, 2, 3, 5, 6, 7, 8, 9), f.Partial6(6)(1, 2, 3, 4, 5, 7, 8, 9))
		assert.Equal(t, f.Partial5(5)(1, 2, 3, 4, 6, 7, 8, 9), 55)
		assert.Equal(t, f.Call(1, 2, 3, 4, 5, 6, 7, 8, 9), 55)
	})
	t.Run("of8", func(t *testing.T) {
		f := fn.PartialR(10).PartialR(9)
		assert.Equal(t, funcs.Of8(f).Partial(1)(2, 3, 4, 5, 6, 7, 8),
			f.PartialR(8)(1, 2, 3, 4, 5, 6, 7))
		assert.Equal(t, f.Partial1(1)(2, 3, 4, 5, 6, 7, 8), f.Partial8(8)(1, 2, 3, 4, 5, 6, 7))
		assert.Equal(t, f.Partial2(2)(1, 3, 4, 5, 6, 7, 8), f.Partial7(7)(1, 2, 3, 4, 5, 6, 8))
		assert.Equal(t, f.Partial3(3)(1, 2, 4, 5, 6, 7, 8), f.Partial6(6)(1, 2, 3, 4, 5, 7, 8))
		assert.Equal(t, f.Partial4(4)(1, 2, 3, 5, 6, 7, 8), f.Partial5(5)(1, 2, 3, 4, 6, 7, 8))
		assert.Equal(t, f.Call(1, 2, 3, 4, 5, 6, 7, 8), 55)
	})
	t.Run("of7", func(t *testing.T) {
		f := funcs.Of7(fn.PartialR(10).PartialR(9).Partial8(8))
		assert.Equal(t, f.Partial(1)(2, 3, 4, 5, 6, 7), f.PartialR(7)(1, 2, 3, 4, 5, 6))
		assert.Equal(t, f.Partial1(1)(2, 3, 4, 5, 6, 7), f.Partial7(7)(1, 2, 3, 4, 5, 6))
		assert.Equal(t, f.Partial2(2)(1, 3, 4, 5, 6, 7), f.Partial6(6)(1, 2, 3, 4, 5, 7))
		assert.Equal(t, f.Partial3(3)(1, 2, 4, 5, 6, 7), f.Partial5(5)(1, 2, 3, 4, 6, 7))
		assert.Equal(t, f.Partial4(4)(1, 2, 3, 5, 6, 7), 55)
		assert.Equal(t, f.Call(1, 2, 3, 4, 5, 6, 7), 55)
	})
	t.Run("of6", func(t *testing.T) {
		f := funcs.Of6(fn.PartialR(10).PartialR(9).Partial8(8).Partial7(7))
		assert.Equal(t, f.Partial(1)(2, 3, 4, 5, 6), f.PartialR(6)(1, 2, 3, 4, 5))
		assert.Equal(t, f.Partial1(1)(2, 3, 4, 5, 6), f.Partial6(6)(1, 2, 3, 4, 5))
		assert.Equal(t, f.Partial2(2)(1, 3, 4, 5, 6), f.Partial5(5)(1, 2, 3, 4, 6))
		assert.Equal(t, f.Partial3(3)(1, 2, 4, 5, 6), f.Partial4(4)(1, 2, 3, 5, 6))
		assert.Equal(t, f.Call(1, 2, 3, 4, 5, 6), 55)
	})
	t.Run("of5", func(t *testing.T) {
		f := funcs.Of5(fn.PartialR(10).PartialR(9).Partial8(8).Partial7(7).PartialR(6))
		assert.Equal(t, f.Partial(1)(2, 3, 4, 5), f.PartialR(5)(1, 2, 3, 4))
		assert.Equal(t, f.Partial1(1)(2, 3, 4, 5), f.Partial5(5)(1, 2, 3, 4))
		assert.Equal(t, f.Partial2(2)(1, 3, 4, 5), f.Partial4(4)(1, 2, 3, 5))
		assert.Equal(t, f.Partial3(3)(1, 2, 4, 5), 55)
		assert.Equal(t, f.Call(1, 2, 3, 4, 5), 55)
	})
	t.Run("of4", func(t *testing.T) {
		f := funcs.Of4(fn.PartialR(10).PartialR(9).PartialR(8).PartialR(7).PartialR(6).PartialR(5))
		assert.Equal(t, f.Partial(1)(2, 3, 4), f.PartialR(4)(1, 2, 3))
		assert.Equal(t, f.Partial1(1)(2, 3, 4), f.Partial4(4)(1, 2, 3))
		assert.Equal(t, f.Partial2(2)(1, 3, 4), f.Partial3(3)(1, 2, 4))
		assert.Equal(t, f.Call(1, 2, 3, 4), 55)
	})
	t.Run("of3", func(t *testing.T) {
		f := funcs.Of3(fn.PartialR(10).PartialR(9).PartialR(8).PartialR(7).
			PartialR(6).PartialR(5).PartialR(4))
		assert.Equal(t, f.Partial(1)(2, 3), f.PartialR(3)(1, 2))
		assert.Equal(t, f.Partial1(1)(2, 3), f.Partial3(3)(1, 2))
		assert.Equal(t, f.Partial2(2)(1, 3), 55)
		assert.Equal(t, f.Call(1, 2, 3), 55)
	})
	t.Run("of2", func(t *testing.T) {
		f := funcs.Of2(fn.PartialR(10).PartialR(9).PartialR(8).PartialR(7).
			PartialR(6).PartialR(5).PartialR(4).PartialR(3))
		assert.Equal(t, f.Partial(1)(2), f.PartialR(2)(1))
		assert.Equal(t, f.Partial1(1)(2), f.Partial2(2)(1))
		assert.Equal(t, f.Call(1, 2), 55)
	})
	t.Run("of1", func(t *testing.T) {
		f := funcs.Of1(fn.PartialR(10).PartialR(9).PartialR(8).PartialR(7).
			PartialR(6).PartialR(5).PartialR(4).PartialR(3).PartialR(2))
		assert.Equal(t, f.Partial(1)(), f.PartialR(1)())
		assert.Equal(t, f.Partial1(1)(), 55)
		assert.Equal(t, f.Call(1), 55)
	})
	t.Run("of0", func(t *testing.T) {
		f := funcs.Of0(fn.PartialR(10).PartialR(9).PartialR(8).PartialR(7).
			PartialR(6).PartialR(5).PartialR(4).PartialR(3).PartialR(2).PartialR(1))
		assert.Equal(t, f(), 55)
		assert.Equal(t, f(), f.Call())
	})
	assert.Equal(t, 55,
		fn.
			Partial(1).
			Partial(2).
			Partial(3).
			Partial(4).
			Partial(5).
			Partial(6).
			Partial(7).
			Partial(8).
			Partial(9).
			Partial(10).
			Call(),
	)
	i2a := funcs.Partial2R(strconv.FormatInt, 10)
	assert.Equal(t, i2a(10), strconv.Itoa(10))

	format10to := funcs.Partial2(strconv.FormatInt, 10)
	assert.Equal(t, format10to(10), "10")
	assert.Equal(t, format10to(2), "1010")
}
