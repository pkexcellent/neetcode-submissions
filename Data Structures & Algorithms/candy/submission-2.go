func candy(ratings []int) int {
	// find monotonic patterns
	// 100, 5, 4, 3, 5, 6, 5, 4, 100
	// 1,   0, 0, 0, 1, 1, 0, 0, 1
	// 1,   0, 0, 1, 2, 3, 0, 0, 1
	// 1,   3, 2, 1, 2, 3, 2, 1, 1
	n := len(ratings) + 2
	rates := append(append([]int{math.MaxInt32}, ratings...), math.MaxInt32)
	candy := make([]int, n)
	flag := make([]int, n)
	flag[0] = 1
	candy[0] = 1
	for i := 1; i < n; i++ {
		if rates[i] > rates[i-1] {
			flag[i] = 1
		} else if rates[i] < rates[i-1] {
			flag[i] = -1
		} else {
			flag[i] = 0
		}
	}
	fmt.Println(rates)
	fmt.Println("l->r", flag)
	for i := 1; i < n; i++ {
		if flag[i] <= 0 && flag[i+1] == 1 {
			candy[i] = 1
		} else if flag[i] == 1 {
			candy[i] = candy[i-1] + 1
		}
	}
	fmt.Println(candy)
	for i := n-2; i > 0; i-- {
		if rates[i] > rates[i+1] {
			flag[i] = 1
		} else if rates[i] < rates[i+1] {
			flag[i] = -1
		} else {
			flag[i] = 0
		}
	}
fmt.Println("r->l", flag)
	for i := n-2; i > 0; i-- {
		if flag[i-1] == 1 && flag[i] <= 0 {
			candy[i] = max(candy[i], 1)
		} else if flag[i] == 1 {
			candy[i] = max(candy[i], candy[i+1]+1)
		}
	}
	
	
	fmt.Println(candy)
	rs := 0
	for i := 1; i <= n-2; i++ {
		if candy[i] == 0 {
			candy[i] = 1
		}
		rs += candy[i]
	}
	return rs
}
