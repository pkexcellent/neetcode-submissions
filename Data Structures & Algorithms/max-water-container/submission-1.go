func maxArea(heights []int) int {

    var i, j = 0, len(heights)-1
    var mostAns = 0
    fmt.Println(mostAns)
    for i < j {
        mostAns = max(mostAns, (j-i) * min(heights[i], heights[j]))
        if heights[i] < heights[j] {
            i++
        } else {
            j--
        }
        
    }
    return mostAns
}
