// description 并查集
// author gmfan
// date 2024/7/15

package graph

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent: parent}
}

func (u *UnionFind) Find(idx int) int {
	if u.parent[idx] != idx {
		u.parent[idx] = u.Find(u.parent[idx])
	}
	return u.parent[idx]
}

func (u *UnionFind) Union(idx1, idx2 int) {
	u.parent[u.Find(idx1)] = u.Find(idx2)
}
