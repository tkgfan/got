// Package mongos
// author gmfan
// date 2023/3/2
package mongos

import (
	"bytes"
	"github.com/tkgfan/got/core/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Bucket(table string) (bucket *gridfs.Bucket, err error) {
	if table == "" {
		bucket, err = gridfs.NewBucket(DB())
	} else {
		bucketOptions := options.GridFSBucket().SetName(table)
		bucket, err = gridfs.NewBucket(DB(), bucketOptions)
	}

	if err != nil {
		err = errors.Wrapf(err, "创建 Bucket 失败，参数:%+v", table)
	}

	return
}

// UploadFile 上传文件，上传成功返回 fid
func UploadFile(table string, data []byte) (fid string, err error) {
	// 获取 bucket
	bucket, err := Bucket(table)
	if err != nil {
		return
	}

	// 上传文件
	id, err := bucket.UploadFromStream(time.Now().String(), bytes.NewBuffer(data))
	if err != nil {
		return fid, errors.Wrapf(err, table)
	}

	return id.Hex(), nil
}

// DownloadFile 下载文件
func DownloadFile(table, fid string) (data []byte, err error) {
	bucket, err := Bucket(table)
	if err != nil {
		return
	}

	id, _ := primitive.ObjectIDFromHex(fid)

	buffer := bytes.NewBuffer(make([]byte, 5<<20))
	_, err = bucket.DownloadToStream(id, buffer)
	if err != nil {
		return nil, errors.Wrapf(err, "下载文件失败 fid=%s", fid)
	}

	return buffer.Bytes(), nil
}
