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

# 二叉树序列化与反序列化
# 直接参照这个把 https://blog.csdn.net/u010005281/article/details/79787278
class Solution:
    def Serialize(self, root):
        if not root:
            return '#'
        return str(root.val) + ',' + self.Serialize(root.left) + ',' + self.Serialize(root.right)
    def Deserialize(self, s):
        list = s.split(',')
        return self.deserializeTree(list)
    def deserializeTree(self,list):
        if len(list) <= 0:
            return None
        val = list.pop(0)
        root = None
        if val != '#':
            root = TreeNode(int(val))
            root.left = self.deserializeTree(list)
            root.right = self.deserializeTree(list)
        return root

# 从上到下分层输出二叉树，必须返回二维列表
# 思路:因为要返回二维列表，所以需要对一般的层次遍历做一些修改。

    # 返回二维列表[[1,2],[4,5]]
    def Print(self, pRoot):
        if pRoot == None:
            return []
        result,nodes = [],[pRoot]
        while nodes :
            curStack,nextStack = [],[]
            for node in nodes:
                curStack.append(node.val)
                if node.left != None:
                    nextStack.append(node.left)
                if node.right != None:
                    nextStack.append(node.right)
            result.append(curStack)
            nodes = nextStack
        return result
# 按之字形打印二叉树
#  请实现一个函数按照之字形打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右至左的顺序打印，第三行按照从左到右的顺序打印，其他行以此类推。
# 其实就是上面那一题，按照求余规则，对子数组翻转一下即可

    def Print(self, pRoot):
        result, nodes, index = [], [pRoot], 1
        if pRoot == None:
            return []
        while nodes:
            curStack, nextStack = [], []
            for node in nodes:
                curStack.append(node.val)
                if node.left != None:
                    nextStack.append(node.left)
                if node.right != None:
                    nextStack.append(node.right)
            nodes = nextStack
            result.append(curStack)
        realResult = []
        for arr in result:
            if index % 2 == 0:
                arr.reverse()
            realResult.append(arr)
            index += 1
        return realResult