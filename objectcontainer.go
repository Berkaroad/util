package util

import (
	"github.com/codegangsta/inject"
	"reflect"
)

const (
	Lifecycle_Singleton = 0
	Lifecycle_Transient = 1
)

// 对象容器
var ObjectContainer = objectContainer{i: inject.New()}

type objectContainer struct {
	i inject.Injector
}

// 映射对象
func (self *objectContainer) Map(val interface{}) {
	if initializer, ok := val.(Initializer); ok {
		self.i.Invoke(initializer.InitFunc())
	}
	self.i.Map(val)
}

// 映射对象到指定类型
func (self *objectContainer) MapTo(val interface{}, typ interface{}) {
	if initializer, ok := val.(Initializer); ok {
		self.i.Invoke(initializer.InitFunc())
	}
	self.i.MapTo(val, typ)
}

// 调用函数，通过调用初始化的函数完成注入操作
func (self *objectContainer) Invoke(fn interface{}) ([]reflect.Value, error) {
	return self.i.Invoke(fn)
}
