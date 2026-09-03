type Edge struct {
	diff int
	row int
	col int
}
type MinHeap []Edge
func (mh MinHeap) Len() int {return len(mh)}
func (mh MinHeap) Less(i, j int) bool {return mh[i].diff < mh[j].diff}
func (mh MinHeap) Swap(i, j int) {mh[i], mh[j] = mh[j], mh[i]}
func (mh *MinHeap) Push(x any) {*mh = append(*mh, x.(Edge))}
func (mh *MinHeap) Pop() any {
	n := len(*mh)
	last := (*mh)[n-1]
	(*mh) = (*mh)[:n-1]
	return last
}
func minimumEffortPath(heights [][]int) int {
	// djstra
	m, n := len(heights), len(heights[0])
	dist := make([][]int, m)
	for i, _ := range dist {
		dist[i] = make([]int, n)
		for j := 0; j< n; j++ {
			dist[i][j] = math.MaxInt64
		}
	}
	dist[0][0] = 0
	q := &MinHeap{}
	heap.Init(q)
	heap.Push(q, Edge{0, 0, 0})
	for q.Len() > 0 {
		top := heap.Pop(q).(Edge)
		if top.row == m-1 && top.col == n-1 {
			return top.diff
		}
		if top.diff > dist[top.row][top.col] {
			continue // no need to update
		}
		directions := [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
		for _, dir := range directions {
			newRow, newCol := top.row + dir[0], top.col + dir[1]
			if newRow >= 0 && newCol >= 0 && newRow < m && newCol < n {
				maxToThisPos := max(top.diff, abs(heights[newRow][newCol] - heights[top.row][top.col]))
				if maxToThisPos < dist[newRow][newCol] {
					dist[newRow][newCol] = maxToThisPos
					heap.Push(q, Edge{maxToThisPos, newRow, newCol})
				}
			}
		}
	}
	return -1
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
