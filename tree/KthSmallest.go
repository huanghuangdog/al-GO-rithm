package tree

func kthSmallest(root *TreeNode, k int) int {
	list := traversa(root)
	return list[k - 1].Val
}

func traversa(node *TreeNode) []*TreeNode {
	list := make([]*TreeNode, 0)
	if node != nil {
		list = append(list, traversa(node.Left)...)
		list = append(list, node)
		list = append(list, traversa(node.Right)...)
	}
	return list
}