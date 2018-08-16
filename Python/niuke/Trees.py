# 二叉搜索树第K个节点
# 给定一棵二叉搜索树，请找出其中的第k小的结点。例如， （5，3，7，2，4，6，8）    中，按结点数值大小顺序第三小结点的值为4。
# 思路:其实就是中序遍历，稍微注意一下边界条件就好了

class Solution:
    # 返回对应节点TreeNode
    def KthNode(self, pRoot, k):
        result = self.mid_stack(pRoot)
        if k <= 0 or k > len(result):
            return None

        node = result[k - 1]
        return node

    def mid_stack(self, root):
        if root == None:
            return []
        stack = []
        result = []
        node = root
        while node or stack:
            while node:
                stack.append(node)
                node = node.left
            node = stack.pop()
            result.append(node)
            node = node.right
        return result