package testing

import (
	"errors"
	. "github.com/dgqypl/go-easybox/str"
)

func Add(a int, b int64, c string) (bool, error) {
	if c == "err" {
		return false, errors.New("c is err")
	}

	cNum := int64(a) + b
	return Int64ToString(cNum) == c, nil
}

type addCase struct {
	A int
	B int64
	C string
	D bool
	E error
}

func ExampleFuncAssertion() {
	FuncAssertion(t, &addCase{1, 2, "3", true, nil}, Add)
	FuncAssertion(t, &addCase{1, 2, "4", false, nil}, Add)
	FuncAssertion(t, &addCase{1, 2, "err", false, errors.New("c is err")}, Add)
}
