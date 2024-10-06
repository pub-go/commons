package assert_test

import (
	"testing"

	"code.gopub.tech/commons/arg"
	"code.gopub.tech/commons/assert"
)

func add(a, b int) int { return a + b }

func foo() any { return nil }

func bar() error { return baz() }

type myerror string

func (e *myerror) Error() string { return string(*e) }

func baz() *myerror { return nil }

func TestAssert(t *testing.T) {
	assert.Equal(t, add(1, 1), 2)
	var s1 = []string{"a", "b"}
	var s2 = []string{"a", "b"}
	// assert.Equal(t, s1, s2) // []string does not satisfy comparable
	assert.DeepEqual(t, s1, s2)
	assert.True(t, add(1, 1) == 2)
	assert.False(t, add(1, 1) == 10)
	assert.Nil(t, foo())
	var err = bar()
	assert.Nil(t, err)
	assert.NotNil(t, 1)
	assert.ShouldNotPanic(t, func() {})
	assert.ShouldPanic(t, func() { panic("blabla") })
	t.Run("skip", func(t *testing.T) {
		 t.Skip()
		got := any(0)
		want := any(int64(0))
		assert.Equal(t, got, want, "got=%v, want=%v", arg.JSON(got), arg.JSON(want))
		assert.True(t, add(1, 1) == 10)
		assert.True(t, err == nil, "bar()==%v(%T)", err, err)
		assert.False(t, add(1, 1) == 2)
		assert.Nil(t, 1)
		assert.NotNil(t, err)
		assert.DeepEqual(t, s1, 1)
		assert.ShouldNotPanic(t, func() { panic("blabla") })
		assert.ShouldNotPanic(t, func() { panic("blabla") }, "panic=%v")
		assert.ShouldPanic(t, func() {})
	})

}
