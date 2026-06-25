func findOrder(numCourses int, prerequisites [][]int) []int {
    // Airflow!
	var indegrees = make([]int, numCourses)
	var downstreams = make(map[int][]int, numCourses)
	
	for _, prerequisite := range prerequisites {
		down, up := prerequisite[0], prerequisite[1]
		downstreams[up] = append(downstreams[up], down)
		indegrees[down]++
	}

	schedulable := []int{}
	q := []int{}
	for i, degree := range indegrees {
		if degree == 0 {
			q = append(q, i)
			schedulable = append(schedulable, i)
		}
	}
	for len(q) > 0 {
		course := q[0]
		q = q[1:]
		for _, down := range downstreams[course] {
			indegrees[down]--
			if indegrees[down] == 0 {
				q = append(q, down)
				schedulable = append(schedulable, down)
			}
		}
	}
	for _, c := range indegrees {
		if c > 0 {
			return []int{}
		}
	}
	return schedulable
}
