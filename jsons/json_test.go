package jsons_test

import (
	"testing"

	"code.gopub.tech/commons/jsons"
)

type User struct {
	ID      int64
	Name    string
	Friends []*User
}

func TestJSON(t *testing.T) {
	var user = &User{}
	s := jsons.ToJSON(user)
	t.Logf("%v", s)
	s = jsons.Indent(user)
	t.Logf("\n%v", s)
}
