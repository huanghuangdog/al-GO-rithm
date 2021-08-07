package stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	tmp := make([]int, len(nums2))
	stack := make([]int, 0)
	m := make(map[int]int)
	for i:= len(nums2) - 1;i >= 0;i-- {
		m[nums2[i]] = i
		for len(stack) != 0 && stack[len(stack) - 1] <= nums2[i] {
			stack = stack[:len(stack) - 1]
		}
		if len(stack) == 0 {
			tmp[i] = -1
		} else {
			tmp[i] = stack[len(stack) - 1]
		}
		stack = append(stack, nums2[i])
	}

	rs := make([]int, 0)
	for _, val := range nums1 {
		rs = append(rs, tmp[m[val]])
	}
	return rs
}
