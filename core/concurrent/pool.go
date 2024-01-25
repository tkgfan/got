// author gmfan
// date 2023/12/4

package concurrent

import (
	"github.com/tkgfan/got/core/errors"
	"sync/atomic"
	"time"
)

const (
	// UseDefaultPolicy 使用默认策略
	UseDefaultPolicy = 0
	// BlockPolicy 阻塞策略
	BlockPolicy = 1 << 0
	// DiscardPolicy 丢弃策略
	DiscardPolicy = 1 << 1
)

// 策略检查，非法策略返回 false，反之 true
func policyCheck(policy int) bool {
	switch policy {
	case UseDefaultPolicy:
		return true
	case BlockPolicy:
		return true
	case DiscardPolicy:
		return true
	}
	return false
}

type TaskPool interface {
	// Submit 提交任务使用默认 policy
	Submit(fn func())
	// SubmitUsePolicy 提交任务并能使用自定义 policy
	SubmitUsePolicy(fn func(), policy int)
	// GetDefaultPolicy 获取默认策略
	GetDefaultPolicy() int
}

type fixedPool struct {
	// 池大小
	size          int64
	curSize       *int64
	defaultPolicy int
}

// NewFixedPool 创建固定大小池，defaultPolicy 为对应策略
func NewFixedPool(size, defaultPolicy int) (*fixedPool, error) {
	cSize := int64(0)
	pool := &fixedPool{
		size:    int64(size),
		curSize: &cSize,
	}
	err := pool.setPolicy(defaultPolicy)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func (f *fixedPool) Submit(fn func()) {
	f.SubmitUsePolicy(fn, f.defaultPolicy)
}

func (f *fixedPool) SubmitUsePolicy(fn func(), policy int) {
	if policy == UseDefaultPolicy {
		policy = f.defaultPolicy
	}
	nextSize := atomic.AddInt64(f.curSize, 1)
	if nextSize <= f.size {
		// 直接执行
		f.worker(fn)
		return
	}

	// 没有成功加入任务需要移除加入操作
	atomic.AddInt64(f.curSize, -1)

	// 达到并发上线，按照具体策略来处理任务
	switch policy {
	case BlockPolicy:
		f.blockHandle(fn)
	case DiscardPolicy:
		return
	}
}

func (f *fixedPool) blockHandle(fn func()) {
	nextSize := atomic.AddInt64(f.curSize, 1)
	// 空转等待
	for nextSize > f.size {
		atomic.AddInt64(f.curSize, -1)
		// 休眠 10 毫秒
		time.Sleep(10 * time.Microsecond)
		nextSize = atomic.AddInt64(f.curSize, 1)
	}

	// 获取工作协程
	f.worker(fn)
}

func (f *fixedPool) setPolicy(policy int) error {
	if policyCheck(policy) {
		f.defaultPolicy = policy
	} else {
		return errors.New("不支持此策略")
	}
	return nil
}

func (f *fixedPool) GetDefaultPolicy() int {
	return f.defaultPolicy
}

func (f *fixedPool) worker(fn func()) {
	go func() {
		defer func() {
			// 任务结束，curSize -1
			atomic.AddInt64(f.curSize, -1)
		}()
		fn()
	}()
}
