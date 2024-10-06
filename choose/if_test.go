package choose_test

import (
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/commons/choose"
)

type User struct {
	ID  int64
	Age int64
}

func (u *User) GetID() int64  { return u.ID }
func (u *User) GetAge() int64 { return u.Age }

func TestChoose(t *testing.T) {
	s := choose.If(true, "T", "F")
	assert.Equal(t, s, "T")
	s = choose.If(false, "T", "F")
	assert.Equal(t, s, "F")

	s = choose.IfLazy(true, func() string { return "T" }, func() string { panic("F") })
	assert.Equal(t, s, "T")
	s = choose.IfLazy(false, func() string { panic("T") }, func() string { return "F" })
	assert.Equal(t, s, "F")

	var u *User
	id := choose.IfLazyT(u != nil, u.GetID, -1)
	assert.Equal(t, id, -1)
	id = choose.IfLazyF(u == nil, -1, u.GetID)
	assert.Equal(t, id, -1)

	u = &User{ID: 1, Age: 18}
	age := choose.IfLazyT(u != nil, u.GetAge, 20)
	assert.Equal(t, age, 18)
	age = choose.IfLazyF(u == nil, 20, u.GetAge)
	assert.Equal(t, age, 18)

}
