func findKthLargest(nums []int, k int) int {
	// method1
	//sort.Slice(nums, func(i, j int) bool {return nums[i] > nums[j]})
	//return nums[k-1]

	// method2
	/*
	k = len(nums)-k
	var quickSelect func(l, r, k int) int
	quickSelect = func(l, r, k int) int {
		pivot := nums[r]
		p := l
		for i := l; i < r; i++ {
			if nums[i] <= pivot {
				nums[i], nums[p] = nums[p], nums[i]
				p++
			}
		}
		nums[p], nums[r] = nums[r], nums[p]
		if p == k {
			return nums[p]
		} else if p < k {
			return quickSelect(p+1, r, k)
		} else {
			return quickSelect(l, p-1, k)
		}
	}
	return quickSelect(0, len(nums)-1, k)
	*/
	// method3
	mp := &minHeap{}
	heap.Init(mp)
	for _, num := range nums {
		if mp.Len() < k {
			heap.Push(mp, num)
		} else if (*mp)[0] < num {
			heap.Push(mp, num)
			heap.Pop(mp)
		}
	}
	return heap.Pop(mp).(int)
}

type minHeap []int
func (p minHeap) Len() int {return len(p)}
func (p minHeap) Less(i, j int) bool {return p[i] < p[j]}
func (p minHeap) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p *minHeap) Push(x any) {*p = append(*p, x.(int))}
func (p *minHeap) Pop() any {
	size := len(*p)
	rs := (*p)[size-1]
	*p = (*p)[0:size-1]
	return rs
}