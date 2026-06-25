type flighInfo struct {
	cost int
	airport int
	stops int
}
type mheap []flighInfo
func (h mheap) Len() int {return len(h)}
func (h mheap) Less(i, j int) bool {return h[i].cost < h[j].cost}
func (h mheap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *mheap) Push(x any) {
	*h = append(*h, x.(flighInfo))
}
func (h *mheap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    // build adjacency table for flights
	adj := make(map[int][][2]int)
	for _, flight := range flights {
		from, airport, price := flight[0], flight[1], flight[2]
		adj[from] = append(adj[from], [2]int{airport, price})
	}

	// state table
	state := make([][]int, n)
	for i, _ := range state {
		state[i] = make([]int, k+2)
		for j, _ := range state[i] {
			state[i][j] = math.MaxInt
		}
	}
	
	h := &mheap{}
	heap.Init(h)
	h.Push(flighInfo{0, src, 0})
	for h.Len() > 0 {
		pickNextFlight := heap.Pop(h).(flighInfo)
		curCost, curAirport, curStops := pickNextFlight.cost, pickNextFlight.airport, pickNextFlight.stops
		if curAirport == dst {
			return curCost
		} 
		if curStops > k {
			continue
		}
		if curCost > state[curAirport][curStops] {
			continue
		}
		for _, adjTo := range adj[curAirport] {
			nextTo, price := adjTo[0], adjTo[1]
			newCost := curCost + price
			newStops := curStops + 1
			if newCost < state[nextTo][newStops] {
				state[nextTo][newStops] = newCost
				heap.Push(h, flighInfo{newCost, nextTo, newStops})
			}
		}
	}
	return -1
}
