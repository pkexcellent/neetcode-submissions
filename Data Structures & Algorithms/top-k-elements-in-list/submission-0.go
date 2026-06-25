type kv struct {
    k int
    v int
}

func topKFrequent(nums []int, k int) []int {
    // use a map to record the frequent
    // while at the same time, record the top K threthod
    var numFreqMap = map[int]int{}
    //var topKThreshold = k
    //var curTop = -100000
    for _, num := range nums {
        if _, exist := numFreqMap[num]; exist {
            numFreqMap[num] ++
        } else {
            numFreqMap[num] = 1
        }
    }
    var kvSlice = make([]kv, 0, len(numFreqMap))
    for k, v := range numFreqMap {
        kvSlice = append(kvSlice, kv{k, v})
    }
    sort.Slice(kvSlice, func (i, j int) bool {return kvSlice[i].v > kvSlice[j].v})
    var ans = make([]int, k)
    for i := 0; i< k; i++ {
        ans[i] = kvSlice[i].k
    }
    return ans
}
