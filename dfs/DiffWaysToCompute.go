package dfs

import "strconv"

/**
20210822
 */
func diffWaysToCompute(expression string) []int {
	rs := make([]int, 0)
	size := len(expression)
	for i:= 0;i < size;i++ {
		cur := expression[i]
		if cur == '+' || cur == '-' || cur == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i + 1:])
			for _, leftVal := range left {
				for _, rightVal := range right {
					switch cur {
					case '+':
						rs = append(rs, leftVal + rightVal)
					case '-':
						rs = append(rs, leftVal - rightVal)
					case '*':
						rs = append(rs, leftVal * rightVal)
					}
				}
			}
		}
	}
	if len(rs) == 0 {
		item, _ := strconv.Atoi(expression)
		return append([]int(nil), item)
	}
	return rs
}