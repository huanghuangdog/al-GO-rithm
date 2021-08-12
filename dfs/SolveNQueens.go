package dfs

/**
20210812
 */
func solveNQueens(n int) (rs [][]string) {
	var dfs func(row int)
	tmp := make([]int, 0)
	dfs = func(row int) {
		if len(tmp) == n {
			rs = append(rs, transform(tmp))
			return
		}
		if row == n {
			return
		}
		for i:= 0;i < n;i++ {
			tmp = append(tmp, i)
			if !check(tmp, row) {
				tmp = tmp[:len(tmp) - 1]
				continue
			} else {
				dfs(row + 1)
			}
			tmp = tmp[:len(tmp) - 1]
		}
	}
	dfs(0)
	return
}

func transform(tmp []int) []string {
	var print func(index int) string
	print = func(index int) string {
		tmpStr := make([]byte,0)
		for i:= 0;i < len(tmp);i++ {
			if i == index {
				tmpStr = append(tmpStr, 'Q')
			} else {
				tmpStr = append(tmpStr, '.')
			}
		}
		return string(tmpStr)
	}
	rs := make([]string, 0)
	for _, val := range tmp {
		rs = append(rs, print(val))
	}
	return rs
}

func check(tmp []int, index int) bool {
	for i:= 0;i < index;i++ {
		if tmp[index] == tmp[i] || abs(index -i) == abs(tmp[index] - tmp[i]) {
			return false
		}
	}
	return true
}
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
