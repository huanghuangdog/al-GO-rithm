package dfs

/**
20210812
 */
func permute(nums []int) (rs [][]int) {
	size := len(nums)
	m := make(map[int]struct{})
	tmp := make([]int, 0)
	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(tmp) == size {
			tmpResult := make([]int, size)
			copy(tmpResult, tmp)
			rs = append(rs, tmpResult)
			return
		}
		for _, val := range nums {
			if _, ok := m[val]; ok {
				continue
			}
			tmp = append(tmp, val)
			m[val] = struct{}{}
			dfs(nums)
			tmp = tmp[:len(tmp) - 1]
			delete(m, val)
		}
	}
	dfs(nums)
	return
}
