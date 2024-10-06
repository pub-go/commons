package iters

import (
	"iter"

	"code.gopub.tech/commons/order"
)

type Seq2[K, V any] iter.Seq2[K, V]

type Entry[K, V any] struct {
	Key K
	Val V
}

func OfSlice2[T any](s []T) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, element := range s {
			if !yield(i, element) {
				return
			}
		}
	}
}

func OfMap[K comparable, V any](m map[K]V) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (it Seq2[K, V]) Sorted(cmp order.Comparator[Entry[K, V]]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for entry := range it.Entry().Sorted(cmp).Seq() {
			if !yield(entry.Key, entry.Val) {
				return
			}
		}

	}
}

func (it Seq2[K, V]) Entry() Seq[Entry[K, V]] {
	return func(yield func(Entry[K, V]) bool) {
		for k, v := range it {
			if !yield(Entry[K, V]{Key: k, Val: v}) {
				return
			}
		}
	}
}
