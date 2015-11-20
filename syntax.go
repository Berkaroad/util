package util

import (
	"reflect"
)

// 构造函数
type Initializer interface {
	InitFunc() interface{}
}

// 调用重载方法
func InvokeOverloadMethod(obj interface{}, methodName string, params ...interface{}) []reflect.Value {
	if methodName == "" {
		panic("Method name is empty!")
	}
	realMethodName := methodName
	paramVals := []reflect.Value{}
	if len(params) > 0 {
		paramVals = make([]reflect.Value, len(params))
		for i, param := range params {
			realMethodName = realMethodName + FromPtrTypeOf(param).Name() + "|"
			paramVals[i] = FromPtrValueOf(param)
		}
		realMethodName = realMethodName[:len(realMethodName)-1]
	}
	method := reflect.ValueOf(obj).MethodByName(realMethodName)
	if !method.IsValid() {
		method = FromPtrValueOf(obj).MethodByName(realMethodName)
	}
	if method.IsValid() {
		return method.Call(paramVals)
	} else {
		panic("Method:'" + realMethodName + "' not found!")
	}
}
