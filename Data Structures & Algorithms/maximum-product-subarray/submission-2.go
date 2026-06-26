func maxProduct(nums []int) int {
    // this is a very very smart algorithm
	// record two states actually
	curMax, curMin := 1, 1
	allMax := math.MinInt

	for _, num := range nums {
		tempMax := curMax*num
		tempMin := curMin*num
		curMax = max(tempMax, max(tempMin, num))
		curMin = min(tempMax, min(tempMin, num))
		allMax = max(allMax, curMax)
	}
	return allMax

}
