// still use minHeap
type Point struct {
	dist int
	point []int
}

type MaxHeap []Point

func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i, j int) bool {return h[i].dist > h[j].dist} 
func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x any) {*h = append(*h, x.(Point))}
func (h *MaxHeap) Pop() any {
	size := len(*h)
	ele := (*h)[size-1]
	*h = (*h)[0:size-1]
	return ele
}

func kClosest(points [][]int, k int) [][]int {
	var minPoints = &MaxHeap{}
	heap.Init(minPoints)
	for _, point := range points {
		dis := point[0]*point[0] + point[1]*point[1]
		heap.Push(minPoints, Point{
				dist: dis, 
				point: []int{point[0], point[1]},
			})
		if minPoints.Len() > k {
			heap.Pop(minPoints)
		}
	}
	rs := [][]int{}
	for minPoints.Len() > 0 {
		rs = append(rs, heap.Pop(minPoints).(Point).point)
	}
	return rs
}
