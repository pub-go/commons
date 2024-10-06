package values

// Zero 返回零值
func Zero[T any]() (zero T) {
	return
}

// IsZero 判断是否是零值
func IsZero[T comparable](a T) bool {
	return a == Zero[T]()
}

// IsNotZero 判断是否不是零值
func IsNotZero[T comparable](a T) bool {
	return a != Zero[T]()
}
