type DSU struct {
	parent map[int]int
	rank map[int]int
}

func newDSU(elements []int) *DSU {
	dsu := &DSU{
		parent: make(map[int]int, len(elements)),
		rank: make(map[int]int, len(elements)),
	}
	for _, element := range elements {
		dsu.parent[element] = element
		dsu.rank[element] = 1
	}
	return dsu
}
func (dsu *DSU) Find(n int) int {
	if dsu.parent[n] != n {
		dsu.parent[n] = dsu.Find(dsu.parent[n])
	}
	return dsu.parent[n]
}
func (dsu *DSU) Union(nodeA, nodeB int) bool {
	rootA := dsu.Find(nodeA)
	rootB := dsu.Find(nodeB)
	if rootA == rootB {
		return false
	}
	if dsu.rank[rootA] >= dsu.rank[rootB] {
		dsu.parent[rootB] = rootA
		dsu.rank[rootA] += dsu.rank[rootB]
	} else {
		dsu.parent[rootA] = rootB
		dsu.rank[rootB] += dsu.rank[rootA]
	}
	return true
}

func findRedundantConnection(edges [][]int) []int {
    // DSU is a way to detect this
	elementsVisited := make(map[int]bool)
	elements := []int{}
	for _, edge := range edges {
		if !elementsVisited[edge[0]] {
			elements = append(elements, edge[0])
			elementsVisited[edge[0]] = true
		}
		if !elementsVisited[edge[1]] {
			elements = append(elements, edge[1])
			elementsVisited[edge[1]] = true
		}
	}
	dsu := newDSU(elements)
	for _, edge := range edges {
		if !dsu.Union(edge[0], edge[1]) {
			return []int{edge[0], edge[1]}
		}
	}
	return []int{}
}
