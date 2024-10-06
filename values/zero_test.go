package values_test

import (
	"testing"

	"code.gopub.tech/commons/values"
)

func TestIsZero(t *testing.T) {

	tests := []struct {
		name string
		got  bool
		want bool
	}{
		{name: "int", got: values.IsZero(0), want: true},
		{name: "str", got: values.IsZero(""), want: true},
		{name: "arr", got: values.IsZero([0]int{}), want: true},
		{name: "any", got: values.IsZero[any](nil), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.got; got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
