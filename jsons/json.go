package jsons

import (
	"encoding/json"

	"code.gopub.tech/commons/conv"
)

// ToJSON 将参数转为 JSON, 如果出现错误则返回空字符串
func ToJSON(obj any, opts ...Opt) string {
	b := ToBytes(obj, opts...)
	return conv.Bytes2String(b)
}

// ToBytes return []byte. return nil if error
func ToBytes(obj any, opts ...Opt) []byte {
	b, err := ToBytesE(obj, opts...)
	if err != nil {
		return nil
	}
	return b
}

func ToBytesE(data any, opts ...Opt) ([]byte, error) {
	option := getOption(opts...)
	if option.prefix == "" && option.indent == "" {
		return option.marshal(data)
	}
	return option.marshalIndent(data, option.prefix, option.indent)
}

// Indent 使用两个空格缩进
func Indent(obj any, opts ...Opt) string {
	opts = append([]Opt{UseIndent("", "  ")}, opts...)
	return ToJSON(obj, opts...)
}

type opt struct {
	marshal        func(v any) ([]byte, error)
	prefix, indent string
	marshalIndent  func(v any, prefix string, indent string) ([]byte, error)
}

type Opt func(*opt)

func getOption(opts ...Opt) *opt {
	option := &opt{
		marshal:       json.Marshal,
		marshalIndent: json.MarshalIndent,
	}
	for _, setter := range opts {
		setter(option)
	}
	return option
}

func UseMarshal(marshal func(v any) ([]byte, error)) Opt {
	return func(o *opt) { o.marshal = marshal }
}

func UseIndent(prefix, indent string) Opt {
	return func(o *opt) { o.prefix, o.indent = prefix, indent }
}

func UseMarshalIndent(marshalIndent func(v any, prefix string, indent string) ([]byte, error)) Opt {
	return func(o *opt) { o.marshalIndent = marshalIndent }
}
