#### 1. 二叉树添加节点
```
def add(self,elem):
	node = Node(elem)
	if self.root.elem == -1: #没有节根点，对根节点赋值
		self.root = node
		self.myQueue.append(self.root)
	else:
		treeNode = self.myQueue[0] #此节点的子树没有齐
		if treeNode.lchild == None: #添加节点时，从左边开始添加
			treeNode.lchild = node
			self.myQueue.append(treeNode.lchild)
		else:
			treeNode.rchild = node
			self.myQueue.append(treeNode.rchild)
			self.myQueue.pop(0) #如果该节点存在右子树，将此节点抛弃。
```

#### 2. 三种遍历，分别递归和非递归算法
```

result = []
def front_recursive(self,root):
	if root == None:
		return
	result.append(root.elem)
	self.front_recursive(root.lchild)
	self.front_recursive(root.rchild)

def front_stack(self,root):
   if root == None:
       return []
   stack,result = [],[]
   while root or stack:
       while root:
           result.append(root.val)
           stack.append(root)
           root = root.left
       root = stack.pop()
       root = root.right
   return result


def mid_recursive(self,root):
	if root == None:
		return
	self.mid_recursive(self.lchild)
	result.append(root.elem)
	self.mind_recursive(self.rchild)

def mid_stack(self,root):
	if root == None:
		return
	stack = []
	result = []
	while node or stack:
		while node:  # 从根节点开始，一直找到左子树
			stack.append(node)
			node = node.left
		node = stack.pop() #while结束，表示当前没有左子树了,最后一个 node.left为空，所以栈上保存的是最下面的左子树
		result.append(node.elem)
		node = node.right 


def later_recursive(self,root):
	if root == None:
		return
	self.later_recursive(self.lchild)
	self.later_recursive(self.rchild)
	result.append(root.elem)

def later_stack(self,root):
	if root == None:
		return
	stack1 = []
	stack2 = []
	result = []
	stack1.append(node)
	while stack1: #找出后序遍历的逆序
		node = stack1.pop()
		if node.lchild:
			stack1.append(node.lchild)
		if node.rchild:
			stack1.append(node.right)
		stack2.append(node)
	while stack2: #stack2为后序遍历的逆序，出栈后即为后序遍历的序列
		result.append(stack2.pop())

```

#### 3.层次遍历
```
def level_queue(self,root):
	if root == None:
		return
	queue = []
	node = root
	queue.append(node)
	while queue:
		node = queue.pop()
		result.append(node.elem)
		if node.left != None:
			queue.append(node.lchild)
		if node.right != None:
			queue.append(node.rchild)
```

#### 4.给出前序和中序，构建二叉树
要记住一点，递归的时候，如果构建左子树，两个参数都是表示左子树的序列，只是一个参数表示中序，一个参数表示前序。构建右子树同理
```
def buildTree(self, preorder, inorder):
   """
   :type preorder: List[int]
   :type inorder: List[int]
   :rtype: TreeNode
   """
   length = len(preorder)
   if length == 0:
       return None
   if length == 1:
       return TreeNode(preorder[0])
   root = TreeNode(preorder[0])
   index = inorder.index(root.val)
   root.left = self.buildTree(preorder[1:index+1],inorder[0:index])
   root.right = self.buildTree(preorder[index+1:length],inorder[index+1:length])
   return root
```
#### 5.给出后序和中序，构建二叉树
和上面思路差不多。
```
def buildTree(self,inorder,postorder):
	 length = len(inorder)
   if length == 0:
       return None
   if length == 1:
       return TreeNode(postorder[length-1])
   root = TreeNode(postorder[length-1])
   index = inorder.index(root.val)
   root.left = self.buildTree(inorder[0:index],postorder[0,index])
   root.right = self.buildTree(inorder[index+1:length],postorder[index+1:length-1]) 
```

#### 6.给定排序好的数组，把它转换为二叉查找树，要求生成的二叉查找树是平衡的。
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

因为数组已经排序好了，取中间的元素作为二叉查找树的根节点，然后递归生成。可以不区分奇偶，注意python3的除法问题即可
```
def sortedArrayToBST(self, nums):
   """
   :type nums: List[int]
   :rtype: TreeNode
   """
   length = len(nums)
   if length == 0:
       return None
   if length == 1:
       return TreeNode(nums[0])
   root = TreeNode(nums[length//2])
   root.left = self.sortedArrayToBST(nums[:length//2])
   root.right = self.sortedArrayToBST(nums[length//2+1:])
   return root
```

