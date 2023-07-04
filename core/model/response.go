// author gmfan
// date 2023/07/01

package model

import "github.com/tkgfan/got/core/structs"

type (
	SuccessResp struct {
		Code uint32 `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data"`
	}

	FailResp struct {
		Code uint32 `json:"code"`
		Msg  string `json:"msg"`
	}
)

func NewSuccessResp(code uint32, msg string, data any) *SuccessResp {
	if structs.IsNil(data) {
		data = struct{}{}
	}
	return &SuccessResp{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewFailResp(code uint32, msg string) *FailResp {
	return &FailResp{
		Code: code,
		Msg:  msg,
	}
}
