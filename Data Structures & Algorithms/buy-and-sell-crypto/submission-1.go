func maxProfit(prices []int) int {
    /*
    var minLeft = make([]int, len(prices))
    var maxRight = make([]int, len(prices))
    minLeft[0] = prices[0]
    for i:=1; i<len(prices); i++ {
        minLeft[i] = min(prices[i], minLeft[i-1])
    }
    maxRight[len(prices)-1] = prices[len(prices)-1]
    for j:=len(prices)-2; j>=0; j-- {
        maxRight[j] = max(prices[j], maxRight[j+1])
    }
    var maxProfit = 0
    for i := 0; i< len(prices); i++ {
        maxProfit = max(maxProfit, maxRight[i] - minLeft[i])
    }
    return maxProfit
    */
    if len(prices) <=1 {
        return 0
    }
    var i, j = 0, 1
    var maxP = 0
    for (j < len(prices)) {
        if prices[i] < prices[j] {
            maxP = max(maxP, prices[j] - prices[i])
        } else {
            i = j
        }
        j++
    }
    return maxP
}
