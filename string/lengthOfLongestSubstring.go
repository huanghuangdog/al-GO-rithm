package string

/**
20210822
 */
func lengthOfLongestSubstring(s string) (rs int) {
	size := len(s)
	m := make(map[byte]int)
	right := 0
	for left := 0; left < size;left++ {
		for ; right < size && m[s[right]] == 0;right++ {
			m[s[right]]++
		}
		if right - left > rs {
			rs = right - left
		}
		delete(m, s[left])
	}
	return
}