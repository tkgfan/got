// author lby
// date 2024/4/18

package fs

import (
	"encoding/json"
	"io"
	"os"
)

// SaveJSON 保存 JSON 对象到文件中
func SaveJSON(fileName string, data any) (err error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return
	}
	dstFile, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer dstFile.Close()
	_, err = dstFile.Write(bs)
	return
}

// LoadJSON 加载文件并反序列化到 res 中
func LoadJSON(fileName string, res any) (err error) {
	fs, err := os.Open(fileName)
	if err != nil {
		return
	}
	bs, err := io.ReadAll(fs)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, res)
	if err != nil {
		return
	}
	return
}
