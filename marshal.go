package util

import (
	"bytes"
	"encoding/gob"
)

type objWrapper struct {
	Obj interface{}
}

// 序列化 Object
func MarshalObj(obj interface{}) ([]byte, error) {
	wrapper := objWrapper{Obj: obj}
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(wrapper); err != nil {
		return nil, err
	} else {
		return buf.Bytes(), nil
	}
}

// 反序列化 Object
func UnmarshalObj(data []byte) (interface{}, error) {
	wrapper := objWrapper{}
	if err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&wrapper); err != nil {
		return nil, err
	} else {
		return wrapper.Obj, nil
	}
}
