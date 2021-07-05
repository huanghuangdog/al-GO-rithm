package dp

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rs := make([][]int, m)

	for i := 0; i < m; i++ {
		rs[i] = make([]int, n)
		if i == 0 {
			rs[i][0] = grid[0][0]
		} else {
			rs[i][0] = grid[i][0] + rs[i-1][0]
		}
	}

	for i := 1; i < n; i++ {
		rs[0][i] = rs[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if rs[i][j-1] <= rs[i-1][j] {
				rs[i][j] = grid[i][j] + rs[i][j-1]
			} else {
				rs[i][j] = grid[i][j] + rs[i-1][j]
			}
		}
	}
	return rs[m-1][n-1]
}
