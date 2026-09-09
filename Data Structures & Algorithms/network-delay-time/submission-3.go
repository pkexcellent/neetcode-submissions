type ToNode struct {
	node int
	weight int
}
type MinHeap []ToNode
func(m MinHeap) Len() int {return len(m)}
func(m MinHeap) Less(i, j int) bool {return m[i].weight < m[j].weight}
func(m MinHeap) Swap(i, j int) {m[i], m[j] = m[j], m[i]}
func(m *MinHeap) Push(x any) {*m = append(*m, x.(ToNode))}
func(m *MinHeap) Pop() any {
	size := len(*m)
	last := (*m)[size-1]
	*m = (*m)[:size-1]
	return last
}
func networkDelayTime(times [][]int, n int, k int) int {
    // dijkstra
	adj := make(map[int][]ToNode)
	for _, time := range times {
		src, dst, weight := time[0], time[1], time[2]
		adj[src] = append(adj[src], ToNode{dst, weight})
	}
	visited := make(map[int]bool)
	mh := &MinHeap{}
	heap.Init(mh)
	curTime := 0
	heap.Push(mh, ToNode{k, curTime})
	//fmt.Println(adj)
	for mh.Len() > 0 {
		curNode := heap.Pop(mh).(ToNode)
		if visited[curNode.node] {
			continue
		}
		curTime = curNode.weight
		visited[curNode.node] = true
		for _, nei := range adj[curNode.node] {
			if !visited[nei.node] {
				heap.Push(mh, ToNode{nei.node, curTime + nei.weight})
			}
		}
	}
	//fmt.Println(visited)
	if len(visited) == n {
		return curTime
	}
	return -1
}
