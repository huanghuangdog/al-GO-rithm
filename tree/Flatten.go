package tree

func flatten(root *TreeNode) {
	list := traversal(root)
	if len(list) == 0 {
		return
	}
	node := list[0]
	for i:= 1;i < len(list);i++ {
		node.Left = nil
		node.Right = list[i]
		node = list[i]
	}
}

func traversal(node *TreeNode) []*TreeNode {
	list := make([]*TreeNode, 0)
	if node != nil {
		list = append(list, node)
		list = append(list, traversal(node.Left)...)
		list = append(list, traversal(node.Right)...)
	}
	return list
}