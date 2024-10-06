package iters

import (
	"iter"

	"code.gopub.tech/commons/funcs"
	"code.gopub.tech/commons/order"
)

type Stream[T any] interface {
	Seq() iter.Seq[T]

	Filter(test funcs.Predicate[T]) Stream[T]
	Map(apply funcs.Function[T, T]) Stream[T]
	FlatMap(flatten funcs.Function[T, iter.Seq[T]]) Stream[T]
	Peek(accept funcs.Consumer[T]) Stream[T]

	Distinct(getUniqKey funcs.Function[T, int]) Stream[T]
	Sorted(cmp order.Comparator[T]) Stream[T]
	Limit(n int64) Stream[T]
	Skip(n int64) Stream[T]

	Count() int64
	ToSlice() []T
	ForEach(accept funcs.Consumer[T])
	AllMatch(test funcs.Predicate[T]) bool
	AnyMatch(test funcs.Predicate[T]) bool
	NoneMatch(test funcs.Predicate[T]) bool
	Reduce(initVal T, acc funcs.BinaryOperator[T]) T
}

// Maps is alias of Map.
// 同 Map, 不过参数是 Stream 而不是 iter.Seq.
func Maps[T, R any](stream Stream[T], apply funcs.Function[T, R]) Stream[R] {
	return Seq[R](Map(iter.Seq[T](stream.Seq()), apply))
}

// FlatMaps is alias of FlatMap.
// 同 FlatMap, 不过参数是 Stream 而不是 iter.Seq.
func FlatMaps[T, R any](stream Stream[T], flatten funcs.Function[T, iter.Seq[R]]) Stream[R] {
	return Seq[R](FlatMap(iter.Seq[T](stream.Seq()), func(input T) iter.Seq[R] {
		return flatten(input)
	}))
}
