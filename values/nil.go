package values

import "unsafe"

type xface struct {
	x    uintptr
	data unsafe.Pointer
}

// IsNil returns whether the given value v is nil.
//
// 💡 NOTE: Typed nil interface (such as fmt.Stringer((*net.IP)(nil))) is nil,
// although fmt.Stringer((*net.IP)(nil)) != nil.
//
// 🚀 EXAMPLE:
//
//	IsNil(nil)                           ⏩ true
//	IsNil(1)                             ⏩ false
//	IsNil((*int)(nil))                   ⏩ true
//	IsNil(fmt.Stringer((*net.IP)(nil)))  ⏩ true
//
// ⚠️ WARNING: This function is implemented using [unsafe].
func IsNil(v any) bool {
	return (*xface)(unsafe.Pointer(&v)).data == nil
}

func IsNotNil(v any) bool {
	return !IsNil(v)
}
