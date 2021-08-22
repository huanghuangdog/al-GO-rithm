package string

/**
20210822
 */
func countBinarySubstrings(s string) int {
	size := len(s)
	prevCnt := 0
	curCnt := 1
	total := 0
	for i:=1;i < size;i++ {
		if s[i] == s[i-1] {
			curCnt++
		} else {
			prevCnt = curCnt
			curCnt = 1
		}
		if prevCnt >= curCnt {
			total++
		}
	}
	return total
}