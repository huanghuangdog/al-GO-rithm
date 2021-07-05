package dp

/**
动态规划
lc 62
*/
func uniquePaths(m int, n int) int {
	/**
	初始化
	*/
	rs := make([][]int, m)
	for i := 0; i < m; i++ {
		rs[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		rs[0][i] = 1
	}
	for i := 0; i < m; i++ {
		rs[i][0] = 1
	}

	/**
	rs[i][j]表示到达第i行第j列的可能性种数
	递归公式 rs[i][j] = rs[i][j-1] + rs[i-1][j]
	*/
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rs[i][j] = rs[i][j-1] + rs[i-1][j]
		}
	}
	return rs[m-1][n-1]
}
