func characterReplacement(s string, k int) int {
    var maxL = 0
    for i := 0; i< len(s); i++ {
        var m = make(map[byte]int)
        var maxNumCnt = 0
        for j := i; j< len(s); j++ {
            m[s[j]] ++
            maxNumCnt = max(maxNumCnt, m[s[j]])
            if (j-i+1) - maxNumCnt <= k { // K个数字足够填充，不是最多number的其他number的话
                maxL = max(maxL, j-i+1)
            }
            
        }
    }
    return maxL
}
