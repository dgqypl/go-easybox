package str

import (
	easyboxt "github.com/dgqypl/go-easybox/testing"
	"testing"
)

type int64ToStringCase struct {
	In  int64
	Out string
}

func TestInt64ToString(t *testing.T) {
	easyboxt.FuncAssertion(t, &int64ToStringCase{10, "10"}, Int64ToString)
	easyboxt.FuncAssertion(t, &int64ToStringCase{-1, "-1"}, Int64ToString)
}
