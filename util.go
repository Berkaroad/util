package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
)

// 获取当前文件执行的路径
func GetExecFilePath() string {
	fullfilename, _ := exec.LookPath(os.Args[0])
	return fullfilename
}

// 获取当前文件执行的路径
func GetExecDirPath() string {
	fullfilename := GetExecFilePath()
	fulldirname := filepath.Dir(fullfilename)
	return fulldirname
}

// 获取真实类型，而非指针
func FromPtrTypeOf(obj interface{}) reflect.Type {
	realType := reflect.TypeOf(obj)
	for realType.Kind() == reflect.Ptr {
		realType = realType.Elem()
	}
	return realType
}

// 获取真实类型，而非指针
func FromPtrType(typ reflect.Type) reflect.Type {
	realType := typ
	for realType.Kind() == reflect.Ptr {
		realType = realType.Elem()
	}
	return realType
}

// 获取真实值，而非指针
func FromPtrValueOf(obj interface{}) reflect.Value {
	realValue := reflect.ValueOf(obj)
	for realValue.Kind() == reflect.Ptr {
		realValue = realValue.Elem()
	}
	return realValue
}

// 获取真实值，而非指针
func FromPtrValue(val reflect.Value) reflect.Value {
	realValue := val
	for realValue.Kind() == reflect.Ptr {
		realValue = realValue.Elem()
	}
	return realValue
}
