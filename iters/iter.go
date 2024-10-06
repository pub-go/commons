package iters

import (
	"iter"
	"sort"

	"code.gopub.tech/commons/funcs"
	"code.gopub.tech/commons/nums"
	"code.gopub.tech/commons/order"
)

// Filter keep elements which satisfy the Predicate.
// 保留满足断言的元素
func Filter[T any](seq iter.Seq[T], test funcs.Predicate[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			val, ok := next()
			if !ok {
				return
			}
			if test(val) && !yield(val) {
				return
			}
		}
	}
}

// Map transform the element use Fuction.
// 使用输入函数对每个元素进行转换
func Map[T, R any](seq iter.Seq[T], apply funcs.Function[T, R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			val, ok := next()
			if !ok || !yield(apply(val)) {
				return
			}
		}
	}
}

// FlatMap transform each element in Seq[T] to a new Seq[R].
// 将原本序列中的每个元素都转换为一个新的序列，
// 并将所有转换后的序列依次连接起来生成一个新的序列
func FlatMap[T, R any](seq iter.Seq[T], flatten funcs.Function[T, iter.Seq[R]]) iter.Seq[R] {
	return func(yield func(R) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			if !ok {
				return
			}
			for v := range flatten(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Peek visit every element in the Seq and leave them on the Seq.
// 访问序列中的每个元素而不消费它
func Peek[T any](seq iter.Seq[T], accept funcs.Consumer[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			if !ok {
				return
			}
			accept(v)
			if !yield(v) {
				return
			}
		}
	}
}

// Distinct remove duplicate elements.
// 对序列中的元素去重
func Distinct[T any, Cmp comparable](seq iter.Seq[T], getUniqKey funcs.Function[T, Cmp]) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		var set = make(map[Cmp]struct{})
		for {
			v, ok := next()
			if !ok {
				return
			}
			k := getUniqKey(v)
			_, ok = set[k]
			set[k] = struct{}{}
			if !ok && !yield(v) {
				return
			}
		}
	}
}

// Sorted sort elements in the Seq by Comparator.
// 对序列中的元素排序
func Sorted[T any](seq iter.Seq[T], cmp order.Comparator[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		vals := ToSlice(seq)
		sort.SliceStable(vals, func(i, j int) bool {
			return cmp(vals[i], vals[j]) < 0
		})
		for _, v := range vals {
			if !yield(v) {
				return
			}
		}
	}
}

// Limit limits the number of elements in Seq.
// 限制元素个数
func Limit[T any, Number nums.Int](seq iter.Seq[T], limit Number) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			limit--
			if limit < 0 {
				return
			}
			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}

// Skip drop some elements of the Seq.
// 跳过指定个数的元素
func Skip[T any, Number nums.Int](seq iter.Seq[T], skip Number) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			skip--
			if skip < 0 {
				if !ok || !yield(v) {
					return
				}
			}
		}
	}
}

// ForEach consume every elements in the Seq.
// 消费序列中的每个元素
func ForEach[T any](seq iter.Seq[T], accept funcs.Consumer[T]) {
	for v := range seq {
		accept(v)
	}
}

// ToSlice return all elements as a slice.
// 将序列中所有元素收集为切片返回
func ToSlice[T any](seq iter.Seq[T]) (result []T) {
	for v := range seq {
		result = append(result, v)
	}
	return
}

// Count return the count of elements in the Seq.
// 返回序列中的元素个数
func Count[T any](seq iter.Seq[T]) (count int64) {
	for _ = range seq {
		count++
	}
	return
}

// AllMatch test if every elements are all match the Predicate.
// 是否每个元素都满足条件
func AllMatch[T any](seq iter.Seq[T], test funcs.Predicate[T]) bool {
	for v := range seq {
		if !test(v) {
			return false
		}
	}
	return true
}

// AnyMatch test if any element matches the Predicate.
// 是否有任意元素满足条件
func AnyMatch[T any](seq iter.Seq[T], test funcs.Predicate[T]) bool {
	for v := range seq {
		if test(v) {
			return true
		}
	}
	return false
}

// NoneMatch test if none element matches the Predicate.
// 是否没有元素满足条件
func NoneMatch[T any](seq iter.Seq[T], test funcs.Predicate[T]) bool {
	for v := range seq {
		if test(v) {
			return false
		}
	}
	return true
}

// Reduce accumulate each element using the BiFunction
// starting from the initial value.
// 从初始值开始, 通过 acc 函数累加每个元素
func Reduce[T, R any](seq iter.Seq[T], initVal R, acc funcs.BiFunction[R, T, R]) (result R) {
	result = initVal
	for v := range seq {
		result = acc(result, v)
	}
	return
}
