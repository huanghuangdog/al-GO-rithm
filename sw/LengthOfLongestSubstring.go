package sw

func lengthOfLongestSubstring(s string) int {
	size := len(s)
	m := make(map[byte]int)
	end := -1
	max := 0
	for i := 0; i < size; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for end+1 < size && m[s[end+1]] == 0 {
			m[s[end+1]]++
			end++
		}
		if end-i+1 > max {
			max = end - i + 1
		}
	}
	return max
}
