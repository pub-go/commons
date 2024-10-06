package conv

import "unsafe"

func Bytes2String(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	// see strings.Builder.String
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// 如果 s 是在常量区，不能修改返回的字符数组，否则会炸
//
//	unexpected fault address 0x1144064
//	fatal error: fault
func String2ReadOnlyBytes(s string) []byte {
	if len(s) == 0 {
		// return nil    // 这里不返回 nil
		return []byte(s) // 直接转换为 []byte 返回
	}
	// see os.File.WriteString
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
