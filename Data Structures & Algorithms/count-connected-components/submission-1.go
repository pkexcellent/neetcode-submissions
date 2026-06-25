func countComponents(n int, edges [][]int) int {
    // the best method is to use DUS, disjoint union set
	// but we can do a normal DFS method first

	// build adjacency table
	adj := make(map[int][]int)
	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		adj[src] = append(adj[src], dst)
		adj[dst] = append(adj[dst], src)
	}
	// record visited indo
	visited := make(map[int]bool)

	// dfs is to search all the nodes adjacent to a given node, make them as one component
	var dfs func(node int) 
	dfs = func(node int) {
		for _, adjNode := range adj[node] {
			if visited[adjNode] {
				continue
			} else {
				visited[adjNode] = true
				dfs(adjNode)
			}
		}
	}
	components := 0
	for i := 0; i < n; i++ {
		if visited[i] == false {
			components++
			visited[i] = true
			dfs(i)
		} 
	}
	return components
}
