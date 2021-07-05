package dp

func numDecodings(s string) int {
	size := len(s)
	rs := make([]int, size)
	if s[0] == '0' {
		return 0
	}

	rs[0] = 1
	lastChar := s[0]
	for i := 1; i < size; i++ {
		c := s[i]
		switch c {
		case '0':
			if lastChar == '1' || lastChar == '2' {
				if i > 1 {
					rs[i] = rs[i-2]
				} else {
					rs[i] = 1
				}
			} else {
				return 0
			}
		case '1', '2', '3', '4', '5', '6':
			if lastChar == '1' || lastChar == '2' {
				if i > 1 {
					rs[i] = rs[i-1] + rs[i-2]
				} else {
					rs[i] = rs[i-1] + 1
				}
			} else {
				rs[i] = rs[i-1]
			}
		case '7', '8', '9':
			if lastChar == '1' {
				if i > 1 {
					rs[i] = rs[i-1] + rs[i-2]
				} else {
					rs[i] = rs[i-1] + 1
				}
			} else {
				rs[i] = rs[i-1]
			}
		}
		lastChar = c
	}
	return rs[size-1]
}
