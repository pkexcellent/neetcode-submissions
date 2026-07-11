func minDistance(word1 string, word2 string) int {
    // use a dfs way to do this
    dp := make([][]int, len(word1))
    for i, _ := range dp {
        dp[i] = make([]int, len(word2))
    }
    for i := 0; i < len(word1); i++ {
        for j := 0; j < len(word2); j++ {
            dp[i][j] = -1
        }
    }

    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i == len(word1) && j == len(word2) {
            return 0
        }
        if i >= len(word1) {
            return len(word2)-1 - j + 1
        }
        if j >= len(word2) {
            return len(word1)-1 - i + 1
        }
        if dp[i][j] != -1 {
            return dp[i][j]
        }
        change := 0
        if word1[i] == word2[j] {
            change += dfs(i+1, j+1)
        } else {
            change = 1 + min(min(dfs(i+1, j), dfs(i, j+1)), 
                    dfs(i+1, j+1)) // del i, insert j to i, or modify i
        }
        dp[i][j] = change
        return change
    }
    return dfs(0, 0)
}
