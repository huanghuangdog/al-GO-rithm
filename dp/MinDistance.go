package dp

func minDistance(word1 string, word2 string) int {
	size1 := len(word1)
	size2 := len(word2)
	dp := make([][]int, size1 + 1)
	for i := 0;i <= size1;i++ {
		dp[i] = make([]int, size2 + 1)
		dp[i][0] = i
	}
	for j := 0;j <= size2;j++ {
		dp[0][j] = j
	}

	for i := 1;i <= size1;i++ {
		for j:= 1;j <= size2;j++ {
			left := dp[i][j - 1] + 1
			up := dp[i - 1][j] + 1
			leftUp := dp[i - 1][j - 1]
			if word1[i - 1] != word2[j - 1] {
				leftUp = leftUp + 1
			}
			dp[i][j] = min(left, up, leftUp)
		}
	}
	return dp[size1][size2]
}

func min(a int, b int ,c int) int {
	rs := a
	if rs > b {
		rs = b
	}
	if rs > c {
		rs = c
	}
	return rs
}
