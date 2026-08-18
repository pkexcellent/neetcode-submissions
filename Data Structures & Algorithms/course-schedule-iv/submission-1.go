func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// use topology sort
	indegrees := make([]int, numCourses)
	adj := make(map[int][]int)
	for _, prerequiste := range prerequisites {
		from, to := prerequiste[0], prerequiste[1]
		adj[from] = append(adj[from], to)
		indegrees[to]++
	}
	//fmt.Println(indegrees)
	// using a full map will be out of memory, prereqOf := make(map[int][]int)
	// change it to a map of map
	prereqOf := make(map[int]map[int]bool)
	q := []int{}
	for i, indegree := range indegrees {
		if indegree == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		take := q[0]
		q = q[1:]
		if viables, exist := adj[take]; exist {
			for _, to := range viables {
				indegrees[to]--
				//prereqOf[to] = append(prereqOf[to], append(prereqOf[take], take)...)
				if _, exist := prereqOf[to]; !exist {
					prereqOf[to] = make(map[int]bool)
				}
				prereqOf[to][take] = true
				for upstream, _ := range prereqOf[take] {
					prereqOf[to][upstream] = true
				}
				if indegrees[to] == 0 {
					q = append(q, to)
				}
			}
		}
	}
	//fmt.Println(prereqOf)
	rs := []bool{}
	for _, query := range queries {
		from, to := query[0], query[1]
		if _, exist := prereqOf[to][from]; exist {
			rs = append(rs, true)
		} else {
			rs = append(rs, false)
		}
	}
	return rs
}
