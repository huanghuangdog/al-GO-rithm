package other

import "math"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	minSize := math.MaxInt32
	for _, str := range strs {
		if len(str) < minSize {
			minSize = len(str)
		}
	}

	if minSize < 1 {
		return ""
	}

	rs := 0
	for i := 0; i < minSize; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != c {
				return strs[0][0:i]
			}
		}
		rs = i
	}
	return strs[0][0 : rs+1]
}
