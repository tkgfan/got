// author gmfan
// date 2024/3/28

package mongos

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DuplicateKeyCode = 11000
)

func IsDuplicateKeyErr(err error) bool {
	var v mongo.CommandError
	if errors.As(err, &v) {
		return v.Code == DuplicateKeyCode
	}
	return false
}
