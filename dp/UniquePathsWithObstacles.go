package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	rs := make([][]int, m)

	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	for i := 0; i < m; i++ {
		rs[i] = make([]int, n)
		if i == 0 || (i > 0 && rs[i-1][0] != 0 && obstacleGrid[i][0] != 1) {
			rs[i][0] = 1
		}
	}

	for i := 1; i < n; i++ {
		if rs[0][i-1] != 0 && obstacleGrid[0][i] != 1 {
			rs[0][i] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				rs[i][j] = rs[i-1][j] + rs[i][j-1]
			}
		}
	}
	return rs[m-1][n-1]
}
