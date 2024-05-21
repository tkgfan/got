// author lby
// date 2023/10/31

package logx

import "testing"

func TestInfo(t *testing.T) {
	Info("info message")
}

func TestInfof(t *testing.T) {
	Infof("format: %s", "message")
}

func TestError(t *testing.T) {
	Error("output red message")
}
