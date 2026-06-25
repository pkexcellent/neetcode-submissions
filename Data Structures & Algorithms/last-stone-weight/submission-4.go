//import ("container/heap")
// we need to use maxHeap
//  a way is to use priortyqueue, but importing a 3rd lib is not viable
// so use the standard queue to contruct one
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(val any)      { *h = append(*h, val.(int)) }
func (h *MaxHeap) Pop() any {
	size := len(*h)
	rs := (*h)[size-1]
	*h = (*h)[:size-1]
	return rs
}

func lastStoneWeight(stones []int) int {
	var maxHeap = &MaxHeap{}
	heap.Init(maxHeap)
	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}
	for maxHeap.Len() > 1 {
		large := heap.Pop(maxHeap).(int)
		secondLarge := heap.Pop(maxHeap).(int)
		if large == secondLarge {
			heap.Push(maxHeap, 0)
		} else if large > secondLarge {
			heap.Push(maxHeap, large-secondLarge)
		} else {
			heap.Push(maxHeap, secondLarge-large)
		}
	}
	return heap.Pop(maxHeap).(int)
}
