type Item struct {
	length int
	right  int
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	if h[i].length == h[j].length {
		return h[i].right < h[j].right
	}
	return h[i].length < h[j].length
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	type Query struct {
		value int
		index int
	}

	qs := make([]Query, len(queries))
	for i, q := range queries {
		qs[i] = Query{value: q, index: i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].value < qs[j].value
	})

	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}

	h := &MinHeap{}
	heap.Init(h)

	i := 0
	for _, q := range qs {
		// 加入所有左端点 <= q 的区间
		for i < len(intervals) && intervals[i][0] <= q.value {
			l, r := intervals[i][0], intervals[i][1]
			heap.Push(h, Item{
				length: r - l + 1,
				right:  r,
			})
			i++
		}

		// 弹出所有右端点 < q 的区间
		for h.Len() > 0 && (*h)[0].right < q.value {
			heap.Pop(h)
		}

		// 堆顶就是包含 q 的最短区间
		if h.Len() > 0 {
			ans[q.index] = (*h)[0].length
		}
	}

	return ans
}
