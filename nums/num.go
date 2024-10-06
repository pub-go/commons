package nums

import "reflect"

type (
	Signed interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}

	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}

	Int interface {
		Signed | Unsigned
	}

	Float interface {
		~float32 | ~float64
	}

	Complex interface {
		~complex64 | ~complex128
	}

	Number interface {
		Int | Float
	}
	AnyNumber any
)

func To[T Number](n AnyNumber) T {
	switch i := n.(type) {
	case int:
		return T(i)
	case int8:
		return T(i)
	case int16:
		return T(i)
	case int32:
		return T(i)
	case int64:
		return T(i)
	case uint:
		return T(i)
	case uint8:
		return T(i)
	case uint16:
		return T(i)
	case uint32:
		return T(i)
	case uint64:
		return T(i)
	case uintptr:
		return T(i)
	case float32:
		return T(i)
	case float64:
		return T(i)
	}
	rt := reflect.TypeOf(n)
	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return T(reflect.ValueOf(n).Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return T(reflect.ValueOf(n).Uint())
	case reflect.Float32, reflect.Float64:
		return T(reflect.ValueOf(n).Float())
	}
	return T(0)
}
