package funcs

type (
	// Supplier 产生一个元素
	Supplier[T any] func() T

	// Consumer 消费一个元素
	Consumer[T any] func(T)

	// Function 将一个类型转为另一个类型
	Function[T, R any] func(T) R

	// Predicate 断言是否满足指定条件
	Predicate[T any] Function[T, bool]

	// UnaryOperator 对输入进行一元运算返回相同类型的结果
	UnaryOperator[T any] Function[T, T]

	// BiFunction 将两个类型转为第三个类型
	BiFunction[X, Y, R any] func(X, Y) R

	// BinaryOperator 输入两个相同类型的参数，对其做二元运算，返回相同类型的结果
	BinaryOperator[T any] BiFunction[T, T, T]
)

// Identidy 一个返回自身的函数
func Identidy[T any](e T) T { return e }

// Not 断言取反
func Not[T any](test Predicate[T]) Predicate[T] {
	return func(t T) bool { return !test(t) }
}
