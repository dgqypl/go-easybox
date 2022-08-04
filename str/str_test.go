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

type intSliceToStringCase struct {
	IntSlice []int
	Sep      string
	Out      string
}

func TestIntSliceToString(t *testing.T) {
	easyboxt.FuncAssertion(t, &intSliceToStringCase{[]int{1, 2, 3}, ",", "1,2,3"}, IntSliceToString)
}

type intSlice64ToStringCase struct {
	Int64Slice []int64
	Sep        string
	Out        string
}

func TestIntSlice64ToString(t *testing.T) {
	easyboxt.FuncAssertion(t, &intSlice64ToStringCase{[]int64{1, 2, 3}, ",", "1,2,3"}, Int64SliceToString)
}

type stringToIntSliceCase struct {
	Str string
	Sep string
	Out []int
	Err error
}

func TestStringToIntSlice(t *testing.T) {
	easyboxt.FuncAssertion(t, &stringToIntSliceCase{"1,2,3", ",", []int{1, 2, 3}, nil}, StringToIntSlice)
}

type stringToInt64SliceCase struct {
	Str string
	Sep string
	Out []int64
	Err error
}

func TestStringToInt64Slice(t *testing.T) {
	easyboxt.FuncAssertion(t, &stringToInt64SliceCase{"1,2,3", ",", []int64{1, 2, 3}, nil}, StringToInt64Slice)
}
