// author gmfan
// date 2023/12/4

package concurrent

import (
	"testing"
	"time"
)

func TestNewFixedPool(t *testing.T) {
	tests := []struct {
		name         string
		size         int
		policy       int
		hasErr       bool
		expectSize   int64
		expectPolicy int
	}{
		{
			name:         "阻塞策略",
			size:         1,
			policy:       BlockPolicy,
			hasErr:       false,
			expectSize:   1,
			expectPolicy: BlockPolicy,
		},
		{
			name:         "抛弃策略",
			size:         1,
			policy:       DiscardPolicy,
			hasErr:       false,
			expectSize:   1,
			expectPolicy: DiscardPolicy,
		},
		{
			name:   "非法策略",
			size:   1,
			policy: -1,
			hasErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool, err := NewFixedPool(tt.size, tt.policy)
			if tt.hasErr {
				if err == nil {
					t.Error("error 不应该为空")
					return
				}
			} else {
				if err != nil {
					t.Error(err)
					return
				}
				if pool.size != tt.expectSize || pool.policy != tt.expectPolicy {
					t.Errorf("got: size=%d,policy=%d expect: size=%d,policy=%d",
						pool.size, pool.policy, tt.expectSize, tt.expectPolicy)
				}
			}
		})
	}
}

func TestFixedPool_Execute(t *testing.T) {
	taskFn := func() {
		time.Sleep(time.Second)
	}
	tests := []struct {
		name   string
		size   int
		policy int
		tasks  int
		minSec int64
		maxSec int64
	}{
		{
			name:   "阻塞策略测试",
			size:   10,
			policy: BlockPolicy,
			tasks:  100,
			minSec: 9,
			maxSec: 10,
		},
		{
			name:   "丢弃策略测试",
			size:   10,
			policy: DiscardPolicy,
			tasks:  100,
			minSec: 0,
			maxSec: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool, err := NewFixedPool(tt.size, tt.policy)
			if err != nil {
				t.Error(err)
				return
			}
			t1 := time.Now().Unix()
			for i := 0; i < tt.tasks; i++ {
				pool.Execute(taskFn)
			}
			t2 := time.Now().Unix()
			res := t2 - t1
			if res < tt.minSec || res > tt.maxSec {
				t.Errorf("耗时 %d 秒，超过[%d,%d]时间限制", res, tt.minSec, tt.maxSec)
			}
		})
	}
}

// 测试 Execute 执行空函数效率(size=10000,policy=BlockPolicy)，测试结果：
// Mac mini M1 2020 16GB
// goos: darwin
// goarch: arm64
// pkg: github.com/tkgfan/got/core/concurrent
// BenchmarkFixedPool_Execute
// BenchmarkFixedPool_Execute-8   	 6328113	       179.0 ns/op
func BenchmarkFixedPool_Execute(b *testing.B) {
	pool, err := NewFixedPool(10000, BlockPolicy)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		pool.Execute(func() {

		})
	}
}
