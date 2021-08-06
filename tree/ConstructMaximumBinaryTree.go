package tree

import "math"

/**
20210806
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index, val := findMax(nums)
	node := &TreeNode{Val: val}
	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index + 1:])
	return node
}

func findMax(nums []int) (int, int) {
	max := math.MinInt64
	index := 0
	for i:= 0;i < len(nums);i++ {
		if nums[i] > max {
			index = i
			max = nums[i]
		}
	}
	return index, max
}