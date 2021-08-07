func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    stack := []*TreeNode{}
    
    for root != nil || len(stack) != 0 {
        if root != nil {
            stack = append(stack,root)
            root = root.Left

        } else {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res = append(res,root.Val)
            root = root.Right
        }

    }
    return res
}