package util

import (
	"testing"
)

type someTypeForTestMarshal struct {
	Name string
	Age  int
}

func Test_Marshal(t *testing.T) {
	TypeContainer.RegisterTypeOf((*someTypeForTestMarshal)(nil))
	obj := someTypeForTestMarshal{Name: "Jerry Bai", Age: 30}
	if data, err := MarshalObj(obj); err != nil {
		t.Fatalf("TestMarshal error: %s!", err.Error())
	} else {
		if objInterface, err := UnmarshalObj(data); err != nil {
			t.Fatalf("TestMarshal error: %s!", err.Error())
		} else {
			if newObj, ok := objInterface.(*someTypeForTestMarshal); ok {
				if newObj.Age != 30 || newObj.Name != "Jerry Bai" {
					t.Fatal("The object info from unmarshal not equal with original one!")
				}
			} else {
				t.Fatal("Type asset error!")
			}
		}
	}
}
