type Meeting struct {
	start int
	end int
	room int
}

type MinHeap []Meeting
func (mh MinHeap) Len() int {return len(mh)}
func (mh MinHeap) Less(i, j int) bool {
	if mh[i].end == mh[j].end {
		return mh[i].room < mh[j].room
	}
	return mh[i].end < mh[j].end
}
func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh MinHeap) Peak() Meeting {
	return mh[0]
}
func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(Meeting))
}
func (mh *MinHeap) Pop() any {
	n := len(*mh)
	last := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return last
}

func mostBooked(n int, meetings [][]int) int {
	// sort meeting with start
	sort.Slice(meetings, func(i, j int) bool {return meetings[i][0] < meetings[j][0]})
	// use heap to order the meeting
	cap := n
	hp := &MinHeap{}
	heap.Init(hp)
	room := 1
	roomCnt := make(map[int]int)
	for cap > 0 {
		heap.Push(hp, Meeting{-1, -1, room})
		room++	
		cap--
	}

	for i := 0; i < len(meetings); i++ {
		thisMeet := meetings[i]
		// free all the rooms that ends before thisMeet starts
		for hp.Len() > 0 && hp.Peak().end < thisMeet[0] {
			freeable := heap.Pop(hp).(Meeting)
			// use start as Meeting's end
			heap.Push(hp, Meeting{thisMeet[0], thisMeet[0], freeable.room})
		}
		endMeeting := heap.Pop(hp).(Meeting)
		heap.Push(hp, Meeting{
			max(endMeeting.end, thisMeet[0]), 
			max(endMeeting.end, thisMeet[0]) + (thisMeet[1] - thisMeet[0]), 
			endMeeting.room,
		})
		roomCnt[endMeeting.room]++
	}
	fmt.Println(roomCnt)
	room, mostMeetings := n+1, -1
	for r, cnt := range roomCnt {
		if cnt > mostMeetings { // golang map has no order, so need to compare == case
			mostMeetings = cnt
			room = r
		} else if cnt == mostMeetings {
			if r < room {
				room = r
			}
		}
	}
	return room-1 // room idx starts from 0
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
