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
		return
	stack = []
	result = []
	node = root
	while node or stack:
		while node:
			result.append(node.elem)
			stack.append(node)
		node = stack.pop()
		node = node.right


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

