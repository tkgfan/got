// author gmfan
// date 2024/3/28

package mongos

import (
	"context"
	"github.com/tkgfan/got/core/strs"
	"sync"
	"testing"
)

func TestAutoIncID(t *testing.T) {
	ctx := context.TODO()
	err := InitDatabase(ctx, Conf{
		URI: "mongodb://root:123456@192.168.1.15:27019/admin",
		DB:  "test_crud",
	})
	if err != nil {
		t.Error(err)
		return
	}

	table := "c_auto_inc"
	set := &sync.Map{}
	wg := &sync.WaitGroup{}
	for h := 0; h < 10; h++ {
		idKey := strs.Rand(10)
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func() {
				defer func() {
					wg.Done()
				}()
				for j := 0; j < 1000; j++ {
					id, err := AutoIncID(ctx, table, idKey)
					if err != nil {
						t.Error(err)
						return
					}
					set.Store(id, id)
				}
			}()
		}
	}

	wg.Wait()
	for i := 1; i <= 10000; i++ {
		if _, ok := set.Load(int64(i)); !ok {
			t.Errorf("缺少：%d", i)
			return
		}
	}
}
