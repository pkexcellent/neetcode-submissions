//import "container/heap"
type Task struct {
	idx int
	start int
	processTime int
}

type minHeap []Task
func (m minHeap) Len() int {return len(m)}
func (m minHeap) Less(i, j int) bool {
	if m[i].processTime == m[j].processTime {
		return m[i].idx < m[j].idx
	}
	return m[i].processTime < m[j].processTime
}
func (m minHeap) Swap(i, j int) {m[i], m[j] = m[j], m[i]}
func (m *minHeap) Push(x any) {
	*m = append(*m, x.(Task))
}
func (m *minHeap) Pop() any {
	n := len(*m)
	last := (*m)[n-1]
	*m = (*m)[:n-1]
	return last
}

func getOrder(tasks [][]int) []int {
	// use sorted array and heap to determine which task to pick
	tks := make([]Task, len(tasks))
	for i, t := range tasks {
		tks[i] = Task{i, t[0], t[1]}
	}
	sort.Slice(tks, func(i, j int) bool {
		if tks[i].start == tks[j].start {
			return tks[i].processTime < tks[j].processTime
		}
		return tks[i].start < tks[j].start
	})
	h := &minHeap{}
	heap.Init(h)
	order := []Task{tks[0]}
	rs := []int{}
	i := 1
	curTime := 0
	for len(order) > 0 {
		cur := order[0]
		rs = append(rs, cur.idx)
		order = order[:len(order)-1]
		curTime = max(curTime, cur.start) + cur.processTime
		for i < len(tks) && tks[i].start <= curTime {
			heap.Push(h, tks[i])
			i++
		}
		if h.Len() > 0 {
			newTask := heap.Pop(h).(Task)
			order = append(order, newTask)
		} else if i < len(tks){
			order = append(order, tks[i])
			i++
		}
	}
	return rs
}
