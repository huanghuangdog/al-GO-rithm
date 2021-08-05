package tree

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(node *TreeNode, minNode *TreeNode, maxNode *TreeNode) bool {
	if node == nil {
		return true
	}
	if minNode != nil && node.Val <= minNode.Val {
		return false
	}
	if maxNode != nil && node.Val >= maxNode.Val {
		return false
	}

	return isValid(node.Left, minNode, node) && isValid(node.Right, node, maxNode)
}
