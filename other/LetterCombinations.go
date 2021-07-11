package other

import "strconv"

func letterCombinations(digits string) (rs []string) {
	m := make([][]string, 0)
	m = append(m, []string{"a", "b", "c"})
	m = append(m, []string{"d", "e", "f"})
	m = append(m, []string{"g", "h", "i"})
	m = append(m, []string{"j", "k", "l"})
	m = append(m, []string{"m", "n", "o"})
	m = append(m, []string{"p", "q", "r", "s"})
	m = append(m, []string{"t", "u", "v"})
	m = append(m, []string{"w", "x", "y", "z"})
	size := len(digits)
	for i := 0; i < size; i++ {
		c, _ := strconv.Atoi(string(digits[i]))
		cInt := int(c) - 2
		if i == 0 {
			rs = append(rs, m[cInt]...)
		} else {
			l := len(rs)
			for j := 0; j < l; j++ {
				for k := 0; k < len(m[cInt]); k++ {
					rs = append(rs, rs[j]+m[cInt][k])
				}
			}
			rs = rs[l:]
		}
	}
	return
}
