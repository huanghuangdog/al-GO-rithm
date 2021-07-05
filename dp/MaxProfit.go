package dp

func maxProfit(prices []int) int {
	l := len(prices)
	max := 0
	tmp := 0
	for i := 1; i < l; i++ {
		gap := prices[i] - prices[i-1]
		if gap+tmp > 0 {
			tmp = gap + tmp
		} else {
			tmp = 0
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
