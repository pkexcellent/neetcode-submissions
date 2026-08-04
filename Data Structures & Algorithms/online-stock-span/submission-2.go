type StockSpanner struct {
	// use monotonic stack, 单调递减, 
	// index in stack
	mstackday []int
	prices []int
	curDay int
}

func Constructor() StockSpanner {
	return StockSpanner {
		mstackday: []int{},
		prices: []int{},
		curDay: -1,
	}
}

func (this *StockSpanner) Next(price int) int {
	this.curDay++
	preDay := this.curDay - 1
	for len(this.mstackday) > 0 && price >= this.prices[this.mstackday[len(this.mstackday)-1]] {
		this.mstackday = this.mstackday[:len(this.mstackday)-1]
		preDay = -1
		if len(this.mstackday) > 0 {
			preDay = this.mstackday[len(this.mstackday)-1]	
		}
	}
	this.prices = append(this.prices, price)
	this.mstackday = append(this.mstackday, this.curDay)
	//fmt.Println(this.curDay, preDay, this.mstackday, this.prices)
	if len(this.mstackday) == 0 {
		return this.curDay + 1
	} else {
		return this.curDay - preDay
	}
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 