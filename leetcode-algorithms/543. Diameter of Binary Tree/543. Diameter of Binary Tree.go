/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var res int

 func diameterOfBinaryTree(root *TreeNode) int {
	 if root == nil {
		 return 0
	 }
	 res = 0
	 // 转化成左右子树深度相加
	 maxDepth(root)
	 return res
 }
 
 func maxDepth(root *TreeNode) int {
	 // 如果节点为空就不再递归下探深度
	 if root == nil {
		 return 0
	 }
	 left := maxDepth(root.Left)
	 right := maxDepth(root.Right)
	 path := max(left,right)
 
	 res = max(left+right,res)
	 
	 return path+1
 }
 
 func max(a, b int) int {
	 if a < b {
		 return b
	 }
	 return a
 }
 