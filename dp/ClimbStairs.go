package dp

func climbStairs(n int) int {
	rs := make([]int, n+1)
	rs[0] = 0
	rs[1] = 1
	for i := 2; i <= n; i++ {
		rs[i] = rs[i-1] + rs[i-2]
	}
	return rs[n]
}
