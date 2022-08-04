package str

import (
	easyboxt "github.com/dgqypl/go-easybox/testing"
	"strconv"
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

type stringToInt64Case struct {
	In  string
	Out int64
	Err error
}

func TestStringToInt64(t *testing.T) {
	easyboxt.FuncAssertion(t, &stringToInt64Case{"10", 10, nil}, StringToInt64)
	easyboxt.FuncAssertion(t, &stringToInt64Case{"-1", -1, nil}, StringToInt64)
	easyboxt.FuncAssertion(t, &stringToInt64Case{"a", 0, &strconv.NumError{Func: "ParseInt", Num: "a", Err: strconv.ErrSyntax}}, StringToInt64)
}
