package choose

import "code.gopub.tech/commons/funcs"

func If[T any](cond bool, onTrue, onFalse T) T {
	if cond {
		return onTrue
	}
	return onFalse
}

func IfLazy[T any](cond bool, onTrue, onFalse funcs.Supplier[T]) T {
	if cond {
		return onTrue()
	}
	return onFalse()
}

func IfLazyT[T any](cond bool, onTrue funcs.Supplier[T], onFalse T) T {
	if cond {
		return onTrue()
	}
	return onFalse
}

func IfLazyF[T any](cond bool, onTrue T, onFalse funcs.Supplier[T]) T {
	if cond {
		return onTrue
	}
	return onFalse()
}
