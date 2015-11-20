package util

import (
	"reflect"
)

type TypeInfo struct {
	TypeName        string
	TypeMemberInfos []TypeMemberInfo
}

func (self TypeInfo) String() string {
	typeInfoStr := self.TypeName
	for _, memberInfo := range self.TypeMemberInfos {
		typeInfoStr = typeInfoStr + "|" + memberInfo.String()
	}
	return typeInfoStr
}

type TypeMemberInfo struct {
	TypeMemberName     string
	TypeMemberTypeName string
}

func (self TypeMemberInfo) String() string {
	return self.TypeMemberName + " " + self.TypeMemberTypeName
}

func GetTypeInfo(typ reflect.Type) TypeInfo {
	realType := FromPtrType(typ)
	switch realType.Kind() {
	case reflect.Struct:
		typeMemberInfoList := make([]TypeMemberInfo, realType.NumField())
		for i := 0; i < len(typeMemberInfoList); i++ {
			typeMemberInfoList[i] = TypeMemberInfo{TypeMemberName: realType.Field(i).Name, TypeMemberTypeName: realType.Field(i).Type.String()}
		}
		return TypeInfo{TypeName: realType.Name(), TypeMemberInfos: typeMemberInfoList}
	case reflect.Interface:
		typeMemberInfoList := make([]TypeMemberInfo, realType.NumMethod())
		for i := 0; i < len(typeMemberInfoList); i++ {
			typeMemberInfoList[i] = TypeMemberInfo{TypeMemberName: realType.Method(i).Name, TypeMemberTypeName: realType.Method(i).Type.String()}
		}
		return TypeInfo{TypeName: realType.Name(), TypeMemberInfos: typeMemberInfoList}
	default:
		return TypeInfo{TypeName: realType.Name()}
	}

}

func GetTypeInfoOf(obj interface{}) TypeInfo {
	realType := FromPtrTypeOf(obj)
	return GetTypeInfo(realType)
}
