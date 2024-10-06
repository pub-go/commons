package arg

import (
	"fmt"

	"code.gopub.tech/commons/conv"
	"code.gopub.tech/commons/jsons"
)

// JSON 使用 %v 打印(实现了 fmt.Stringer 的)返回值 将会得到 JSON
func JSON(argument any, opts ...jsons.Opt) fmt.Stringer {
	return &Arg{data: argument, opts: opts}
}

// Indent 使用 %v 打印(实现了 fmt.Stringer 的)返回值 将会得到缩进格式的 JSON
func Indent(argument any, opts ...jsons.Opt) fmt.Stringer {
	opts = append([]jsons.Opt{jsons.UseIndent("", "  ")}, opts...)
	return &Arg{data: argument, opts: opts}
}

type Arg struct {
	data any
	opts []jsons.Opt
}

// String 实现 Stringer 接口，打印自身将会返回 JSON 字符串
func (a *Arg) String() string {
	b, err := jsons.ToBytesE(a.data, a.opts...)
	if err != nil {
		return fmt.Sprintf("!(BADJSON|err=%+v|data=%#v)", err, a.data)
	}
	return conv.Bytes2String(b)
}
