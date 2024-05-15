// author lby
// date 2024/5/8

package spider

import (
	"testing"
)

func TestBrowser_Get(t *testing.T) {
	br := NewHttpSpider(nil)
	resp, err := br.Get("https://baidu.com")
	if err != nil {
		t.Error(err)
		return
	}
	bs, err := br.ReadBody(resp)
	if err != nil {
		t.Error(err)
		return
	}
	if len(bs) == 0 {
		t.Error("请求 https://baidu.com 失败")
	}
}
