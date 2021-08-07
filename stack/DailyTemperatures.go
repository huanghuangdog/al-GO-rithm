package stack

func dailyTemperatures(temperatures []int) []int {
	rs := make([]int, len(temperatures))
	stack := make([]int, 0)
	m := make(map[int]int)
	for i := len(temperatures) - 1;i >= 0;i-- {
		m[temperatures[i]] = i
		for len(stack) != 0 && temperatures[i] >= stack[len(stack) - 1] {
			stack = stack[:len(stack) - 1]
		}
		if len(stack) == 0 {
			rs[i] = 0
		} else {
			rs[i] = m[stack[len(stack) - 1]] - i
		}
		stack = append(stack, temperatures[i])
	}
	return rs
}