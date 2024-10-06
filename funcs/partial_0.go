package funcs

type Func0[R any] func() R

func Of0[R any, F ~func() R](fn F) Func0[R] {
	return Func0[R](fn)
}

func (fn Func0[R]) Call() R {
	return fn()
}
