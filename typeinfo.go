package util

import (
	"crypto/md5"
	"github.com/berkaroad/uuid"
	"reflect"
	"sort"
)

// 类型信息
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

func (t TypeInfo) HashCode() uuid.UUID {
	return uuid.UUID(md5.Sum([]byte(t.String())))
}

// 类型成员信息
type TypeMemberInfo struct {
	TypeMemberName     string
	TypeMemberTypeName string
}

func (self TypeMemberInfo) String() string {
	return self.TypeMemberName + " " + self.TypeMemberTypeName
}

// 获取类型的类型信息
func GetTypeInfo(typ reflect.Type) TypeInfo {
	realType := FromPtrType(typ)
	switch realType.Kind() {
	case reflect.Struct:
		typeMemberNameList := make([]string, realType.NumField())
		typeMemberInfoList := make([]TypeMemberInfo, realType.NumField())
		for i := 0; i < len(typeMemberNameList); i++ {
			typeMemberNameList[i] = realType.Field(i).Name
		}
		sort.Strings(typeMemberNameList)
		for i, typeMemberName := range typeMemberNameList {
			typeMember, _ := realType.FieldByName(typeMemberName)
			typeMemberInfoList[i] = TypeMemberInfo{TypeMemberName: typeMember.Name, TypeMemberTypeName: typeMember.Type.String()}
		}
		return TypeInfo{TypeName: realType.Name(), TypeMemberInfos: typeMemberInfoList}
	case reflect.Interface:
		typeMemberNameList := make([]string, realType.NumMethod())
		typeMemberInfoList := make([]TypeMemberInfo, realType.NumMethod())
		for i := 0; i < len(typeMemberNameList); i++ {
			typeMemberNameList[i] = realType.Method(i).Name
		}
		sort.Strings(typeMemberNameList)
		for i, typeMemberName := range typeMemberNameList {
			typeMember, _ := realType.MethodByName(typeMemberName)
			typeMemberInfoList[i] = TypeMemberInfo{TypeMemberName: typeMember.Name, TypeMemberTypeName: typeMember.Type.String()}
		}
		return TypeInfo{TypeName: realType.Name(), TypeMemberInfos: typeMemberInfoList}
	default:
		return TypeInfo{TypeName: realType.Name()}
	}

}

// 获取对象的类型信息
func GetTypeInfoOf(obj interface{}) TypeInfo {
	realType := FromPtrTypeOf(obj)
	return GetTypeInfo(realType)
}
