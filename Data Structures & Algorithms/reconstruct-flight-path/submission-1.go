func findItinerary(tickets [][]string) []string {
    // DFS, backtracking
	// build adj list and sort, make sure the smallest is in front
	adj := make(map[string][]string)
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		adj[src] = append(adj[src], dst)
	}
	for k, _ := range adj {
		sort.Strings(adj[k])
	}

	// dfs
	path := []string{"JFK"}
	var dfs func(airport string) bool
	dfs = func(airport string) bool {
		if len(path) == len(tickets) + 1 {
			return true
		}

		// pick a ticket and try to dfs
		dsts, exist := adj[airport]
		if !exist {return false}
		for i, dst := range dsts {
			adj[airport] = append(adj[airport][:i], adj[airport][i+1:]...) // pick the ith ticket
			path = append(path, dst)
			if dfs(dst) {return true}
			adj[airport] = append(adj[airport][:i], append([]string{dst}, adj[airport][i:]...)...)
			path = path[:len(path)-1]
		}
		return false
	}
	if dfs("JFK") {return path}
	return []string{}
}
