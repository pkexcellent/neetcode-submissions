func validTree(n int, edges [][]int) bool {
    // there are two conditions
	// 1. edges must be n-1
	// 2. there is no loop
	if len(edges) != n-1 {return false}
	var adj = make(map[int][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	
	var visited = make(map[int]bool)
	var dfs func(int, int) bool
	dfs = func(node int, parent int) bool {
		if visited[node] {return false}
		visited[node] = true
		for _, nAdj := range adj[node] {
			if nAdj == parent {
				continue
			}
			if !dfs(nAdj, node) {
				return false
			}
		}
		return true
	}
	if dfs(0, -1) && len(visited) == n {
		return true
	}
	return false
}
