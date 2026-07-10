func numDistinct(s string, t string) int {
    // build a 2D metrix to show [i][j] i is the ith letter in t, j is the jth letter
    // value is 1 means hit
    dp := make([][]int, len(t))
    for i, _ := range dp {
        dp[i] = make([]int, len(s))
    }
    for i := 0; i < len(t); i++ {
        target := t[i]
        for j := 0; j < len(s); j++ {
            if s[j] == target {
                dp[i][j] = 1
            }
        }
    }
    fmt.Println(dp)
    for i := 1; i < len(t); i++ {
        for j := 0; j < len(s); j++ {
            if dp[i][j] != 0 {
                dp[i][j] = 0 // reassign it to 0, it should be decided by the previous row
                for k := 0; k < j; k++ {
                    dp[i][j] += dp[i-1][k]
                }
            }   
        }
    }
    fmt.Println(dp)
    rs := 0
    for i := 0; i < len(s); i++ {
        rs += dp[len(t)-1][i]
    }
    return rs
}
