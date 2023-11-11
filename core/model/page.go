// author gmfan
// date 2023/6/27
package model

import "github.com/tkgfan/got/core/slices"

const (
	// ASC 升序
	ASC = "asc"
	// DESC 降序
	DESC = "desc"
)

type (
	// Page 分页模型，推荐 Num 从第一页开始
	Page struct {
		Num   int64  `json:"num"`
		Size  int64  `json:"size"`
		Sorts []Sort `json:"sorts"`
	}

	Sort struct {
		Condition string `json:"condition"`
		Order     string `json:"order"`
	}

	PageResult struct {
		Rows  any   `json:"rows"`
		Total int64 `json:"total"`
	}
)

func NewPageResult(rows any, total int64) *PageResult {
	return &PageResult{
		Rows:  slices.ToInterfaceSlice(rows),
		Total: total,
	}
}
