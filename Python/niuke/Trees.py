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



    # 生成镜像二叉树
    def mirror(self, root):
        if root == None:
            return
        if root.left != None:
            self.mirror(root.left)
        if root.right != None:
            self.mirror(root.right)
        root.left, root.right = root.right, root.left
        return root

    # 判断二叉树是否对称
    def isSameTree(p, q):
        if not p and not q:
            return True
        if p and q and p.val == q.val:
            l = isSameTree(p.left, q.right)
            r = isSameTree(p.right, q.left)
            return l and r
        else:
            return False

    def symmetry(root):
        if root == None:
            return True
        else:
            return isSameTree(root.left, root.right)

    # 二叉树的下一个节点
    # 给定一个二叉树和其中一个节点，找出中序遍历的下一个节点并返回。
    def mid_stack(root):
        if root == None:
            return []
        result = []
        stack = []
        node = root
        while node or stack:
            while node:
                stack.append(node)
                node = node.left
            node = stack.pop()
            result.append(node)
            node = node.right


    # 二叉搜索树与双向链表
    # 中序遍历后，遍历排序数组，right指向下一个元素，下一个元素left指向当前
    def Convert(pRootOfTree):

         if pRootOfTree == None:
            return
        result = mid_stack(pRootOfTree)
        for i, v in enumerate(result[:-1]):  # result[:-1]表示遍历到倒数第二个，防止越界，如果遍历最后一个，执行i+1就越界了

            v.right = result[i + 1]
            result[i + 1].left = v
        return result[0]

    # 二叉搜索树的后序遍历序列
    # 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。
    # 思路：抛开写程序的舒服，手工分析。给定一个明确地数组例子，可以直接看出根节点和左右子树部分，然后对左右子树再次递归

    def isBST(array,sublen):
        if array is None or len(array) = 0:
            return False
        devide = 0
        for i,v in enumerate(array[:sublen-1]):
            if v > array[sublen-1]:
                devide = i
                break



        left = isBST(array,devide-1)
        right = isBST(array,devide,sublen)


        return left and right


