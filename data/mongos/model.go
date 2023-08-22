// Package mongos
// author gmfan
// date 2023/2/27
package mongos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedBy int64              `bson:"created_by,omitempty" json:"createdBy"`
	CreatedAt int64              `bson:"created_at,omitempty" json:"createdAt"`
	UpdatedBy int64              `bson:"updated_by,omitempty" json:"updatedBy"`
	UpdatedAt int64              `bson:"updated_at,omitempty" json:"updatedAt"`
	IsDeleted int8               `bson:"is_deleted" json:"isDeleted"`
}

func (b *BaseModel) CreateByUID(uid int64) {
	now := time.Now().Unix()
	b.CreatedBy = uid
	b.CreatedAt = now
	b.UpdatedBy = uid
	b.UpdatedAt = now
}

func (b *BaseModel) UpdateByUID(uid int64) {
	b.UpdatedBy = uid
	b.UpdatedAt = time.Now().Unix()
}
