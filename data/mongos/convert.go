// author lby
// date 2023/4/25

package mongos

import "go.mongodb.org/mongo-driver/bson/primitive"

// ToObjectIDS 将十六进制字符串类型ID转换为primitive.ObjectID
func ToObjectIDS(strs []string) (ids []primitive.ObjectID) {
	ids = make([]primitive.ObjectID, len(strs))
	for i := 0; i < len(strs); i++ {
		ids[i].UnmarshalText([]byte(strs[i]))
	}
	return
}
