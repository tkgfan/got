// author gmfan
// date 2024/7/15

package graph

import (
	"reflect"
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	uf := NewUnionFind(5)
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 0)
	uf.Union(3, 4)
	res := []int{2, 2, 2, 4, 4}
	if !reflect.DeepEqual(res, uf.parent) {
		t.Error("union find error")
	}
}
