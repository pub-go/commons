package iters

import (
	"iter"

	"code.gopub.tech/commons/funcs"
	"code.gopub.tech/commons/nums"
)

// OfSeq convert from iter.Seq.
// 类型转换, 从 iter.Seq 转为 Seq.
func OfSeq[T any](s iter.Seq[T]) Seq[T] {
	return Seq[T](s)
}

// Repeat create an infinite Seq,
// which all elements is same as the input element.
// 使用输入的单个元素创建一个无限序列
func Repeat[T any](e T) Seq[T] {
	return Generate(func() T { return e })
}

// Generate build a Seq,
// which each element is generate by the Supplier.
// 通过生成器生成一个序列
func Generate[T any](get funcs.Supplier[T]) Seq[T] {
	return func(yield func(T) bool) {
		for {
			if !yield(get()) {
				return
			}
		}
	}
}

// Of build a Seq by input elements.
// 使用输入的任意个元素创建一个序列
func Of[T any](elements ...T) Seq[T] {
	return OfSlice(elements)
}

// OfSlice build a Seq by input slice.
// 使用输入的切片创建一个序列
func OfSlice[T any](s []T) Seq[T] {
	return func(yield func(T) bool) {
		for _, element := range s {
			if !yield(element) {
				return
			}
		}
	}
}

// Range build a Seq count from fromInclude to toExclude.
// 构造一个左闭右开的升序区间序列
func Range[T nums.Number](fromInclude, toExclude T) Seq[T] {
	if fromInclude <= toExclude {
		return RangeStep(fromInclude, toExclude, 1)
	}
	return RangeStep(fromInclude, toExclude, -1)
}

// RangeStep build a Seq count from fromInclude to toExclude.
// step may negetive.
// 按指定步进大小构造一个左闭右开的序列, 步进大小可以是负数
func RangeStep[T, S nums.Number](fromInclude, toExclude T, step S) Seq[T] {
	return func(yield func(T) bool) {
		if step >= 0 {
			for i := fromInclude; i < toExclude; i += T(step) {
				if !yield(i) {
					return
				}
			}
		} else {
			// 需要处理 T=uint step=-1 的情况:
			// -step 转为正数; 然后依次递减
			for i := fromInclude; i > toExclude; i -= T(-step) {
				if !yield(i) {
					return
				}
			}
		}
	}
}
