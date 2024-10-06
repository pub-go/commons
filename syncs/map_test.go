package syncs_test

import (
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/commons/syncs"
)

func TestMap(t *testing.T) {
	m := syncs.NewMap[string, string]()
	v, loaded := m.Load("a") // {}
	assert.Equal(t, v, "")
	assert.Equal(t, loaded, false)

	v, loaded = m.LoadOrStore("a", "b") // {"a": "b"}
	assert.Equal(t, v, "b")
	assert.Equal(t, loaded, false)
	v, loaded = m.LoadOrStore("a", "b1") // {"a": "b"}
	assert.Equal(t, v, "b")
	assert.Equal(t, loaded, true)

	v, loaded = m.LoadAndDelete("a") // {}
	assert.Equal(t, v, "b")
	assert.Equal(t, loaded, true)
	v, loaded = m.LoadAndDelete("a") // {}
	assert.Equal(t, v, "")
	assert.Equal(t, loaded, false)

	m.Store("k", "v") // {"k": "v"}
	v, loaded = m.Load("k")
	assert.Equal(t, v, "v")
	assert.Equal(t, loaded, true)
	v, loaded = m.Swap("k", "v1") // {"k": "v"} -> {"k": "v1"}
	assert.Equal(t, v, "v")
	assert.Equal(t, loaded, true)

	v, loaded = m.Swap("x", "y") //  {"k": "v1", "x", "y"}
	assert.Equal(t, v, "")
	assert.Equal(t, loaded, false)
	v, loaded = m.Load("x")
	assert.Equal(t, v, "y")
	assert.Equal(t, loaded, true)

	swaped := m.CompareAndSwap("x", "y", "y1") //  {"k": "v1", "x", "y"} -> {"k": "v1", "x", "y1"}
	assert.True(t, swaped)
	v, loaded = m.Load("x")
	assert.Equal(t, v, "y1")
	assert.Equal(t, loaded, true)

	swaped = m.CompareAndSwap("x", "y", "y2") // {"k": "v1", "x", "y1"}
	assert.False(t, swaped)
	v, loaded = m.Load("x")
	assert.Equal(t, v, "y1")
	assert.Equal(t, loaded, true)

	deleted := m.CompareAndDelete("x", "y1") // {"k": "v1", "x", "y1"} -> {"k": "v1"}
	assert.True(t, deleted)
	v, loaded = m.Load("x")
	assert.Equal(t, v, "")
	assert.Equal(t, loaded, false)
	m.Delete("x") 

	assert.DeepEqual(t, m.ToMap(), map[string]string{"k": "v1"})
	m.Clear()
	assert.DeepEqual(t, m.ToMap(), map[string]string{})
}
