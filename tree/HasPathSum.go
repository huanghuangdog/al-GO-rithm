package tree

/**
20210813
 */
func pathSum(root *TreeNode, targetSum int) (rs [][]int) {
	tmp := make([]int, 0)
	var track func(node *TreeNode, num int)
	track = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		tmp = append(tmp, node.Val)
		defer func() { tmp = tmp[:len(tmp) - 1]}()
		if node.Left == nil && node.Right == nil && node.Val == num {
			tmpResult := make([]int, len(tmp))
			copy(tmpResult, tmp)
			rs = append(rs, tmpResult)
		}
		track(node.Left, num - node.Val)
		track(node.Right, num - node.Val)
	}
	track(root, targetSum)
	return
}
