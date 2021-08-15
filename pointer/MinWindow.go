package pointer

/**
20210815
 */
func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)
	sm := make(map[byte]int)
	tm := make(map[byte]int)
	for i:= 0;i < tLen;i++ {
		tm[t[i]]++
	}
	var match func() bool
	match = func() bool {
		for k, v := range tm {
			if sm[k] < v {
				return false
			}
		}
		return true
	}

	rsLen := sLen + 1
	rs := ""
	for left, right := 0, 0; right < sLen;right++ {
		if tm[s[right]] > 0 {
			sm[s[right]]++
		}
		for match() && right >= left {
			if right - left + 1 < rsLen {
				rsLen = right - left + 1
				rs = s[left:left + rsLen]
			}
			sm[s[left]]--
			left++
		}
	}
	return rs
}

