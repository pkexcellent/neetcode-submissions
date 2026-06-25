func findRedundantConnection(edges [][]int) []int {
    // using DFS to detect cycle
	visited := make(map[int]bool)
	adj := make(map[int][]int)

	var dfs func(cur, parent int) bool
	dfs = func(cur, parent int) bool {
		if visited[cur] {
			return true
		}

		visited[cur] = true
		for _, node := range adj[cur] {
			if node == parent {
				continue // is the adj is parent, given we have marked it as true, so skip it
			} else {
				if dfs(node, cur) {
					return true
				}
			}
		}
		return false
	}
	for _, edge := range edges {
		// should add edge step by step, building the adjacent map
		src, dst := edge[0], edge[1]
		adj[src] = append(adj[src], dst)
		adj[dst] = append(adj[dst], src)
		// need to reset visited every loop
		for node, _ := range visited {
			visited[node] = false
		}
		if dfs(src, -1) {
			return []int{src, dst}
		}
	}
	return []int{}
}
