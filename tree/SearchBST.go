package tree

/**
20210805 思念是一种病
 */
func searchBST(root *TreeNode, val int) (rs *TreeNode) {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if val > root.Val {
		rs = searchBST(root.Right, val)
	} else {
		rs = searchBST(root.Left, val)
	}
	return
}