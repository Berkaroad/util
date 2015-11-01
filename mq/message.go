package mq

import (
	"github.com/berkaroad/util"
	"reflect"
)

type Message interface{}

var registerdTypeGallary map[string]reflect.Type = make(map[string]reflect.Type, 0)

// 注册 MessageType
func RegisterMessageType(typ reflect.Type) {
	realType := util.FromPtrType(typ)
	registerdTypeGallary[realType.Name()] = realType
	consoleLog.Printf("[info] RegisterMessageType:%s.\n", realType.Name())
}

// 注册 MessageType
func RegisterMessageTypeOf(obj interface{}) {
	realType := util.FromPtrTypeOf(obj)
	registerdTypeGallary[realType.Name()] = realType
	consoleLog.Printf("[info] RegisterMessageType:%s.\n", realType.Name())
}

func findType(typeName string) reflect.Type {
	return registerdTypeGallary[typeName]
}

// 序列化 Message
func MarshalMessage(message Message) ([]byte, error) {
	return util.MarshalObj(message)
}

// 反序列化 Message
func UnmarshalMessage(data []byte) (Message, error) {
	if obj, err := util.UnmarshalObj(data, findType); err == nil {
		return Message(obj), nil
	} else {
		return nil, err
	}
}
