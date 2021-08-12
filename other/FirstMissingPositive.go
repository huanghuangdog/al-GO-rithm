package other

/**
20210812 充分利用已有的空间
 */
func firstMissingPositive(nums []int) int {
	size := len(nums)
	for i:= 0;i < size;i++ {
		for nums[i] > 0 && nums[i] <= size && nums[i] != nums[nums[i] - 1] {
			nums[nums[i] - 1], nums[i] = nums[i], nums[nums[i] - 1]
		}
	}
	for i := 0;i < size;i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return size + 1
}