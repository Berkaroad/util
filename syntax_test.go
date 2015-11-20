package util

import (
	"testing"
)

type someTypeForTestSyntax struct {
	isInitialized bool
	field01       string
}

func (self *someTypeForTestSyntax) InitFunc() interface{} {
	return func(field01 string) {
		if !self.isInitialized {
			// 初始化
			self.field01 = field01
		}
		self.isInitialized = true
	}
}

func (self *someTypeForTestSyntax) Method1string_int(para1 string, para2 int) string {
	return "Method1string_int:" + para1
}
func (self *someTypeForTestSyntax) Method1string(para1 string) string {
	return "Method1string:" + para1
}

func Test_InvokeOverloadMethod(t *testing.T) {
	obj := new(someTypeForTestSyntax)
	returnVal, _ := InvokeOverloadMethod(obj, "Method1", "hello")[0].Interface().(string)
	if returnVal != "Method1string:hello" {
		t.Fatal("InvokeOverloadMethod 'Method1' error!")
	}
	returnVal, _ = InvokeOverloadMethod(obj, "Method1", "hello", 9)[0].Interface().(string)
	if returnVal != "Method1string_int:hello" {
		t.Fatal("InvokeOverloadMethod 'Method1' error!")
	}
}
