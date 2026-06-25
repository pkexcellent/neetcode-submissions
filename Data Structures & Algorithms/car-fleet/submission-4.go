type carInfo struct {
	position int
	time float32
}
func carFleet(target int, position []int, speed []int) int {
	// calculate the time t of each car reaches to destination
	// as long as the previous t, is larger than the later one, they are in one fleet
	var fleets int = 0
	//var time = make([]int, len(speed))
	var cars = make([]carInfo, len(speed))
	for i := 0; i< len(position); i++ {
		t := float32(target - position[i]) / float32(speed[i])
		cars[i] = carInfo{position[i], t}
	}
	sort.Slice(cars, func(i, j int) bool {return cars[i].position < cars[j].position})
	
	fmt.Println(cars)
	// 接着就是跟上一题类似，找最右边第一个比它大的时间的idx，
	// 最后看有多少组被贴到0上的，就是被堵了多少组
	var fleetGroup = make([]int, len(cars))
	var stack = []int{}
	for i := 0; i < len(cars); i ++ {
		for len(stack) > 0 && cars[i].time >= cars[stack[len(stack)-1]].time {
			fleetGroup[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, idx := range stack {
		fleetGroup[idx] = 0
	}

	fmt.Println(fleetGroup)
	for _, nextGroup := range fleetGroup {
		if nextGroup == 0 {
			fleets++
		}
	}
	return fleets
}
