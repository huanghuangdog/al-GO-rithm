package sw

/**
20210807
 */
func maxSlidingWindow(nums []int, k int) []int {
	rs := make([]int, 0)
	queue := make([]int, 0)
	for i := 0;i < k;i++ {
		queue = push(queue, nums, i, k)
	}

	for i:= k;i < len(nums);i++ {
		rs = append(rs, nums[queue[0]])
		queue = push(queue, nums, i, k)
	}
	rs = append(rs, nums[queue[0]])
	return rs
}

func push(queue []int, nums []int, index int, k int) []int {
	for len(queue) != 0 && index >= k && queue[0] <= index -k {
		queue = queue[1:]
	}

	for i:= len(queue) - 1;i >= 0;i-- {
		if nums[index] >= nums[queue[i]] {
			queue = queue[:len(queue) -1]
		}
	}
	queue = append(queue, index)
	return queue
}
