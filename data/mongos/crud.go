// author gmfan
// date 2023/6/28

package mongos

import (
	"context"
	"github.com/tkgfan/got/core/errs"
	"github.com/tkgfan/got/core/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"time"
)

// UpsertOne 如果upsert=true则有则修改，无则新增。如果 update 类型不是 bson.D
// 则会被设置为 bson.D{{"$set",update}}
func UpsertOne(ctx context.Context, table string, filter bson.D, update any, upsert bool) (*mongo.UpdateResult, error) {
	filter = wrapNeDeleted(filter)
	opts := options.Update().SetUpsert(upsert)
	if _, ok := update.(bson.D); !ok {
		update = bson.D{{"$set", update}}
	}

	result, err := DB().Collection(table).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateOne(ctx context.Context, table string, filter bson.D, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	filter = wrapNeDeleted(filter)
	return DB().Collection(table).UpdateOne(ctx, filter, update, opts...)
}

func UpdateMany(ctx context.Context, table string, filter bson.D, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	filter = wrapNeDeleted(filter)
	return DB().Collection(table).UpdateMany(ctx, filter, update, opts...)
}

func InsertOne(ctx context.Context, table string, document any, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return DB().Collection(table).InsertOne(ctx, document, opts...)
}

func InsertMany(ctx context.Context, table string, documents []any, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return DB().Collection(table).InsertMany(ctx, documents, opts...)
}

func wrapNeDeleted(filter bson.D) bson.D {
	return append(filter, bson.E{Key: "is_deleted", Value: bson.M{"$ne": 1}})
}

func elemValueIfPointer(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Pointer {
		return v.Elem()
	}
	return v
}

// UpsertMany 批量更新，filters与updates必须一致且filters与updates中元素一一对应
func UpsertMany(ctx context.Context, table string, filters []bson.D,
	updates any, upsert bool) (res *mongo.BulkWriteResult, err error) {

	upVal := reflect.ValueOf(updates)
	upVal = elemValueIfPointer(upVal)

	// 数据格式效验
	if upVal.Kind() != reflect.Slice && upVal.Kind() != reflect.Array {
		err = errs.New("updates必须为数组或切片")
		return
	} else if upVal.Len() != len(filters) {
		err = errs.New("filters与updates长度必须一致")
		return
	}

	var models []mongo.WriteModel
	for i := 0; i < len(filters); i++ {
		update := upVal.Index(i).Interface()
		model := mongo.NewUpdateOneModel().
			SetUpsert(upsert).
			SetFilter(filters[i]).
			SetUpdate(bson.D{{"$set", update}})
		models = append(models, model)
	}

	opts := options.BulkWrite().SetOrdered(false)
	res, err = DB().Collection(table).BulkWrite(ctx, models, opts)

	return
}

// DelMany 根据 filter 软删除数据
func DelMany(ctx context.Context, table string, filter bson.D, deletedBy int64) (err error) {
	update := bson.D{{"$set", bson.D{
		{"is_deleted", 1},
		{"deleted_by", deletedBy},
		{"deleted_at", time.Now().Unix()},
	}}}

	_, err = DB().Collection(table).UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}
	return
}

func handleProjections(projections []string) bson.M {
	m := bson.M{}
	for _, p := range projections {
		m[p] = 1
	}
	return m
}

// Query 简化多文档查询
func Query(ctx context.Context, table string, filter bson.D, res any, projections ...string) error {
	// 只查询未删除文档
	filter = wrapNeDeleted(filter)
	cur, err := DB().Collection(table).Find(ctx, filter, &options.FindOptions{
		Projection: handleProjections(projections),
	})
	if err != nil {
		return err
	}
	return cur.All(ctx, res)
}

// QueryOne 简化查询集合单条信息函数
func QueryOne(ctx context.Context, table string, filter bson.D, res any, projections ...string) error {
	// 只查询未删除文档
	filter = wrapNeDeleted(filter)
	return DB().Collection(table).FindOne(ctx, filter, &options.FindOneOptions{
		Projection: handleProjections(projections),
	}).Decode(res)
}

// PageQuery 分页查询
func PageQuery(ctx context.Context, table string, filter bson.D, page *model.Page, res any, projections ...string) (total int64, err error) {
	collection := DB().Collection(table)
	filter = wrapNeDeleted(filter)
	opts := &options.FindOptions{Projection: handleProjections(projections)}

	handlerPage(opts, page)

	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return
	}

	total, err = collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return total, cur.All(ctx, res)
}

func handlerPage(opts *options.FindOptions, page *model.Page) {
	// 默认每页大小
	opts.SetLimit(20)
	// 处理排序
	filter := bson.D{}
	for _, v := range page.Sorts {
		order := 1
		if v.Order == model.DESC {
			order = -1
		}
		filter = append(filter, bson.E{v.Condition, order})
	}

	if len(filter) > 0 {
		opts.SetSort(filter)
	}

	// 处理分页
	if page.Size != 0 {
		opts.SetLimit(page.Size)
		if page.Num > 0 {
			opts.SetSkip((page.Num - 1) * page.Size)
		}
	}
}

type autoIncIDModule struct {
	ID  string `bson:"_id"`
	Val int64  `bson:"val"`
}

// AutoIncID 获取自增 ID，key 为自增键。
func AutoIncID(ctx context.Context, table string, key string) (int64, error) {
	filter := bson.D{{"_id", key}}
	update := bson.D{{"$inc", bson.M{"val": 1}}}
	var res autoIncIDModule
	opts := &options.FindOneAndUpdateOptions{}
	opts.SetReturnDocument(options.After)
	opts.SetUpsert(true)
	err := DB().Collection(table).FindOneAndUpdate(ctx, filter, update, opts).Decode(&res)
	for err != nil && IsDuplicateKeyErr(err) {
		// 并发冲突，继续尝试获取自增 ID
		err = DB().Collection(table).FindOneAndUpdate(ctx, filter, update, opts).Decode(&res)
		if err == nil {
			return res.Val, nil
		}
	}
	if err != nil {
		return -1, errs.Wrap(err)
	}
	return res.Val, nil
}
