// Package structs
// author gmfan
// date 2022/8/31

package structs

import (
	"reflect"
)

// CopyFields 复制结构体 src 到 dst 中，会尽可能复制名称相同的字段。
// 如果 src 为切片或数组，则 dst 必须也为切片或数组。需要注意的是 dst
// 必须为指针类型
func CopyFields(dst, src any) {
	copyCheck(dst, src)

	if isArrOrSlice(dst) {
		// dst 为数组或切片
		copyArrayOrSlice(dst, src)
		return
	}

	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	copyValue(dstVal, srcVal)
}

// 检查 dst 与 src 是否能进行复制，若不能复制则会 panic
func copyCheck(dst, src any) {
	// dst 必须为指针类型
	dstType := reflect.ValueOf(dst)
	if dstType.Kind() != reflect.Pointer {
		panic("dst 必须为 Pointer")
	}

	dstType = dstType.Elem()
	srcType := reflect.TypeOf(src)

	dstKind := dstType.Kind()
	srcKind := srcType.Kind()

	// src 为数组或切片
	if srcKind == reflect.Array || srcKind == reflect.Slice {
		if dstKind != reflect.Array && dstKind != reflect.Slice {
			panic("src 为数组或切片，dst 也必须是数组或切片")
		}
	}
}

// 复制数组与切片
func copyArrayOrSlice(dst, src any) {
	dstValRes := reflect.ValueOf(dst)
	dstVal := elemIfPointer(dstValRes)
	srcVal := elemIfPointer(reflect.ValueOf(src))

	elementType := dstVal.Type().Elem()

	for i := 0; i < srcVal.Len(); i++ {
		// 扩容
		if dstVal.Len() == i {
			newElem := reflect.New(elementType).Elem()
			dstVal = reflect.Append(dstVal, newElem)
		}

		copyValue(dstVal.Index(i), srcVal.Index(i))
	}

	if dstVal.Kind() == reflect.Array {
		dstValRes.Elem().Set(dstVal)
	} else {
		dstValRes.Elem().Set(dstVal.Slice(0, srcVal.Len()))
	}
	return
}

// 复制src的值到dst
func copyValue(dst, src reflect.Value) {
	dst = elemIfPointer(dst)
	src = elemIfPointer(src)

	// 接口可以直接赋值
	if dst.Kind() == reflect.Interface {
		dst.Set(src)
		return
	}

	if dst.Kind() != reflect.Struct {
		panic("dst 不为结构体")
	}

	if src.Kind() != reflect.Struct {
		panic("src不为结构体")
	}

	srcType := src.Type()
	for i := 0; i < src.NumField(); i++ {
		stf := srcType.Field(i)
		df := dst.FieldByName(stf.Name)
		if ok := df.IsValid(); !ok {
			// stf 可能是匿名结构体
			if stf.Type.Kind() == reflect.Struct {
				copyValue(dst, src.Field(i))
			}
			continue
		}
		// 如果 dst 是 interface 则直接复制
		if df.Kind() != reflect.Interface {
			// 类型不一致
			if df.Kind() != src.Field(i).Kind() {
				continue
			}

			// 如果字段是数组或切片需要检查元素类型是否一致，否则跳过
			if src.Field(i).Kind() == reflect.Array || src.Field(i).Kind() == reflect.Slice {
				if src.Field(i).Type() != df.Type() {
					continue
				}
			}
		}

		df.Set(src.Field(i))
	}
}

// 类型检查，如果 dst 为数组或切片则返回 true。
func isArrOrSlice(dst any) bool {
	dstType := reflect.ValueOf(dst)
	dstType = dstType.Elem()
	dstKind := dstType.Kind()
	if dstKind == reflect.Array || dstKind == reflect.Slice {
		// dst 为数组或切片
		return true
	}
	return false
}

// 如果v是Pointer则返回其指向的值
func elemIfPointer(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Pointer {
		return v.Elem()
	}
	return v
}
