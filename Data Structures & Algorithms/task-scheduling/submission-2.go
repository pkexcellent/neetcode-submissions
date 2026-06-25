func leastInterval(tasks []byte, n int) int {
	m := make(map[byte]int)
	maxTask := tasks[0]
	maxNum := -1
	for _, task := range tasks {
		m[task]++
		if m[task] >= maxNum {
			maxTask = task
			maxNum = m[task]
		}
	}
	slots := (maxNum - 1) * n
	for task, num := range m {
		if task == maxTask {continue}
		slots -= min(num, maxNum-1)
	}

	// slots is the final remaining idle
	// the total  steps will be idle+task_count
	return len(tasks) + max(0, slots)
	
}
