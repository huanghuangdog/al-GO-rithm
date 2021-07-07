package dp

/**
step1:
定义函数 f(i)表示第i天结束的最大收益
step2:
f(i)(0)状态 持有股票
f(i)(1)状态 无股票可以买
f(i)(2)状态 无股票不可买
step3:
目标值：max(f(n)(0)、f(n)(1)、f(n)(2))
step4:
转移函数
f(i)(0) = max(f(i-1)(1) - prices[i], f(i-1)(0))
f(i)(1) = max(f(i-1)(1), f(i-1)(2))
f(i)(2) = f(i-1)(0) + prices[i]
临界值
f(0)(0）= -prices[0]
f(0)(1) = 0
f(0)(0) = 0
*/
func maxProfit2(prices []int) int {
	l := len(prices)
	dp := make([][3]int, l)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][1]-prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i]
	}
	return max(dp[l-1][1], dp[l-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
