package pointer

/**
20210815
 */
func twoSum(numbers []int, target int) (rs []int) {
	size := len(numbers)
	left := 0
	right := size - 1
	for left < right {
		if numbers[left] + numbers[right] == target {
			rs = append(rs, left + 1, right + 1)
			return
		} else if numbers[left] + numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return
}