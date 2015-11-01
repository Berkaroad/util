package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type objWrapper struct {
	Type   string `json:"type"`
	Object string `json:"obj"`
}

// 序列化 Object
func MarshalObj(obj interface{}) ([]byte, error) {
	typeName := FromPtrTypeOf(obj).Name()
	if data, err := json.Marshal(obj); err != nil {
		return nil, err
	} else {
		wrapper := objWrapper{Type: typeName, Object: string(data)}
		return json.Marshal(wrapper)
	}
}

// 反序列化 Object
func UnmarshalObj(data []byte, findType func(typeName string) reflect.Type) (interface{}, error) {
	wrapper := objWrapper{}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	} else {
		typ := findType(wrapper.Type)
		if typ != nil {
			obj := reflect.New(typ).Interface()
			err = json.Unmarshal([]byte(wrapper.Object), &obj)
			return obj, err
		} else {
			return nil, errors.New(fmt.Sprintf("Object Type \"%s\" not found.", wrapper.Type))
		}
	}
}
