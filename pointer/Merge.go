package pointer

/**
20210815
 */
func merge(nums1 []int, m int, nums2 []int, n int)  {
	total := m + n
	last := total - 1
	i := m - 1
	j := n - 1
	for j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[last] = nums1[i]
			i--
		} else {
			nums1[last] = nums2[j]
			j--
		}
		last--
	}
}