#### 7. 反向层次遍历二叉树,按层次输出
https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
层次遍历结果反过来即可。
```
def levelOrderBottom(self, root):
   """
   :type root: TreeNode
   :rtype: List[List[int]]
   """
   tree = []
   if root == None:
       return tree
   curr_level = [root]
   while curr_level:
       level_list = []
       next_list = []
       for tmp in curr_level:
           level_list.append(tmp.val)
           if tmp.left != None:
               next_list.append(tmp.left)
           if tmp.right != None:
               next_list.append(tmp.right)
       tree.append(level_list)
       curr_level = next_list
   tree.reverse()
   return tree
```
#### 8.之字形层次遍历二叉树
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
其实也是层次遍历，和上面那题一样,加个奇偶判断就可以了。层次遍历基本都是这个套路
```
def zigzagLevelOrder(self, root):
   """
   :type root: TreeNode
   :rtype: List[List[int]]
   """
   tree = []
   if root == None:
       return tree
   curr_level = [root]
   num = 1
   while curr_level:
       curr_list = []
       next_list = []
    
       for tmp in curr_level:
           curr_list.append(tmp.val)
           if tmp.left:
               next_list.append(tmp.left)
           if tmp.right:
               next_list.append(tmp.right)
               
       if num % 2 == 0:
           curr_list.reverse()
       tree.append(curr_list)
       curr_level = next_list
       num += 1

   return tree
```
#### 9.二叉树路径总和
https://leetcode.com/problems/path-sum/description/
给定一棵树和一个和，判断是否存在根到叶子节点的路径，使得路径上节点总和为给定的值。必须贯穿到最下面的叶子节点。

题目说到从根到XX节点，那肯定是优先考虑先序遍历,或者深度优先。
然后就是二叉树题，没什么额外要求的话，优先考虑递归解决。
```
def hasPathSum(self, root, sum):
   """
   :type root: TreeNode
   :type sum: int
   :rtype: bool
   """
   if root == None:
       return False
   if root.left == None and root.right == None:
       return sum == root.val
   return self.hasPathSum(root.left,sum-root.val) or self.hasPathSum(root.right,sum-root.val)
```
#### 10.列举二叉树路径总和的结果
https://leetcode.com/problems/path-sum-ii/description/
这里需要使用嵌套函数


#### 11.二叉树的下一个节点
https://www.nowcoder.com/practice/9023a0c988684a53960365b889ceaf5e?tpId=13&tqId=11210&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
这题分三个步骤，首先是通过next找到根节点，然后进行中序遍历，最后找到目标节点的下一个节点。
```
def GetNext(self, pNode):
        dummy = pNode
        while dummy.next:
            dummy = dummy.next
        #找到根节点
        self.result = []
        #中序遍历
        self.midReverse(dummy)
        #条件检查，输出目标节点
        if self.result.index(pNode) != len(self.result) -1:
            pIndex = self.result.index(pNode)
            return self.result[pIndex+1]
        else:
            return None
        
    def midReverse(self,node):
        if node == None:
            return 
        self.midReverse(node.left)
        self.result.append(node)
        self.midReverse(node.right)
```
#### 12.判读对称二叉树（镜像）
https://www.nowcoder.com/practice/ff05d44dfdb04e1d83bdbdab320efbcb?tpId=13&tqId=11211&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
主要思路就是递归遍历二叉树的左右子树
```
    def isSymmetrical(self, pRoot):
        # write code here
        def isSameTree(p, q):
            if q == None and p == None:
                return True
            if q and p and q.val == p.val:
                l = isSameTree(p.left,q.right)
                r = isSameTree(q.left,p.right)
                return l and r
            else:
                return False
            
        if pRoot == None:
            return True
        else:
            return isSameTree(pRoot.left,pRoot.right)
```
#### 12.X 相似题，判断二叉树是否相同
使用递归
```
    def isSameTree(self, p, q):
        """
        :type p: TreeNode
        :type q: TreeNode
        :rtype: bool
        """
        if p == None and q == None:
            return True
        if p and q and p.val == q.val:
            l = self.isSameTree(p.left,q.left)
            r = self.isSameTree(p.right,q.right)
            return l and r
        else:
            return False
```

