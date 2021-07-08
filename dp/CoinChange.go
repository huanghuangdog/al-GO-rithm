package dp

import "math"

/**
step1
定义函数：f(i)为组成总金额为i所需的最少数量
step2
状态转移：f(i) = (for 0<j<n) min(f(i -coins[j]) + 1
step3
边界值
f(0) = 0
f(i)匹配不到的情况需要用大数表示，因为有效值是取min运算
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	var min func(int, int) int
	min = func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}

	size := len(coins)
	for i := 1; i <= amount; i++ {
		for j := 0; j < size; j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i-coins[j]], dp[i])
			}
		}
		dp[i] += 1
	}

	if dp[amount] >= math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
