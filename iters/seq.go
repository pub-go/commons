package iters

import (
	"iter"

	"code.gopub.tech/commons/funcs"
	"code.gopub.tech/commons/order"
)

type Seq[T any] iter.Seq[T]

func (it Seq[T]) Seq() iter.Seq[T] {
	return iter.Seq[T](it)
}

func (it Seq[T]) Seq2() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var i int
		for e := range it {
			if !yield(i, e) {
				return
			}
			i++
		}
	}
}

func (it Seq[T]) Filter(test funcs.Predicate[T]) Stream[T] {
	return Seq[T](Filter(iter.Seq[T](it), test))
}

func (it Seq[T]) Map(apply funcs.Function[T, T]) Stream[T] {
	// return Maps(it, apply) // 多几层函数调用 最终也是下一行的形式
	return Seq[T](Map(iter.Seq[T](it), apply))
}

func (it Seq[T]) FlatMap(flatten funcs.Function[T, iter.Seq[T]]) Stream[T] {
	// return FlatMaps(it, flatten)
	return Seq[T](FlatMap(iter.Seq[T](it), flatten))
}

func (it Seq[T]) Peek(accept funcs.Consumer[T]) Stream[T] {
	return Seq[T](Peek(iter.Seq[T](it), accept))
}

func (it Seq[T]) Distinct(getUniqKey funcs.Function[T, int]) Stream[T] {
	return Seq[T](Distinct(iter.Seq[T](it), getUniqKey))
}

func (it Seq[T]) Sorted(cmp order.Comparator[T]) Stream[T] {
	return Seq[T](Sorted(iter.Seq[T](it), cmp))
}

func (it Seq[T]) Limit(n int64) Stream[T] {
	return Seq[T](Limit(iter.Seq[T](it), n))
}

func (it Seq[T]) Skip(n int64) Stream[T] {
	return Seq[T](Skip(iter.Seq[T](it), n))
}

func (it Seq[T]) ForEach(accept funcs.Consumer[T]) {
	ForEach(iter.Seq[T](it), accept)
}

func (it Seq[T]) ToSlice() []T {
	return ToSlice(iter.Seq[T](it))
}

func (it Seq[T]) Count() int64 {
	return Count(iter.Seq[T](it))
}

func (it Seq[T]) AllMatch(test funcs.Predicate[T]) bool {
	return AllMatch(iter.Seq[T](it), test)
}

func (it Seq[T]) AnyMatch(test funcs.Predicate[T]) bool {
	return AnyMatch(iter.Seq[T](it), test)
}

func (it Seq[T]) NoneMatch(test funcs.Predicate[T]) bool {
	return NoneMatch(iter.Seq[T](it), test)
}

func (it Seq[T]) Reduce(initVal T, acc funcs.BinaryOperator[T]) T {
	return Reduce(iter.Seq[T](it), initVal, funcs.BiFunction[T, T, T](acc))
}
