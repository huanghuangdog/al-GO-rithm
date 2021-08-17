package dfs

/**
20210817
 */
func maxAreaOfIsland(grid [][]int) (rs int) {
	rowLen := len(grid)
	if rowLen == 0 {
		return 0
	}
	colLen := len(grid[0])

	var dfs func(row int, col int) int
	dfs = func(row int, col int) int {
		if row < 0 || row > rowLen || col < 0 || col > colLen {
			return 0
		}
		if grid[row][col] == 0 {
			return 0
		}
		grid[row][col] = 0
		return 1 + dfs(row - 1, col) + dfs(row + 1, col) + dfs(row, col - 1) + dfs(row, col + 1)
	}

	for i:= 0;i < rowLen;i++ {
		for j:= 0;j < colLen;j++ {
			if grid[i][j] == 1 {
				tmp := dfs(i,j)
				if tmp > rs {
					rs = tmp
				}
			}
		}
	}
	return
}


