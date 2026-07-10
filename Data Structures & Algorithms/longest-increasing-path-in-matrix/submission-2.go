func longestIncreasingPath(matrix [][]int) int {
    // using BFS and topological sort
    // treat it as a graph with direction
    // direction is pointing a smaller number to a bigger
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    r, c := len(matrix), len(matrix[0])
    indegrees := make([][]int, r)
    for i, _ := range indegrees {
        indegrees[i] = make([]int, c)
    }
    directions := [][2]int{{1,0}, {0, 1}, {-1, 0}, {0, -1}}
    for i := 0; i < r; i++ {
        for j:= 0; j < c; j++ {
            for _, direction := range directions {
                nr, nc := i+direction[0], j+direction[1]
                if nr >= 0 && nr < r && nc >= 0 && nc < c && matrix[nr][nc] < matrix[i][j] {
                    indegrees[i][j]++
                }
            }
        }
    }
    q := [][2]int{}
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            if indegrees[i][j] == 0 {
                q = append(q, [2]int{i, j})
            }
        }
    }
    maxL := 0
    for len(q) > 0 {
        stepSize := len(q)
        for i := 0; i < stepSize; i++ {
            node := q[0]
            q = q[1:]
            for _, direction := range directions {
                nr, nc := node[0]+direction[0], node[1]+direction[1]
                if nr >= 0 && nr < r && nc >= 0 && nc < c && matrix[nr][nc] > matrix[node[0]][node[1]] {
                    indegrees[nr][nc]--
                    if indegrees[nr][nc] == 0 {
                        q = append(q, [2]int{nr, nc})
                    }
                }
            }
        }
        maxL++
    }
    return maxL

}
