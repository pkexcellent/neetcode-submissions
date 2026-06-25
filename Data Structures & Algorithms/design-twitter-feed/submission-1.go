type Twitter struct {
    count int
	tMap map[int][][]int // user_id -> list of tweet [count, tid]
	followMap map[int]map[int]bool // user_id -> followedID
}


func Constructor() Twitter {
    return Twitter {
		count: 0,
		tMap: make(map[int][][]int),
		followMap: make(map[int]map[int]bool),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    if this.tMap[userId] == nil {
		this.tMap[userId] = make([][]int, 0)
	} 
	this.tMap[userId] = append(this.tMap[userId], []int{this.count, tweetId})
	this.count--
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    mh := &minHeap{}
	heap.Init(mh)
	// push each one's latest tw into minHeap, including themselves
	if this.followMap[userId] == nil {
        this.followMap[userId] = make(map[int]bool)
    }
	this.followMap[userId][userId] = true
	for followee := range this.followMap[userId] {
		tInfos := this.tMap[followee]
		if len(tInfos) > 0 {
			curIdx := len(tInfos) - 1
			heap.Push(mh, []int{tInfos[curIdx][0], tInfos[curIdx][1], followee, curIdx-1})
		}
	}
	rs := []int{}
	for mh.Len() > 0 && len(rs) < 10 {
		// pick the latest one from heap top
		latest := heap.Pop(mh).([]int)
		userId := latest[2]
		//count := latest[0]
		tid := latest[1]
		nextIdx := latest[3]
		rs = append(rs, tid)
		// don't forget to check whether this one has more tids
		if nextIdx >= 0 {
			heap.Push(mh, []int{this.tMap[userId][nextIdx][0], this.tMap[userId][nextIdx][1], userId, nextIdx-1})
		}
	}
	return rs
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.followMap[followerId] == nil {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true 
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if this.followMap[followerId] != nil {
		delete(this.followMap[followerId], followeeId)
	}
}

type minHeap [][]int
func(h minHeap) Len() int {return len(h)}
func(h minHeap) Less(i, j int) bool {return h[i][0] < h[j][0]}
func(h minHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func(h *minHeap) Push(x any) {
	*h = append(*h, x.([]int))
}
func(h *minHeap) Pop() any {
	size := len(*h)
	top := (*h)[size-1]
	*h = (*h)[0:size-1]
	return top
}
