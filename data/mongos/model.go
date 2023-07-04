// Package mongos
// author gmfan
// date 2023/2/27
package mongos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedBy int64              `bson:"created_by,omitempty"`
	CreatedAt int64              `bson:"created_at,omitempty"`
	UpdatedBy int64              `bson:"updated_by,omitempty"`
	UpdatedAt int64              `bson:"updated_at,omitempty"`
	IsDeleted int8               `bson:"is_deleted"`
}

func CreateBaseModel(uid int64) BaseModel {
	return BaseModel{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now().Unix(),
		CreatedBy: time.Now().Unix(),
		UpdatedBy: uid,
		UpdatedAt: time.Now().Unix(),
	}
}

func (b *BaseModel) Update(uid int64) {
	b.UpdatedBy = uid
	b.UpdatedAt = time.Now().Unix()
}
