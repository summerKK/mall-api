package util

import (
	"reflect"
	"time"
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

func TimeFormat(t time.Time) string {
	var timeString = t.Format("2006/01/02 15:04:05")

	return timeString
}
