# -*- coding:utf8 -*-

from collections import deque

# binary tree mirror
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    # solution 1 --> recursion, 直接粗暴递归,交换两个子树即可

    def tree_mirror_1(root):
        if root:
            root.left, root.right = root.right, root.left
            tree_mirror_1(root.left)
            tree_mirror_1(root.right)

    # solution 2 从右到左，从上到下遍历。将以层次遍历形式保存到数组中
    def tree_mirror_2(self,root):
        ret = []
        queue = deque([root])
        while queue:
            node = queue.popleft()

            if node:
                ret.append(node)
                queue.append(node.right)
                queue.append(node.left)
        return ret

    # 从上往下打印二叉树，即层次遍历二叉树
    # 和上面一个类似
    def print_tree(self,root):
        ret = []
        queue = deque([root])
        while queue:
            node = queue.popleft()
            if node:
                ret.append(node.val)
                queue.append(node.left)
                queue.append(node.right)
        return ret

    # 二叉搜索树的后续遍历：判断给定的整数数组是不是二叉搜索树的后序遍历序列
    # 首先要清楚二叉搜索树的概念，右子树都大于根节点，左子树都小于根节点。后续遍历根节点在最右边。根据二叉搜索特性，可以把左右子树的序列区分开。然后再递归处理。

    def is_post_order(self,order):
        length = len(order)
        if length:
            root = order[-1]
            left = 0
            while order[left] < root:
                left += 1
            right = left
            while right < length - 1:
                if order[right] < root:
                    return False
                right += 1
            left_ret = True if left == 0 else is_post_order(order[:left])
            right_ret = True if left == right else is_post_order(order[left:right])
            return left_ret and right_ret  # 只有当左右子树都为后序遍历时，结果为后序遍历
        return False

    # 要求：输入一棵二叉树和一个值，求从根结点到叶结点的和等于该值的路径
    # 二叉树深度优先搜索，就是先序，中序，后序遍历。 广度优先搜索就是层次遍历
    # 其实问的就行先序遍历, 把遍历的节点保存到数组中
    #
    def find_path(self,root,num):
        res = []
        if not root:
            return res

        self.target = num
        self.dfs(root,res,[root.val])
        return res


    def dfs(self,root,res,path):
        if not root.left and root.right and sum(path) == self.target:
            res.append(path)
        if root.left:
            self.dfs(root.left,res,path + [root.left.val])
        if root.right:
            self.dfs(root,res,path + [root.right.val])




# 顺时针打印矩阵
# 将第一行加入result，然后删掉第一行，整个矩阵逆时针转90度（使用临时空间来辅助存储）,像魔方旋转那样。然后继续取第一行。知道整个矩阵被删空为止。
def print_matrix(matrix):
    result = []
    while matrix:
        result += matrix.pop(0)
        if not matrix or not matrix[0]:
            break
        matrix = turn(matrix)
    return result

def turn(matrix):
    num_r = len(matrix)
    num_c = len(matrix[0])
    print(num_r,num_c)
    newmat = []
    for i in range(num_c):
        newmat2 = []
        for j in range(num_r): # j 和 i 反过来，变成转置矩阵
            newmat2.append(matrix[j][i])
        newmat.append(newmat2)
    newmat.reverse()
    return newmat

'''
matrix =  [[1,2,3],[4,5,6]]
print(matrix)
print(print_matrix(matrix))
'''


# 设计一个栈，使得PUSH、POP以及GetMin（获取栈中最小元素）能够在常数时间内完成。
# 使用一个辅助栈，主栈每次push的时候，检查压入的元素是否小于或等于辅助栈栈顶元素，如果小于或等于，那么这个元素也push到辅助栈中。
# 主栈pop的时候，检查pop的元素是否等于辅助栈的栈顶元素，如果等于，辅助栈的栈顶元素也跟着pop
class min_stack:
    def __init__(self):
        self.stack = []
        self.min = []
    def push(self,val):
        self.stack.append(val)
        if self.min:
            minval = self.min[-1]
            if val <= minval:
                self.min.append(val)
        else:
            self.min.append(val)

    def pop(self):
        if self.stack:
            val = self.stack.pop()
            if val == self.min[-1]:
                self.min.pop()
            return val


# 输入两个整数序列，第一个序列表示栈的压入顺序，请判断二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。 这两个序列长度是相等的
# 思路：使用一个辅助栈. 例如:入栈顺序为 1 2 3 4 5，出栈顺序为 4 5 3 2 1 。遍历入栈，先取第一个元素1，加入辅助栈。判断辅助栈栈顶元素是否等于出栈，
# 显然1不等于4。继续遍历入栈，压入辅助栈。直到4。此时辅助栈从顶部到底部，分别是 4 ，3 ，2 ，1，栈顶刚好等于出栈序列的4。辅助栈出栈4，出栈序列去掉4.
# 此时辅助栈：3 2 1  ，出栈序列 5 3 2 1，接着操作，辅助栈继续入栈直到没元素入为止
# 辅助栈 5 3 2 1 出栈序列 5 3 2 1, 辅助栈出栈，出栈序列移除第一个
# 辅助栈  3  2 1 ，出栈序列 3  2 1
# 最终思路描述:构造一个辅助栈，还有一个索引值0，该索引值用于出栈序列元素读取。遍历入栈序列，在循环里面，入栈序列一个元素压入辅助栈，
# 然后此时如果辅助栈的栈顶元素等于出栈序列的索引位置的值，辅助栈顶部pop，索引加1。不满足就循环。最终，如果辅助栈为空，就返回true,否则返回false
def pop_order(push_stack,pop_stack):
    # 两个输入元素其实是数组
    if not push_stack or not pop_stack:
        return False
    # 辅助栈
    stack = []
    count = 0 #相当于指针标记，用于出栈队列
    for i in range(len(push_stack)):

        stack.append(push_stack[i])
        while stack and stack[-1] == pop_stack[count]:
            stack.pop()
            count += 1

    return not stack

# print(pop_order([1,2,3,4,5],[4,5,2,3,1]))


# ================================ 分解让复杂问题简单化
# 字符串全排列
def my_permutation(s):
    str_set = []
    ret = []  # 最后的结果

    def permutation(string):
        for i in string:
            str_tem = string.replace(i, '')
            str_set.append(i)
            if len(str_tem) > 0:
                permutation(str_tem)
            else:
                ret.append(''.join(str_set))
            str_set.pop()

    permutation(s)
    return ret

str = "1234"
print(my_permutation(str))







