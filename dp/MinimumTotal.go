package dp

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	l := len(triangle)
	dp := make([][]int, l)
	dp[0] = make([]int, 1)
	dp[0][0] = triangle[0][0]
	for m := 1; m < l; m++ {
		dp[m] = make([]int, m+1)
		for i := 0; i < m+1; i++ {
			if i == 0 {
				dp[m][i] = dp[m-1][i] + triangle[m][i]
			} else if i == m {
				dp[m][i] = dp[m-1][i-1] + triangle[m][i]
			} else {
				if dp[m-1][i-1] > dp[m-1][i] {
					dp[m][i] = dp[m-1][i] + triangle[m][i]
				} else {
					dp[m][i] = dp[m-1][i-1] + triangle[m][i]
				}
			}
		}
	}

	size := len(dp[l-1])
	min := math.MaxInt32
	for i := 0; i < size; i++ {
		if dp[l-1][i] < min {
			min = dp[l-1][i]
		}
	}
	return min
}
