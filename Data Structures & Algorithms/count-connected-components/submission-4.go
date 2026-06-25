type DSU struct {
	parent []int
	rank []int
}

func newDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i< n; i ++ {
		dsu.parent[i] = i
		dsu.rank[i] = 1 // rank is the size for future union weight
	}
	return dsu
}

func (dsu *DSU) Find(node int) int {
	if dsu.parent[node] != node {
		dsu.parent[node] = dsu.Find(dsu.parent[node])
	}
	return dsu.parent[node]
}
func (dsu *DSU) Union(nodeA, nodeB int) bool {
	rootA := dsu.Find(nodeA)
	rootB := dsu.Find(nodeB)
	if rootA == rootB {
		return false // already merged
	} else { // merge them together
		if dsu.rank[rootA] >= dsu.rank[rootB] {
			dsu.parent[rootB] = rootA
			dsu.rank[rootA] += dsu.rank[rootB]
		} else {
			dsu.parent[rootA] = nodeB
			dsu.rank[rootB] += dsu.rank[rootA]
		}
		return true
	}
}

func countComponents(n int, edges [][]int) int {
    // use disjoint set union - DSU
	dsu := newDSU(n)
	rs := n
	for _, edge := range edges {
		if dsu.Union(edge[0], edge[1]) {
			// merge two together via the edge
			rs--
		}
	}
	return rs

}
