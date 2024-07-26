// author gmfan
// date 2023/2/15

package structs

import (
	"encoding/json"
	"reflect"
	"unicode"
)

// IsNil 判断 val 是否为 nil
func IsNil(val any) bool {
	if val == nil {
		return true
	}

	vv := reflect.ValueOf(val)
	if vv.Kind() == reflect.Pointer {
		return vv.IsNil()
	}

	return false
}

// IsBasicType 判断 val 是否为基本类型
func IsBasicType(val any) bool {
	valueType := reflect.TypeOf(val)
	kind := valueType.Kind()

	switch kind {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		return true
	default:
		return false
	}
}

// IsSerializable 判断 val 是否可被 json.Marshal 序列化
func IsSerializable(val any) bool {
	if IsNil(val) {
		return true
	}

	switch val.(type) {
	case json.Marshaler:
		return true
	}

	v := reflect.ValueOf(val)
	v = elemIfPointer(v)
	if v.Kind() == reflect.Map {
		// Map 可以序列化
		return true
	}
	if IsBasicType(v.Interface()) {
		// 基础数据类型
		return true
	}

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if v.Len() > 0 {
			return IsSerializable(v.Index(0).Interface())
		}
		return true
	}

	// 如果是结构体类型，检查是否存在可导出字段
	t := v.Type()
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			name := t.Field(i).Name
			if unicode.IsUpper([]rune(name)[0]) {
				// 首字母大写即可导出
				return true
			}
		}
	}
	return false
}
