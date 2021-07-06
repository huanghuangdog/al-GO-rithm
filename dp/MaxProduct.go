package dp

func maxProduct(nums []int) int {
	size := len(nums)
	var maxNum func([]int) int
	maxNum = func(ints []int) int {
		rs := ints[0]
		for i := 1; i < len(ints); i++ {
			if ints[i] > rs {
				rs = ints[i]
			}
		}
		return rs
	}

	var minNum func([]int) int
	minNum = func(ints []int) int {
		rs := ints[0]
		for i := 1; i < len(ints); i++ {
			if ints[i] < rs {
				rs = ints[i]
			}
		}
		return rs
	}

	max := nums[0]
	tmpMin, tmpMax := nums[0], nums[0]
	for i := 1; i < size; i++ {
		tmpMin, tmpMax = minNum([]int{nums[i], nums[i] * tmpMin, nums[i] * tmpMax}), maxNum([]int{nums[i], nums[i] * tmpMin, nums[i] * tmpMax})
		if tmpMax > max {
			max = tmpMax
		}
	}
	return max
}
