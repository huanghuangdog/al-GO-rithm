package dp

func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = 1
	for i := 1; i < l; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + 1
	}
	rs := dp[0]
	for i := 1; i < l; i++ {
		if dp[i] > rs {
			rs = dp[i]
		}
	}
	return rs
}
