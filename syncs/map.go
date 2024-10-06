package syncs

import (
	"sync"
)

type Map[K comparable, V any] struct {
	m sync.Map
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{}
}

// Clear wraps [sync.Map.Clear].
// 清除 map 中的所有键值对.
func (sm *Map[K, V]) Clear() {
	sm.m.Clear()
}

// CompareAndDelete wraps [sync.Map.CompareAndDelete].
// 比较并删除, 如果 key 存在且旧值等于 old, 则删除该键值对.
// 值的类型必须是可比较的.
func (sm *Map[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return sm.m.CompareAndDelete(key, old)
}

// CompareAndSwap wraps [sync.Map.CompareAndSwap].
// 比较并交换, 如果 key 存在且旧值等于 old, 则将新值存储到 map 中.
// 值的类型必须是可比较的.
func (sm *Map[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	return sm.m.CompareAndSwap(key, old, new)
}

// Delete wraps [sync.Map.Delete].
// 删除一个键值对.
func (sm *Map[K, V]) Delete(key K) {
	sm.m.LoadAndDelete(key)
}

// Load wraps [sync.Map.Load].
// 返回 map 中的值, 如果没有则返回零值.
// ok 字段表示 key 是否存在.
func (sm *Map[K, V]) Load(k K) (value V, ok bool) {
	v, ok := sm.m.Load(k)
	if ok {
		return v.(V), true
	}
	return
}

// LoadAndDelete wraps [sync.Map.LoadAndDelete].
// 删除一个键值对, 并返回旧值(如有).
// loaded 字段表示 key 是否存在.
func (sm *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	val, ok := sm.m.LoadAndDelete(key)
	if ok {
		return val.(V), ok
	}
	return
}

// LoadOrStore wraps [sync.Map.LoadOrStore].
// 如果 key 存在, 则返回旧值; 否则存储一个键值对到 map 中, 并返回新值.
// loaded 字段为 true 表示 key 已经存在, false 表示返回的是存储的新值.
func (sm *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	val, loaded := sm.m.LoadOrStore(key, value)
	return val.(V), loaded
}

// Range wraps [sync.Map.Range].
// 遍历 map 中的所有键值对.
// 对 map 中的每个键值对依次调用 f, 如果 f 返回 false, 则停止迭代.
func (sm *Map[K, V]) Range(f func(key K, value V) bool) {
	sm.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Store wraps [sync.Map.Store].
// 存储一个键值对到 map 中.
func (sm *Map[K, V]) Store(key K, value V) {
	sm.m.Swap(key, value)
}

// Swap wraps [sync.Map.Swap].
// 存储一个键值对到 map 中, 如果 key 已经存在, 则返回旧值.
// loaded 字段表示 key 是否存在.
func (sm *Map[K, V]) Swap(key K, value V) (actual V, loaded bool) {
	val, ok := sm.m.Swap(key, value)
	if ok {
		return val.(V), ok
	}
	return
}

func (sm *Map[K, V]) ToMap() map[K]V {
	mm := make(map[K]V)
	sm.Range(func(k K, v V) bool {
		mm[k] = v
		return true
	})
	return mm
}
