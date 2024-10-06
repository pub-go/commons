package assert

import (
	"reflect"
	"testing"

	"code.gopub.tech/commons/fmts"
	"code.gopub.tech/commons/values"
)

func Equal[T comparable](t *testing.T, got, want T, msgs ...any) {
	if got != want { // 使用泛型确保入参类型一致, 避免 int(0) != int64(0) 的情况
		t.Helper()
		t.Errorf("[%v] assert equal failed: got %v(%T), want %v(%T).%s",
			t.Name(), got, got, want, want, msg(msgs...))
	}
}

func DeepEqual(t *testing.T, got, want any, msgs ...any) {
	if !reflect.DeepEqual(got, want) {
		t.Helper()
		t.Errorf("[%v] assert deep equal failed: got %v(%T), want %v(%T).%s",
			t.Name(), got, got, want, want, msg(msgs...))
	}
}

func True(t *testing.T, cond bool, msgs ...any) {
	if !cond {
		t.Helper()
		t.Errorf("[%v] assert true failed: got %v(%T).%s",
			t.Name(), cond, cond, msg(msgs...))
	}
}

func False(t *testing.T, cond bool, msgs ...any) {
	if cond {
		t.Helper()
		t.Errorf("[%v] assert false failed: got %v(%T).%s",
			t.Name(), cond, cond, msg(msgs...))
	}
}

func Nil(t *testing.T, a any, msgs ...any) {
	if !values.IsNil(a) {
		t.Helper()
		t.Errorf("[%v] assert nil failed: got %v(%T).%s",
			t.Name(), a, a, msg(msgs...))
	}
}

func NotNil(t *testing.T, a any, msgs ...any) {
	if values.IsNil(a) {
		t.Helper()
		t.Errorf("[%v] assert not nil failed: got %v(%T).%s",
			t.Name(), a, a, msg(msgs...))
	}
}

func ShouldNotPanic(t *testing.T, fn func(), msgs ...any) {
	t.Helper()
	defer func() {
		if x := recover(); x != nil {
			t.Errorf("[%v] assert should not panic failed: got panic %v(%T).%s",
				t.Name(), x, x, msg(append(msgs, x)...))
		}
	}()
	fn()
}

func ShouldPanic(t *testing.T, fn func(), msgs ...any) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Errorf("[%v] assert should panic failed: not panic.%s",
				t.Name(), msg(msgs...))
		}
	}()
	fn()
}

func msg(msgs ...any) string {
	if len(msgs) > 0 {
		if format, ok := msgs[0].(string); ok {
			return fmts.Sprintf(" "+format, msgs[1:]...)
		}
	}
	return ""
}
