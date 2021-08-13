package dfs

/**
20210813
 */
func subsets(nums []int) (rs [][]int) {
	var dfs func(index int)
	tmp := make([]int, 0)
	dfs = func(index int) {
		if index == len(nums) {
			rs = append(rs, append([]int(nil), tmp...))
			return
		}
		tmp = append(tmp, nums[index])
		dfs(index + 1)
		tmp = tmp[:len(tmp) - 1]
		dfs(index + 1)
	}
	dfs(0)
	return
}
