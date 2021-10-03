func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 前序左子树index
	preoLeftIndex := len(inorder[:i]) + 1
	root.Left = buildTree(preorder[1:preoLeftIndex], inorder[:i])
	root.Right = buildTree(preorder[preoLeftIndex:], inorder[i+1:])
	return root
}