package fmts_test

import (
	"testing"

	"code.gopub.tech/commons/fmts"
)

func TestFormat(t *testing.T) {
	s := fmts.Sprintf("a=%v")
	t.Logf("%v", s) // a=%v
	s = fmts.Sprintf("a=%v", 1, 2)
	t.Logf("%v", s) // a=1
	// 仅处理 参数多于占位符 的情况; 参数更少可能不符合预期
	s = fmts.Sprintf("a=%v,b=%v,c=%v,d=%v", 1, 2)
	t.Logf("%v", s) // a=1,b=2,c=,d=%!v(MISSING)
}
