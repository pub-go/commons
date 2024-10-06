package funcs

type Func1[T1, R any] func(T1) R

func Of1[T1, R any, F ~func(T1) R](fn F) Func1[T1, R] {
	return Func1[T1, R](fn)
}

func (fn Func1[T1, R]) Call(t1 T1) R {
	return fn(t1)
}

func (fn Func1[T1, R]) Partial(t1 T1) Func0[R] {
	return Partial1(fn, t1)
}
func (fn Func1[T1, R]) Partial1(t1 T1) Func0[R] {
	return Partial1N1(fn, t1)
}

func (fn Func1[T1, R]) PartialR(t1 T1) Func0[R] {
	return Partial1R(fn, t1)
}

func Partial1[T1, R any, F ~func(T1) R](fn F, t1 T1) Func0[R] {
	return func() R {
		return fn(t1)
	}
}

func Partial1N1[T1, R any, F ~func(T1) R](fn F, t1 T1) Func0[R] {
	return func() R {
		return fn(t1)
	}
}

func Partial1R[T1, R any, F ~func(T1) R](fn F, t1 T1) Func0[R] {
	return func() R {
		return fn(t1)
	}
}
