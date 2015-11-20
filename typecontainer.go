package util

import (
	"reflect"
)

var TypeContainer = typeContainer{mapper: make(map[string]reflect.Type, 0)}

type typeContainer struct {
	mapper map[string]reflect.Type
}

// 注册 Type
func (self *typeContainer) RegisterType(typ reflect.Type) {
	realType := FromPtrType(typ)
	realTypeInfo := GetTypeInfo(realType)
	self.mapper[realTypeInfo.String()] = realType
	switch realType.Kind() {
	case reflect.Array, reflect.Slice:
		self.RegisterType(realType.Elem())
	}
}

// 注册 Type
func (self *typeContainer) RegisterTypeOf(obj interface{}) {
	realType := FromPtrTypeOf(obj)
	self.RegisterType(realType)
}

// 根据类型完整名称，获取 Type
func (self *typeContainer) FindType(typeInfo TypeInfo) reflect.Type {
	return self.mapper[typeInfo.String()]
}
