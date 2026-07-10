func numDistinct(s string, t string) int {
    // build a 2D metrix to show [i][j] i is the ith letter in t, j is the jth letter
    // value is 1 means hit
    r, c := len(t), len(s)
    dp := make([][]int, r)
    for i, _ := range dp {
        dp[i] = make([]int, c)
    }
    for i := 0; i < r; i++ {
        target := t[i]
        for j := 0; j < c; j++ {
            if s[j] == target {
                dp[i][j] = 1
            }
        }
    }
    for i := 1; i < r; i++ {
        for j := 0; j < c; j++ {
            if dp[i][j] != 0 {
                dp[i][j] = 0 // reassign it to 0, it should be decided by the previous row
                for k := 0; k < j; k++ {
                    dp[i][j] += dp[i-1][k]
                }
            }   
        }
    }
    rs := 0
    for i := 0; i < c; i++ {
        rs += dp[r-1][i]
    }
    return rs
}
