type Event struct {
	pos int
	num int
}
func carPooling(trips [][]int, capacity int) bool {
	// interval? 
	events := make([]Event, 0, len(trips)*2)
	for _, trip := range trips {
		pass, start, end := trip[0], trip[1], trip[2]
		events = append(events, Event {
			pos: start,
			num: pass,
		})
		events = append(events, Event {
			pos: end,
			num: -pass,
		})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].pos == events[j].pos {
			return events[i].num < events[j].num // first out, then in
		}
		return events[i].pos < events[j].pos
	})
	fmt.Println(events)
	cap := capacity
	for _, event := range events {
		cap = cap - event.num
		if cap < 0 {
			return false
		}
	}
	return true
}
