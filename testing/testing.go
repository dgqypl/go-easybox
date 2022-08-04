package testing

import (
	"fmt"
	"reflect"
	"testing"
)

// FuncAssertion 简化编写函数的单元测试
//
// 比如对于如下函数：
//
// func Add(a int, b int64, c string) (bool, error) {
// 	if c == "err" {
// 		return false, errors.New("c is err")
// 	}
//
// 	cNum := int64(a) + b
// 	return Int64ToString(cNum) == c, nil
// }
//
// 先定义一个测试用例的结构体：
//
// type addCase struct {
// 	A int
// 	B int64
// 	C string
// 	D bool
// 	E error
// }
//
// 然后调用本函数来简化单元测试：
func FuncAssertion(t *testing.T, fnCasePtr any, fn any) {
	t.Helper()

	ptrValue := reflect.ValueOf(fnCasePtr)
	if ptrValue.Kind() != reflect.Ptr {
		panic("fnCasePtr is not a ptr")
	}

	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		panic("fn is not a function")
	}

	_case := ptrValue.Elem()
	caseFieldNum := _case.NumField()

	fnType := fnValue.Type()
	fnInNum := fnType.NumIn()
	fnOutNum := fnType.NumOut()

	if caseFieldNum != (fnInNum + fnOutNum) {
		panic("case field num and fn in+out num not match")
	}

	for i := 0; i < fnInNum; i++ {
		caseFieldType := _case.Field(i).Type()
		fnInType := fnType.In(i)
		if caseFieldType != fnInType {
			panic(fmt.Sprintf("case %d'th field(type '%s') not math fn %d'th in(type '%s')", i+1, caseFieldType.Name(), i+1, fnInType.Name()))
		}
	}

	for i := fnInNum; i < caseFieldNum; i++ {
		caseFieldType := _case.Field(i).Type()
		fnOutType := fnType.Out(i - fnInNum)
		if caseFieldType != fnOutType {
			panic(fmt.Sprintf("case %d'th field(type '%s') not math fn %d'th out(type '%s')", i+1, caseFieldType.Name(), i-fnInNum+1, fnOutType.Name()))
		}
	}

	var caseIn []reflect.Value
	for i := 0; i < fnInNum; i++ {
		caseIn = append(caseIn, _case.Field(i))
	}
	var caseOut []reflect.Value
	for i := fnInNum; i < caseFieldNum; i++ {
		caseOut = append(caseOut, _case.Field(i))
	}

	callOut := fnValue.Call(caseIn)

	for i := 0; i < fnOutNum; i++ {
		if !reflect.DeepEqual(caseOut[i].Interface(), callOut[i].Interface()) {
			t.Fatalf("case %d'th field(type '%s') expected '%+v' , but '%+v' got", fnInNum+i+1, caseOut[i].Type().Name(), caseOut[i], callOut[i])
		}
	}

}
