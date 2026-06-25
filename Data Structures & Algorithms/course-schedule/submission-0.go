func canFinish(numCourses int, prerequisites [][]int) bool {
    // this is very like how Airflow schedule tasks
	// use indegree to indicate which task can be scheduled

	var cursesIndegree = make([]int, numCourses)
	var downstream = make(map[int][]int, numCourses)

	for _, prerequiste := range prerequisites {
		cursesIndegree[prerequiste[0]]++
		downstream[prerequiste[1]] = append(downstream[prerequiste[1]], prerequiste[0])
	}

	q := []int{}
	for i, indegree := range cursesIndegree {
		if indegree == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		course := q[0]
		q = q[1:]
		for _, c := range downstream[course] {
			cursesIndegree[c]--
			if cursesIndegree[c] == 0 {
				q = append(q, c)
			}
		}
	}
	for _, indegree := range cursesIndegree {
		if indegree > 0 {return false}
	}
	return true
}
