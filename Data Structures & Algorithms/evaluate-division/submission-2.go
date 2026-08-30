type ele struct {
	num string
	result float64
}
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    // build adjancy map
	adj := make(map[string][]ele)
	for idx, equation := range equations {
		x, y := equation[0], equation[1]
		result := values[idx]
		adj[x] = append(adj[x], ele{y, result})
		adj[y] = append(adj[y], ele{x, 1.0/result})
	}

	// dfs
	var dfs func(x string, result float64, target string, visited map[string]bool) float64 
	dfs = func(x string, result float64, target string, visited map[string]bool) float64 {
		if x == target {
			return result
		}
		if visited[x] {
			return -1.0
		} else {
			visited[x] = true
		}
		if divs, exist := adj[x]; !exist {
			return -1.0
		} else {
			for _, div := range divs {
				rs := dfs(div.num, result*div.result, target, visited)
				if rs != -1.0 {
					return rs
				}
			}
			return -1.0
		}
	}

	rs := []float64{}
	for _, query := range queries {
		a, b := query[0], query[1]
		if _, exist := adj[a]; exist {
			if a == b {
				rs = append(rs, 1.0)
				continue
			}
			m := make(map[string]bool)
			curRs := dfs(a, 1.0, b, m)
			rs = append(rs, curRs)
		} else {
			rs = append(rs, -1.0)
		}
	}
	return rs
}
