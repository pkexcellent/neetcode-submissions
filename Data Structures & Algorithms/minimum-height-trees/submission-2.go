func findMinHeightTrees(n int, edges [][]int) []int {
	// like unpeering onion
	// all the edges contruct a graph, only when picking the most central one(s)
	// can make the tree's height shortest
	if len(edges) == 0 {
		return []int{0}
	}
	adj := make(map[int][]int)
	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		/*if _, exist := adj[src]; !exist {
			adj[src] = []int{}
		}
		if _, exist := adj[dst]; !exist {
			adj[dst] = []int{}
		}
		*/
		adj[src] = append(adj[src], dst)
		adj[dst] = append(adj[dst], src)
	}
	
	indegree := make(map[int]int) // indegree is at least 1, becasue dual direction
	q := []int{}
	for src, dsts := range adj {
		indegree[src] = len(dsts)
		if len(dsts) == 1 {
			q = append(q, src)
		}
	}
	fmt.Println(q)
	remainV := len(adj)
	for len(q) > 0 {
		if remainV <= 2 {
			return q
		}
		size := len(q)
		remainV -= size
		for size > 0 {
			node := q[0]
			q = q[1:]
			for _, dst := range adj[node] {
				indegree[dst]--
				if indegree[dst] == 1 {
					q = append(q, dst)
				}
			}
			size--
		}
	}
	return q

}
