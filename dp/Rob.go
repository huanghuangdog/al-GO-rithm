package dp

func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	dp := make([]int, l)
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[0] + nums[2]
	max := 0
	for i := 3; i < l; i++ {
		if dp[i-2] > dp[i-3] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-3] + nums[i]
		}
	}

	for i := 0; i < l; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
