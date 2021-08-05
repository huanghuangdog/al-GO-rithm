package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := preorder[0]
	node := &TreeNode{Val: root}
	rootIndex := findRoot(inorder, root)
	if rootIndex != 0 {
		leftNode := buildTree(preorder[1:1 + rootIndex], inorder[0:rootIndex])
		node.Left = leftNode
	}
	if rootIndex != len(inorder) {
		rightNode := buildTree(preorder[1 + rootIndex:], inorder[rootIndex + 1:])
		node.Right = rightNode
	}
	return node
}

func findRoot(tree []int, val int) int {
	index := 0
	for ; index < len(tree); index++ {
		if tree[index] == val {
			return index
		}
	}
	return index
}