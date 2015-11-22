package util

import (
	"encoding/gob"
	"reflect"
)

// 类型容器
var TypeContainer = typeContainer{mapper: make(map[string]reflect.Type, 0)}

type typeContainer struct {
	mapper map[string]reflect.Type
}

// 注册 Type
func (self *typeContainer) RegisterType(typ reflect.Type) {
	realType := FromPtrType(typ)
	realTypeHashCode := GetTypeInfo(realType).HashCode().String()
	if self.mapper[realTypeHashCode] == nil {
		self.mapper[realTypeHashCode] = realType
		gob.RegisterName(realTypeHashCode, reflect.New(realType).Interface())
		switch realType.Kind() {
		case reflect.Array, reflect.Slice:
			self.RegisterType(realType.Elem())
		}
	}
}

// 注册 Type
func (self *typeContainer) RegisterTypeOf(obj interface{}) {
	realType := FromPtrTypeOf(obj)
	self.RegisterType(realType)
}

// 根据类型完整名称，获取 Type
func (self *typeContainer) FindType(typeInfo TypeInfo) reflect.Type {
	realTypeHashCode := typeInfo.HashCode().String()
	return self.mapper[realTypeHashCode]
}
