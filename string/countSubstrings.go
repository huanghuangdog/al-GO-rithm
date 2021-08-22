package string

/**
20210821
 */
func countSubstrings(s string) int {
	size := len(s)
	var count func(left int, right int) int
	count = func(left int, right int) int {
		cnt := 0
		for ;left >= 0 && right < size && s[left] == s[right]; {
			left--
			right++
			cnt++
		}
		return cnt
	}
	total := 0
	for i:= 0;i < size;i++ {
		total += count(i, i)
		total += count(i, i + 1)
	}
	return total
}
