package values_test

import (
	"testing"

	"code.gopub.tech/commons/values"
)

type MyError struct{}

func (e *MyError) Error() string { return "" }

func foo() *MyError { return nil }

func TestIsNil(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "nil字面量", args: args{v: nil}, want: true},
		{name: "custom-err", args: args{v: foo()}, want: true},
		{name: "int", args: args{v: 0}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := values.IsNil(tt.args.v); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
