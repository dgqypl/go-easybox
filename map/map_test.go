package _map

import (
	easyboxt "github.com/dgqypl/go-easybox/testing"
	"testing"
)

type sliceToMapCase[T any, k comparable] struct {
	Slice []T
	F     func(T) k
	Out   map[k]T
}

func TestSliceToMap(t *testing.T) {
	var slice []*User
	u1 := &User{Id: 1, Name: "helen"}
	u2 := &User{Id: 2, Name: "tim"}

	slice = append(slice, u1, u2)

	f := func(u *User) int64 {
		return u.Id
	}

	easyboxt.FuncAssertion(t, &sliceToMapCase[*User, int64]{slice, f, map[int64]*User{1: u1, 2: u2}}, SliceToMap[*User, int64])
}
