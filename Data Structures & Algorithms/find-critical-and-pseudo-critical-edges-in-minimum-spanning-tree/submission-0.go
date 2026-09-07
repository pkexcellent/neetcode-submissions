type UnionFind struct {
	parent []int
	rank []int
}
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}
func (u *UnionFind) Find(n int) int {
	if u.parent[n] == n {
		return n
	} else {
		u.parent[n] = u.Find(u.parent[n])
	}
	return u.parent[n]
}
func (u *UnionFind) Union(a, b int) bool {
	pa, pb := u.Find(a), u.Find(b)
	if pa == pb {
		return false
	}
	if u.rank[pa] >= u.rank[pb] {
		u.rank[pa] += u.rank[pb]
		u.parent[pb] = pa
	} else {
		u.rank[pb] += u.rank[pa]
		u.parent[pa] = pb
	}
	return true
}
func (u *UnionFind) Connected(totalNode int) bool {
	for _, r := range u.rank {
		if r == totalNode {
			return true
		}
	}
	return false
}
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// this is hard
	// firstly get the min weight
	msWeight := 0
	edgeList := [][4]int{}
	for i, edge := range edges {
		edgeList = append(edgeList, [4]int{edge[0], edge[1], edge[2], i})
	}
	sort.Slice(edgeList, func(i, j int) bool {return edgeList[i][2] < edgeList[j][2]})
	uf := NewUnionFind(n)
	for _, e := range edgeList {
		if uf.Union(e[0], e[1]) {
			msWeight += e[2]
		}
	}

	// check which one is critical and which one isn't
	critical, notCritical := []int{}, []int{}
	for _, excludeEdge := range edgeList {
		// try exclude this edge
		uft := NewUnionFind(n)
		weight := 0
		for _, includeEdge := range edgeList {
			if includeEdge[3] == excludeEdge[3] {
				continue
			}
			if uft.Union(includeEdge[0], includeEdge[1]) {
				weight += includeEdge[2]
			}
		}
		if !uft.Connected(n) || weight > msWeight {
			critical = append(critical, excludeEdge[3])
			continue // if critical, then not possible to be notCritical
		}
		// if not critical, need to check if this edge is in any MST
		// if it can be in any MST, then it's a non-critical edge
		// if it's in any MST, adding it forcely, won't increase msWeight
		ufn := NewUnionFind(n)
		fweight := excludeEdge[2]
		ufn.Union(excludeEdge[0], excludeEdge[1])
		for _, e := range edgeList {
			if e[3] == excludeEdge[3] {
				continue
			}
			if ufn.Union(e[0], e[1]) {
				fweight += e[2]
			}
		}
		if fweight == msWeight {
			notCritical = append(notCritical, excludeEdge[3])
		}
	}
	return [][]int{critical, notCritical}

}
