import "math"

func isValidBST(root *TreeNode) bool {
	var stack = make([]*TreeNode, 0)
	inorder := math.MinInt64

	for len(stack) > 0 || root != nil {
		for root != nil {
			//入栈
			stack = append(stack, root)
			root = root.Left
		}
		//出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}

		inorder = root.Val
		root = root.Right

	}
	return true
}