func orangesRotting(grid [][]int) int {
    totalFresh := 0
    q := [][2]int{}
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 2 {
                q = append(q, [2]int{i, j})
            }
            if grid[i][j] == 1 {
                totalFresh++
            }
        }
    }

    if totalFresh == 0 {
        return 0
    }

    directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    minutes := 0

    for len(q) > 0 && totalFresh > 0 {
        levelSize := len(q)  // snapshot: all currently rotten fruits
        for i := 0; i < levelSize; i++ {
            cur := q[0]
            q = q[1:]
            for _, dir := range directions {
                r, c := cur[0]+dir[0], cur[1]+dir[1]
                if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != 1 {
                    continue
                }
                grid[r][c] = 2
                totalFresh--
                q = append(q, [2]int{r, c})
            }
        }
        minutes++
    }

    if totalFresh > 0 {
        return -1
    }
    return minutes
}