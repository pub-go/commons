package order

import (
	"code.gopub.tech/commons/funcs"
	"code.gopub.tech/commons/nums"
)

type Ordered interface {
	nums.Number | ~string
}

// Comparator 比较两个元素.
// 第一个元素大于第二个元素时，返回正数;
// 第一个元素小于第二个元素时，返回负数;
// 否则返回 0.
type Comparator[T any] funcs.BiFunction[T, T, int]

func Natural[T Ordered](t1, t2 T) int {
	switch {
	case t1 < t2:
		return -1
	case t1 > t2:
		return 1
	}
	return 0
}

func Reversed[T Ordered](t1, t2 T) int {
	return Reverse(Natural[T])(t1, t2)
}

func Reverse[T Ordered](cmp Comparator[T]) Comparator[T] {
	return func(t1, t2 T) int {
		return cmp(t2, t1)
	}
}
