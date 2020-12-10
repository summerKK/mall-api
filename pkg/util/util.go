package util

import "reflect"

// 判断是否是指针结构体
func IsStructPtr(i interface{}) bool {
	typ := reflect.TypeOf(i)
	if typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct {
		return true
	}

	return false
}
