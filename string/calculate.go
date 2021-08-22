package string

/**
20210822
 */
func calculate(s string) (ans int) {
	size := len(s)
	num := 0
	prevOp := '+'
	tmp := make([]int, 0)
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = num * 10 + int(ch - '0')
		}
		if !isDigit && ch != ' ' || i == size - 1 {
			switch prevOp {
			case '+':
				tmp = append(tmp, num)
			case '-':
				tmp = append(tmp, -num)
			case '*':
				tmp[len(tmp) - 1] = tmp[len(tmp) - 1] * num
			case '/':
				tmp[len(tmp) - 1] = tmp[len(tmp) - 1] / num
			default:
			}
			prevOp = ch
			num = 0
		}

	}
	for _, val := range tmp {
		ans += val
	}
	return
}
