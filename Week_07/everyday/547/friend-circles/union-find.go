package friend_circles

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		parent: make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

func (u *UnionFind) Find(x int) int {
	p := u.parent
	for p[x] != x {
		p[x] = p[p[x]]
		x = p[x]
	}
	return x
}

func (u *UnionFind) Union(p, q int) {
	rootP, rootQ := u.Find(p), u.Find(q)
	if rootP == rootQ {
		return
	}
	u.parent[rootP] = rootQ
	u.count--
}

func (u *UnionFind) Count() int {
	return u.count
}

