func characterReplacement(s string, k int) int {
    var maxL = 0
    var l = 0
    var m = make(map[byte]int)
    var maxCharLength = 0
    for r := 0; r< len(s); r++ {
        m[s[r]] ++
        if m[s[r]] > maxCharLength {
            maxCharLength = m[s[r]]
        }
        
        if r-l+1 - maxCharLength > k {
            // 这步非常关键，在挪动左指针的时候要把最左边的count也-1
            m[s[l]]--
            l++
        } else {
            maxL = max(maxL, r-l+1)
        }
    }
    return maxL
}
