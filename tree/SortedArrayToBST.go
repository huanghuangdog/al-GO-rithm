package tree

/**
20210701
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	rootIdx := len(nums) / 2
	node := &TreeNode{Val: nums[rootIdx]}
	leftNode := sortedArrayToBST(nums[:rootIdx])
	rightNode := sortedArrayToBST(nums[rootIdx+1:])
	node.Left = leftNode
	node.Right = rightNode
	return node
}
