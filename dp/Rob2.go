package dp

import "fmt"

func rob2(nums []int) int {
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

	var innerRob func([]int) int
	innerRob = func(ints []int) int {
		l := len(ints)
		if l == 0 {
			return 0
		} else if l == 1 {
			return ints[0]
		} else if l == 2 {
			return maxNum([]int{ints[1], ints[0]})
		}

		prev, next := ints[0], maxNum([]int{ints[0], ints[1]})
		for i := 2; i < l; i++ {
			prev, next = next, maxNum([]int{prev + ints[i], next})
		}
		fmt.Printf("input:%v,output:%v\n", ints, next)
		return next
	}

	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	} else if l == 2 {
		return maxNum([]int{nums[0], nums[1]})
	} else {
		return maxNum([]int{innerRob(nums[:l-1]), innerRob(nums[1:])})
	}
}
