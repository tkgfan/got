// author gmfan
// date 2024/5/8

package spider

import (
	"compress/flate"
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type DoFunc func(req *http.Request) (*http.Response, error)

type HttpSpider struct {
	CookieMap map[string]string
	Headers   map[string]string
	HttpDo    DoFunc
}

// NewHttpSpider 创建一个 http 爬虫，do 为 nil 时默认使用 http.DefaultClient.Do
func NewHttpSpider(do DoFunc) *HttpSpider {
	if do == nil {
		do = http.DefaultClient.Do
	}
	return &HttpSpider{
		CookieMap: make(map[string]string),
		Headers:   make(map[string]string),
		HttpDo:    do,
	}
}

func (b *HttpSpider) getCookies() string {
	var arr []string
	for k, v := range b.CookieMap {
		arr = append(arr, k+"="+v)
	}
	return strings.Join(arr, "; ")
}

func (b *HttpSpider) ReadBody(resp *http.Response) ([]byte, error) {
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
	case "deflate":
		reader = flate.NewReader(resp.Body)
	default:
		reader = resp.Body
	}
	defer reader.Close() // nolint

	return io.ReadAll(reader)
}

func (b *HttpSpider) Do(req *http.Request) (*http.Response, error) {
	if b.Headers["Referer"] == "" {
		b.Headers["Referer"] = req.URL.String()
	}
	if b.Headers["User-Agent"] == "" {
		b.Headers["User-Agent"] = RandomUA()
	}
	req.Header.Set("Cookie", b.getCookies())
	for k, v := range b.Headers {
		req.Header.Set(k, v)
	}

	resp, err := b.HttpDo(req)
	if err != nil {
		return nil, err
	}
	err = b.handleResp(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (b *HttpSpider) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return b.Do(req)
}

// 处理返回结果，更新 cookie 上下文
func (b *HttpSpider) handleResp(resp *http.Response) (err error) {
	// 提取 cookie
	cookies := resp.Cookies()
	for _, c := range cookies {
		b.CookieMap[c.Name] = c.Value
	}
	return
}
