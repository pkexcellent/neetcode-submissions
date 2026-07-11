func isMatch(s string, p string) bool {
    // use dfs first, and use memo to record status
    memo := map[[2]int]bool{}
    /*
    memo := make([][]int, len(s))
    for i, _ := range memo {
        memo[i] = make([]int, len(p))
        for j := 0; j < len(p); j++ {
            memo[i][j] = -1
        }
    }
    */

    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        if i == len(s) {
            if j == len(p) || j+2 == len(p) && p[j+1] == byte('*') {
                return true
            } else {
                return false
            }
        } else if j == len(p) {
            return false
        }
        if v, exist := memo[[2]int{i, j}]; exist {
            return v
        }
        status := false

        if j+1 < len(p) && p[j+1] == byte('*') { // * should be together with preceding char 
            if p[j] == byte('.') || s[i] == p[j] {
                status = dfs(i+1, j+2) || // exact match once
                     dfs(i, j+2) || // match zero
                     dfs(i+1, j) // repeated match
            } else { // if not equal, only match zero is possible
                status = dfs(i, j+2)
            }
        } else if p[j] == byte('.') {
            status = dfs(i+1, j+1)
        } else if s[i] == p[j] { // this must be last, because s=abc, p=a*bc
            status = dfs(i+1, j+1)
        }  else {
            status = false
        }
        memo[[2]int{i, j}] = status
        return memo[[2]int{i, j}]
    }
    dfs(0, 0)
    return memo[[2]int{0, 0}]
}
