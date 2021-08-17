package sort

/**
20210817
 */
func findKthLargest(nums []int, k int) int {
	var quickSort func(start int, stop int) int
	quickSort = func(start int, stop int) int {
		if start >= stop {
			return start
		}
		left := start
		right := stop
		tmp := nums[left]
		for left != right {
			for left < right && nums[right] >= tmp {
				right--
			}
			for left < right && nums[left] <= tmp {
				left++
			}
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
			}
		}
		nums[right], nums[start] = tmp, nums[right]
		if right == len(nums) - k {
			return nums[right]
		} else if right > len(nums) - k {
			return quickSort(start, right - 1)
		} else {
			return quickSort(right + 1, stop)
		}
	}
	return quickSort(0, len(nums) - 1)
}
