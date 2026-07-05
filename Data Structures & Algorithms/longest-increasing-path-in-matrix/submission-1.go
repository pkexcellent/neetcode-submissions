func longestIncreasingPath(matrix [][]int) int {
    rows, cols := len(matrix), len(matrix[0])
    dp := make([][]int, rows)
    for i := range dp {
        dp[i] = make([]int, cols)
    }

    var dfs func(r, c, prevVal int) int
    dfs = func(r, c, prevVal int) int {
        if r < 0 || r >= rows || c < 0 || c >= cols ||
           matrix[r][c] <= prevVal {
            return 0
        }
        if dp[r][c] != 0 {
            return dp[r][c]
        }

        res := 1
        res = max(res, 1 + dfs(r+1, c, matrix[r][c]))
        res = max(res, 1 + dfs(r-1, c, matrix[r][c]))
        res = max(res, 1 + dfs(r, c+1, matrix[r][c]))
        res = max(res, 1 + dfs(r, c-1, matrix[r][c]))
        dp[r][c] = res
        return res
    }

    maxPath := 0
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            maxPath = max(maxPath, dfs(r, c, -1))
        }
    }
    return maxPath
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}