type DSU struct {
	parent map[int]int
	rank map[int]int
}
func (dsu *DSU) Find(n int) int {
	if dsu.parent[n] != n {
		dsu.parent[n] = dsu.Find(dsu.parent[n])
	}
	return dsu.parent[n]
}

func (dsu *DSU) Union(nodeA, nodeB int) bool {
	// return true means they are previously not connected
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
func newDSU(n int) *DSU {
	dsu := &DSU{
		parent: make(map[int]int, n),
		rank: make(map[int]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.rank[i] = 1
	}
	return dsu
}

type Edge struct {
	src int // []int // use i, j to represent a node, make it easier
	dst int // []int
	dis int
}

func minCostConnectPoints(points [][]int) int {
    // use DSU to calculate the distance from smallest to larger
	// until all the nodes are connected to one root
	dsu := newDSU(len(points))
	// calculate edges between every two nodes
	edges := []Edge{}
	for i := 0; i < len(points); i++ {
		ax, ay := points[i][0], points[i][1]
		for j := i+1; j < len(points); j++ {
			bx, by := points[j][0], points[j][1]
			distance := int(math.Abs(float64(ax-bx)) + math.Abs(float64(ay-by)))
			edges = append(edges, Edge{
				src: i,
				dst: j,
				dis: distance,
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {return edges[i].dis < edges[j].dis})

	// put edges into DSU to merge
	totalDis := 0
	for _, edge := range edges {
		if dsu.Union(edge.src, edge.dst) {
			totalDis += edge.dis
		}
	}
	return totalDis

}
