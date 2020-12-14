package util

import (
	"reflect"
)

// 判断是否是指针结构体
func IsStructPtr(i interface{}) bool {
	typ := reflect.TypeOf(i)
	if typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct {
		return true
	}

	return false
}

func IsSliceElemPtr(i interface{}) bool {
	typ := reflect.TypeOf(i)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.Ptr && typ.Elem().Elem().Kind() == reflect.Struct {
		return true
	}

	return false
}
