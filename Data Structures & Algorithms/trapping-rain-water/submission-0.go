func trap(height []int) int {
    var ans = 0
    for i, _ := range(height) {
        var leftMax = height[i]
        var rightMax = height[i]
        for j := 0; j< i; j++ {
            if height[j] > leftMax {
                leftMax = height[j]
            }
        }
        for j := len(height) -1; j>i; j-- {
            if height[j] > rightMax {
                rightMax = height[j]
            }
        } 
        ans += min(leftMax, rightMax) - height[i]
    }
    return ans
}

func min(i, j int) int {
    if i > j {
        return j
    } else {
        return i
    }
}