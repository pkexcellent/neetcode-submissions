type UnionFind struct {
	parent []int
	rank []int
}
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind {
		parent: make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}
func (uf *UnionFind) Find(n int) int {
	if uf.parent[n] != n {
		uf.parent[n] = uf.Find(uf.parent[n])
	} 
	return uf.parent[n]
}
func (uf *UnionFind) Union(a, b int) bool {
	pa, pb := uf.Find(a), uf.Find(b)
	if pa == pb {return false}
	if uf.rank[pa] >= uf.rank[pb] {
		uf.parent[pb] = pa
		uf.rank[pa] += uf.rank[pb]
	} else {
		uf.parent[pa] = pb
		uf.rank[pb] += uf.rank[pa]
	}
	return true
}

func (uf *UnionFind) ConnectedReady(n int) bool {
	for i := 0; i < n; i++ {
		if uf.rank[i] == n {
			return true
		}
	}
	return false
}

func canTraverseAllPairs(nums []int) bool {
	// use an edge to represent num[i] and num[j] are passable when gcd > 1
	// -- then use topology sort to check whether the graph can be built
	// then use unionFind to check if using the edges can form a connected component
	n := len(nums)
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if gcd(nums[i], nums[j]) > 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.ConnectedReady(n)
}

func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
