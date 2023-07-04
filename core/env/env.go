// Package env
// author gmfan
// date 2023/2/28
package env

func init() {
	LoadStr(&CurModel, ModelKey, false)
}

// 服务部署环境
const (
	ModelKey = "MODEL"
	// DevModel 开发环境
	DevModel = "dev"
	// TestModel 测试环境
	TestModel = "test"
	// ProdModel 生产环境
	ProdModel = "prod"
)

// CurModel 当前环境
var CurModel = DevModel