#### 13.层次输出二叉树
https://www.nowcoder.com/practice/445c44d982d04483b04a54f298796288?tpId=13&tqId=11213&tPage=2&rp=2&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
这个其实是层次遍历二叉树的升级版，主要是构造两个栈来辅助
```
   def isSameTree(self, p, q):
        """
        :type p: TreeNode
        :type q: TreeNode
        :rtype: bool
        """
        if p == None and q == None:
            return True
        if p and q and p.val == q.val:
            l = self.isSameTree(p.left,q.left)
            r = self.isSameTree(p.right,q.right)
            return l and r
        else:
            return False
```
#### 14.序列化二叉树
https://www.nowcoder.com/practice/cf7e25aa97c04cc1a68c8f040e71fb84?tpId=13&tqId=11214&tPage=2&rp=2&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
算法思路:
序列化就是将节点值转成字符串，逗号隔开，空白节点用 # 代替。使用前序遍历
```
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
```
#### 15.二叉搜索树第K个节点
https://www.nowcoder.com/practice/ef068f602dde4d28aab2b210e859150a?tpId=13&tqId=11215&tPage=2&rp=2&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
中序遍历搞定，注意判断条件即可。
```
# 返回对应节点TreeNode
    def KthNode(self, pRoot, k):
        self.result = []
        self.mid(pRoot)
        if 0<k<=len(self.result):
            return self.result[k-1]
        else:
            return None
    
    def mid(self,root):
        if root == None:
            return 
        self.mid(root.left)
        self.result.append(root)
        self.mid(root.right)
```

#### 16.翻转二叉树
https://leetcode.com/problems/invert-binary-tree/
经典的翻转二叉树，提供递归和非递归两种解法
```
    def invertTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if root == None:
            return None
        root.left,root.right = self.invertTree(root.right),self.invertTree(root.left)

        return root
```
```
        if root == None:
            return None
        queue = []
        queue.append(root)
        while queue:
            node = queue.pop(0)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
            node.left,node.right = node.right,node.left
        return root
```


```
fast = head.next
slow = head
while fast:
	fast = fast.next.next
	slow = slow.next

pHead = head
last = None
while pHead.val != slow.val
	temp = pHead.next
	pHead.next = last
	last = pHead
	pHead = temp

```

```
def shift_down(arr,start,end):
	root = start
	while True:
		child = 2 * root + 1
		if child > end:
			return 
		if child+1 <= end and arr[child] < arr[child+1]:
			child +=1
		if arr[root] < arr[child]:
			arr[root],arr[child] = arr[root],arr[child]
			root = child
		else:
		   break

def heapSort(arr):
	firt = len(arr) // 2 - 1
	for start in range(first,-1,-1):
		shift_down(arr,start,len(arr)-1)
	for end in range(len(arr)-1,0,-1):
	   shift_down(arr,0,end-1)
	return arr
```
##### 98. 验证二叉搜索树
只要说道二叉搜索树，就可以考虑和中序遍历关联起来。本题思路非常简单，中序遍历结果只要不是升序的，就不是二叉搜索树。这里要注意一个技巧，go语言没有栈，只能用数组操作来模拟栈。
```
import "math"

func isValidBST(root *TreeNode) bool {
    var stack = make([]*TreeNode,0)
    inorder := math.MinInt64

    for len(stack) > 0 || root != nil {
        for root != nil {
            //入栈
            stack = append(stack,root)
            root = root.Left
        }
        //出栈
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Val <= inorder {
            return false
        } 
        
        inorder = root.Val
        root = root.Right
        
    }
    return true
}

```

##### 530 二叉搜索树的最小绝对差

中序遍历可以解决大多数二叉搜索树问题
非递归解法：
```
func getMinimumDifference(root *TreeNode) int {
    stack := make([]*TreeNode,0)
    re := -1
    var pre *TreeNode = nil

    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack,root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if pre != nil {
            if re == -1 {
                re = root.Val - pre.Val
            } else {
                re = Min(re,root.Val - pre.Val)
            }
        }

        pre = root
        root = root.Right
    }

    return re
}

func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

```
这题的递归解法，在leetcode上反而更快。下面是递归解法


