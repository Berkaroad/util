package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type objWrapper struct {
	Type       TypeInfo `json:"type"`
	ObjectData []byte   `json:"obj"`
}

// 序列化 Object
func MarshalObj(obj interface{}) ([]byte, error) {
	typ := FromPtrTypeOf(obj)
	typeInfo := GetTypeInfo(typ)
	if data, err := json.Marshal(obj); err != nil {
		return nil, err
	} else {
		wrapper := objWrapper{Type: typeInfo, ObjectData: data}
		return json.Marshal(wrapper)
	}
}

// 反序列化 Object
func UnmarshalObj(data []byte) (interface{}, error) {
	wrapper := objWrapper{}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	} else {
		typ := TypeContainer.FindType(wrapper.Type)
		if typ != nil {
			obj := reflect.New(typ).Interface()
			err = json.Unmarshal(wrapper.ObjectData, &obj)
			return obj, err
		} else {
			return nil, errors.New(fmt.Sprintf("Object Type '%s' is not found.", wrapper.Type))
		}
	}
}
