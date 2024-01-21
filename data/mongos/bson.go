// author lby
// date 2023/6/28

package mongos

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strings"
)

const (
	MODEL1 = 1
	MODEL2 = 2
	MODEL3 = 3
	inline = ",inline"
)

// StructureBsonD 构建bson D对象同时解析含_id后缀的值为ObjectID），MODEL1 直接构建不做处理
// MODEL2 跳过默认值构建，MODEL3跳过默认值其中整型的默认值为-1
func StructureBsonD(arg any, model int) bson.D {
	res := bson.D{}

	argT := reflect.TypeOf(arg)
	argV := reflect.ValueOf(arg)
	// 指针变量需要解
	if argT.Kind() == reflect.Pointer {
		argT = argT.Elem()
		argV = argV.Elem()
	}

	// 遍历所有字段
	for i := 0; i < argT.NumField(); i++ {
		sf := argT.Field(i)
		val := argV.Field(i).Interface()

		// 开启默认值不构建bson
		if model != MODEL1 {
			isDefault := false
			// 类型过滤
			if model == MODEL2 {
				isDefault = handleModel2(val)
			} else if model == MODEL3 {
				isDefault = handleModel3(val)
			}

			if isDefault {
				continue
			}
		}

		tag := getTagValue(sf.Tag)

		//参考文档 https://www.mongodb.com/docs/drivers/go/v1.9/fundamentals/bson/#struct-tags
		if tag == inline {
			inlineB := StructureBsonD(val, model)
			res = append(res, inlineB...)
		} else {
			if strings.Contains(tag, "_id") {
				switch val.(type) {
				case string:
					val, _ = primitive.ObjectIDFromHex(val.(string))
				}
			}
			res = append(res, bson.E{Key: tag, Value: val})
		}
	}

	return res
}

func getTagValue(st reflect.StructTag) string {
	tag := st.Get("bson")
	if tag == "" {
		tag = st.Get("json")
	}
	if tag == inline {
		return tag
	}
	splitIndex := strings.Index(tag, ",")
	if splitIndex != -1 {
		tag = tag[:splitIndex]
	}

	return tag
}

func handleModel2(val interface{}) (isDefault bool) {
	if handleIfArrayOrSlice(val) {
		return true
	}

	switch val.(type) {
	case int8:
		isDefault = val.(int8) == 0
	case int16:
		isDefault = val.(int16) == 0
	case int32:
		isDefault = val.(int32) == 0
	case int64:
		isDefault = val.(int64) == 0
	case int:
		isDefault = val.(int) == 0
	case string:
		isDefault = val.(string) == ""
	case primitive.ObjectID:
		isDefault = val.(primitive.ObjectID) == primitive.NilObjectID
	case []int:
		isDefault = len(val.([]int)) == 0
	case []int8:
		isDefault = len(val.([]int8)) == 0
	case map[string]int:
		isDefault = len(val.(map[string]int)) == 0
	case map[int]string:
		isDefault = len(val.(map[int]string)) == 0
	case map[int64]int64:
		isDefault = len(val.(map[int64]int64)) == 0
	case map[string]interface{}:
		isDefault = len(val.(map[string]interface{})) == 0
	}

	return isDefault
}

func handleModel3(val interface{}) (isDefault bool) {
	if handleIfArrayOrSlice(val) {
		return true
	}

	switch val.(type) {
	case int8:
		isDefault = val.(int8) == -1
	case int16:
		isDefault = val.(int16) == -1
	case int32:
		isDefault = val.(int32) == -1
	case int64:
		isDefault = val.(int64) == -1
	case int:
		isDefault = val.(int) == -1
	case string:
		isDefault = val.(string) == ""
	case primitive.ObjectID:
		isDefault = val.(primitive.ObjectID) == primitive.NilObjectID
	case []int:
		isDefault = len(val.([]int)) == 0
	case []int8:
		isDefault = len(val.([]int8)) == 0
	case map[string]int:
		isDefault = len(val.(map[string]int)) == 0
	case map[int64]int64:
		isDefault = len(val.(map[int64]int64)) == 0
	case map[string]interface{}:
		isDefault = len(val.(map[string]interface{})) == 0
	}

	return isDefault
}

func handleIfArrayOrSlice(val any) bool {
	rv := reflect.ValueOf(val)
	elemValueIfPointer(rv)
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		return rv.Len() == 0
	}
	return false
}
