package arg_test

import (
	"testing"

	"code.gopub.tech/commons/arg"
)

func TestArgs(t *testing.T) {
	data := map[string]any{"key": true}
	t.Logf("%v", data)
	t.Logf("%v", arg.JSON(data))
	t.Logf("%v", arg.Indent(data))
	t.Logf("%v", arg.JSON(arg.JSON))
}
