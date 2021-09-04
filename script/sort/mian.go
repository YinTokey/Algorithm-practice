package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Initer interface {
	SetVal (val interface{})
}

type Operater interface {
	PrintBT()
	Depth() int
	LeafCount() int
}

type Order interface {
	PreOrder()
	InOrder()
	PostOrder()
}

type Node struct {
	Left *Node
	Val int
	Right *Node
}
// ---------------------------------------
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// 当数组用就行了，相当于给数组实现扩展方法
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Top() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	return x
}



// -----------------------------------------------

func main()  {
	// 默认小顶堆
	//h := &IntHeap{100,16,4,8,70,2,36,22,5,12}
	//heap.Init(h)
//	heap.Push(h, 3) // 往堆中推入元素
	//for h.Len() > 0 {
	//	fmt.Printf("%d ", heap.Pop(h))
	//}

	arr := []int{1,5,6,2,3,4}
	re := quickSort(arr,0,5)
	fmt.Println(re)
}



func minimumDifference(s []string, k int) string {

	sort.Slice(s, func(i, j int) bool {
		numA, _ := strconv.Atoi(s[i])
		numB, _ := strconv.Atoi(s[j])
		return numA < numB
	})

	fmt.Println(s)

	return s[len(s)-k+1]

}

//func findMin(nums []int) int {
//	heap.Init(nums)
//	heap.Push(h, 3) // 往堆中推入元素
//	for h.Len() > 0 {
//		fmt.Printf("%d ", heap.Pop(h))
//	}
//}


func minStoneSum(piles []int, k int) int {

	for k > 0 {
		idx := findBigest(piles)
		fmt.Println(idx)
		piles[idx] = int(math.Ceil(float64(piles[idx]) / 2))
		k--
	}

	return sumArr(piles)
}

func findBigest(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	re := 0
	big := arr[0]
	for i := 1; i < len(arr) ; i++ {
		v := arr[i]
		if v > big {
			big = v
			re = i
		}
	}

	return re
}

func sumArr(arr []int) int {
	re := 0
	for _,v := range arr {
		re += v
	}
	return re
}

func isPrefixString(s string, words []string) bool {
	if len(words) == 0 {
		return false
	}
	if len(words) == 1 {
		return s == words[0]
	}

	re := true

	for i := 0; i < len(words); i++ {
		if len(s) == 0 {
			return true
		}
		v := words[i]
		var sub string
		if len(v) <= len(s) {
			sub = s[:len(v)]
		} else {
			return false
		}
		fmt.Println(sub)
		fmt.Println(v)
		fmt.Println(s)
		if sub != v {
			return false
		}
		s = s[len(v):]
	}

	if len(s) > 0 {
		return false
	}

	return re
}

func preOrder(root *Node) {
	if root != nil {
		fmt.Println(root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}
}

func inOrder(root *Node) {
	if root != nil {
		preOrder(root.Left)
		fmt.Println(root.Val)
		preOrder(root.Right)
	}
}

func postOrder(root *Node) {
	if root != nil {
		preOrder(root.Left)
		preOrder(root.Right)
		fmt.Println(root.Val)
	}
}

func leftCount(root *Node) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		return leftCount(root.Left) + leftCount(root.Right)
	}
}

func depth(root *Node) int {

	if root == nil {
		return 0
	} else {
		depthLeft := depth(root.Left);
		depthRight := depth(root.Right)
		if depthLeft > depthRight {
			return depthLeft + 1
		} else {
			return depthRight + 1
		}
	}
}

//1) Create an empty stack S.
//2) Initialize current node as root
//3) Push the current node to S and set current = current->left until current is NULL
//4) If current is NULL and stack is not empty then
//a) Pop the top item from stack.
//b) Print the popped item, set current = popped_item->right
//c) Go to step 3.
//5) If current is NULL and stack is empty then we are done.
func preorderStack(root *Node) []int {

	var res = []int{}
	if root == nil {
		return res
	}
	var stack = []*Node{}
	for root != nil || len(stack) != 0 {
		if root != nil {
			res = append(res,root.Val)
			stack = append(stack,root)
			root = root.Left

		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}

func inorderStack(root *Node) []int  {
	var res = []int{}
	if root == nil {
		return res
	}
	var stack = []*Node{}
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack,root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			res = append(res,root.Val)
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}

// 左右根，就左右孩子都入栈，为了让结果是先左后右，入栈时要反过来，先右入栈，后左入栈
func postorderStack(root *Node) []int  {
	var res = []int{}
	if root == nil {
		return res
	}
	var pre = &Node{}
	var stack = []*Node{}
	stack = append(stack,root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if (cur.Left == nil && cur.Right == nil) || (pre != nil && (pre == cur.Left || pre == cur.Right)) {
			res = append(res,cur.Val)
			pre = cur
			stack = stack[:len(stack) -1]
		} else {
			if cur.Right != nil {
				stack = append(stack,cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack,cur.Left)
			}
		}
	}
	return res
}

// 二维数组是必然的
// 本质上是 BFS 的一种应用，可以总结出 BFS 的模板，借助队列来实现
func levelOrder(root *Node) [][]int {
	levelOrder := [][]int{}
	if root == nil {
		return levelOrder
	}
	queue := []*Node{}
	queue = append(queue,root)
	for len(queue) > 0 {
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level,node.Val)
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
		}
		levelOrder = append(levelOrder,level)

	}
	return levelOrder
}

//func BFS(root *Node) {
//	visited :=
//}

//------------------------------------------------

// 冒泡
func bubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j:= 0; j < len(arr) - 1; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}

	return arr
}

// 插入排序
// 遍历后面的，一个一个往前比较插入有序集合
func insertSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		for j := i; j >= 1; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1],arr[j] = arr[j],arr[j-1]
			} else {
				// 因为是往有序集合里从右往左依次比较，所以只要一个不满足，就可以break了
				break
			}
		}
	}
	return arr

}

// 选择排序
// 遍历选出最小放第一个，选出第二小放第二个
//func selectSort(arr []int) []int  {
//
//}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr)/2
	return merge(mergeSort(arr[:mid]),mergeSort(arr[mid:]))

}

func merge(left []int,right []int) []int {
	tmp := make([]int,0)
	i, j := 0,0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp,left[i])
			i++
		} else {
			tmp = append(tmp,right[j])
			j++
		}
	}

	if i < len(left) {
		tmp = append(tmp,left[i:]...)
	}
	if j < len(right) {
		tmp = append(tmp,right[j:]...)
	}
	return tmp
}


