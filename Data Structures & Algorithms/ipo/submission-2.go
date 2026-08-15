type MaxHeap []Task
func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i, j int) bool {
	return h[i].profit > h[j].profit
}
func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x any) {*h = append(*h, x.(Task))}
func (h *MaxHeap) Pop() any {
	n := len(*h)
	last := (*h)[n-1]
	*h = (*h)[:n-1]
	return last
}
func (h *MaxHeap) Peak() any {
	return (*h)[len(*h)-1]
}

type Task struct {
	profit int
	cap int
}
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// object task {profit, capitcal}, order by captical, 2nd profit DESC
	// with the inital w, push all the actionable tasks in to maxHeap
	// maxheap order by profit, while picking from maxHeap with maxprofile, 
	// pick more actionable tasks into maxheap
	tasks := make([]Task, len(profits))
	for i, profit := range profits {
		tasks[i] = Task{
			profit: profit,
			cap: capital[i],
		}
	}
	sort.Slice(tasks, func(i, j int) bool {return tasks[i].cap < tasks[j].cap})
	//fmt.Println(tasks)
	h := &MaxHeap{}
	heap.Init(h)
	noww := w
	i := 0
	pickedTasks := 0
	for i < len(tasks) || h.Len() > 0 {
		for i < len(tasks) && tasks[i].cap <= noww {
			heap.Push(h, tasks[i])
			i++
		}
		//fmt.Println(h)

		if h.Len() == 0 {
			break
		}
		pick := heap.Pop(h).(Task)
		noww = noww + pick.profit
		pickedTasks++
		
		if pickedTasks >= k {
			break
		}
	}
	return noww
}
