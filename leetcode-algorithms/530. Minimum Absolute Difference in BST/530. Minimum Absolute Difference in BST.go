func getMinimumDifference(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	re := -1
	var pre *TreeNode = nil

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil {
			if re == -1 {
				re = root.Val - pre.Val
			} else {
				re = Min(re, root.Val-pre.Val)
			}
		}

		pre = root
		root = root.Right
	}

	return re
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}