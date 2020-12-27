func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	} else {
		return min(left, right) + 1
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}