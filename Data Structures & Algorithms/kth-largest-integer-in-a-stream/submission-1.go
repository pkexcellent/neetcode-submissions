// ok, directly use priority_queue
// import ("github.com/emirpasic/gods/queues/priorityqueue")
/*
type KthLargest struct {
    pq *priorityqueue.Queue
	k int
}


func Constructor(k int, nums []int) KthLargest {
	var pq = priorityqueue.NewWith(utils.IntComparator)
	for i, num := range nums {
		pq.Enqueue(num)
		if i >= k {
			pq.Dequeue()
		}
	}
	return KthLargest{
		pq: pq,
		k: k,
	}
}


func (this *KthLargest) Add(val int) int {
	this.pq.Enqueue(val)
	if this.pq.Size() > this.k {
		this.pq.Dequeue()
	}
	rs, _ := this.pq.Peek()
	return rs.(int) // Peek() returns interface and ok
}
*/
type KthLargest struct {
    k   int
    arr []int
}

func Constructor(k int, nums []int) KthLargest {
    return KthLargest{k: k, arr: nums}
}

func (this *KthLargest) Add(val int) int {
    this.arr = append(this.arr, val)
    sort.Ints(this.arr)
    return this.arr[len(this.arr)-this.k]
}
