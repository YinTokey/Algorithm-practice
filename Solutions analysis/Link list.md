# 链表题
##### 1.给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null。
这题基本上基于数学分析，代码实现逻辑反而很简单。解析直接看这个
https://www.nowcoder.com/questionTerminal/253d2c59ec3e4bc68da16833f79a38e4
看ID为`求一个大大的offer`的解释，这里把他的实现翻译成自己熟悉的Python
时间复杂度O(n)，空间复杂度O(1)
```
class Solution:
    def EntryNodeOfLoop(self, pHead):
        if pHead == None or pHead.next == None or pHead.next.next == None:
            return None
        fast = pHead.next.next
        slow = pHead.next
        # 第一个循环判断是否有环
        while fast != slow:
            if fast.next != None and fast.next.next != None:
                fast = fast.next.next
                slow = slow.next
            else:
                return None
        fast = pHead
        # 第二个循环找到入口节点
        while fast != slow:
            fast = fast.next
            slow = slow.next
        return slow
```

go版本求环形链表交叉节点，也可以用这个判断链表是否有环
```
func detectCycle(head *ListNode) *ListNode {

    if head == nil || head.Next == nil {
        return nil
    }
    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    if fast == nil || fast.Next == nil {
        return nil
    }
    fast = head
    for fast != slow {
        fast = fast.Next
        slow = slow.Next
    }
    return slow
}

```


##### 2.反转链表
牛客网就有分析，具体不贴了，直接贴AC代码，画个图很好理解。总结一下就是temp保存next,指向last,
```
    def ReverseList(self, pHead):
        if pHead == None or pHead.next == None:
            return pHead
        last = None
        while pHead:
            tmp = pHead.next
            pHead.next = last
            last = pHead
            pHead = tmp
        return last
```

go语言的解法，和上面差不多```
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var prev *ListNode
    for head != nil {
        tmp := head.Next
        head.Next = prev
        prev = head
        head = tmp
    }

    return prev
}
```
再补一个递归法，时间 空间 复杂度都为O(n)，因为递归用到栈，最大O（n）

思路：假设head后面的部分已经被翻转了，head的下一个节点必须反过来指向它，head自己的next先断开。
即 head.Next.Next = head ,   head.Next = nil

```
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var p *ListNode = reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil

    return p
}

```

##### 3.两个链表第一个公共节点
解析 https://www.nowcoder.com/questionTerminal/6ab1d9a29e88450685099d45c9e31e46
核心思路就是全部入栈，然后再依次出栈对比数值，如果不一样，那么久找到了分叉点
```
    def FindFirstCommonNode(self, pHead1, pHead2):
        if pHead1 == None or pHead2 == None:
            return None
        stack1 = []
        stack2 = []
        result = None
        while pHead1:
            stack1.append(pHead1)
            pHead1 = pHead1.next
        while pHead2:
            stack2.append(pHead2)
            pHead2 = pHead2.next
        while stack1 and stack2:
            top1 = stack1.pop()
            top2 = stack2.pop()
            if top1 == top2:
                result = top1
            else:
                break
        return result
```

**解法2**
统计两个链表的长度，以及二者的差值。假设差值为N，那么长的那条先走N步，然后再一起走，一旦发现相同，就是第一个公共节点。这种方式可以达到空间复杂度O(1)
```
    def FindFirstCommonNode(self, pHead1, pHead2):
        if pHead1 == None or pHead2 == None:
            return None
        length1,length2,step = 0,0,0
        p1,p2 = pHead1,pHead2
        while p1:
            length1 +=1
            p1 = p1.next
        while p2:
            length2 +=1
            p2 = p2.next
        step = abs(length2-length1)
        if length2 > length1:
            while step != 0:
                step -= 1
                pHead2 = pHead2.next
        else:
            while step != 0:
                step -= 1
                pHead1 = pHead1.next
        while pHead1 and pHead2:
            if pHead1.val == pHead2.val:
                return pHead1
            pHead1 = pHead1.next
            pHead2 = pHead2.next
        return None
```

go语言解法，O（n）时间复杂度，O(1) 空间复杂度
```

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    l1 := 1
    l2 := 1
    tmp1 := headA
    tmp2 := headB
    for tmp1 != nil  {
        tmp1 = tmp1.Next
        l1 += 1
    }

    for tmp2 != nil  {
        tmp2 = tmp2.Next
        l2 += 1
    }

    if l2 > l1 {
        offset := l2 - l1
        for offset > 0 {
            headB = headB.Next
            offset -= 1
        }
    } else if l2 < l1 {
        offset := l1 - l2
        for offset > 0 {
            headA = headA.Next
            offset -= 1
        }
    }

    // 一起走
    for headA != nil {

        if headA == headB {
            return headA
        } else {
            headA = headA.Next
            headB = headB.Next
        }

    }

    return nil
}

```



##### 4.删除链表重复节点
有序链表：就是使用两个指针操作。这题很巧妙的地方在于建一个辅助空节点，指向第一个节点，这么做的目的是为了解决第一个节点就是重复节点的情况，这种情况如果不建辅助节点，那么第一个节点删不掉。
```
   def deleteDuplication(self, pHead):
        if pHead == None or pHead.next == None:
            return pHead
        head = ListNode(0)
        head.next = pHead
        pre = head
        last = head.next
        while last:
            if last.next != None and last.val == last.next.val:
                while last.next != None and last.val == last.next.val:
                    last = last.next
                pre.next = last.next
                last = last.next
            else:
                pre = pre.next
                last = last.next
            
        return head.next
```

`无序链表` 会比较麻烦。有两种解法，方法一利用哈希表，时间复杂度O(n) , 空间复杂度O(n)。方法二不使用额外空间，但是时间复杂度会增加到O(n^2)。下面分别用Go实现
```
方法1：有额外空间
func removeDuplicateNodes(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var nodemap map[int]bool
    // map 记得要初始化
    nodemap = make(map[int]bool)
    nodemap[head.Val] = true
    pos := head
    for pos.Next != nil {
        cur := pos.Next
        if nodemap[cur.Val] == false {
            nodemap[cur.Val] = true
            pos = pos.Next
        } else {
            pos.Next = pos.Next.Next
        }
    }
    
    return head
}

```
```
方法2 无额外空间。头部保留一个节点，然后遍历删除后续和该节点值相同的节点。两重循环，达到目的。
二重循环，在leetcode上运行时间，比方法1慢了20几倍

func removeDuplicateNodes(head *ListNode) *ListNode {
    p1 := head
    for p1 != nil {
        p2 := p1
        for p2.Next != nil {
            if p2.Next.Val == p1.Val {
                p2.Next = p2.Next.Next
            } else {
                p2 = p2.Next
            }
        }
        p1 = p1.Next
    }
    return head
}


```

##### 5.单链表排序
要求常数级空间复杂度
https://leetcode.com/problems/sort-list/discuss/46711/Python-merge-sort-with-comments.
```
# merge sort, recursively 
def sortList(self, head):
    if not head or not head.next:
        return head
    # divide list into two parts
    fast, slow = head.next, head
    while fast and fast.next:
        fast = fast.next.next
        slow = slow.next
    second = slow.next
    # cut down the first part
    slow.next = None
    l = self.sortList(head)
    r = self.sortList(second)
    return self.merge(l, r)
    
# merge in-place without dummy node        
def merge(self, l, r):
    if not l or not r:
        return l or r
    if l.val > r.val:
        l, r = r, l
    # get the return node "head"
    head = pre = l
    l = l.next
    while l and r:
        if l.val < r.val:
            l = l.next
        else:
            nxt = pre.next
            pre.next = r
            tmp = r.next
            r.next = nxt
            r = tmp
        pre = pre.next
    # l and r at least one is None
    pre.next = l or r
    return head

```
##### 6. 旋转链表
https://leetcode.com/problems/rotate-list/description/
给定一个链表，将链表向右旋转k个位置
思路：统计长度，求余算出实际要走多少步。但是没必要真的走，构建出循环链表，指针走到要返回的节点上面，然后断开前一个节点连接，使它恢复到单向链表。最后返回那个节点即可。
```
def rotateRight(self, head, k):
   """
   :type head: ListNode
   :type k: int
   :rtype: ListNode
   """
   if head is None or head.next is None:
       return head
   length = 0
   dummy = ListNode(0)
   dummy.next = head
   p = dummy
   while p.next:
       length += 1
       p = p.next
   step = length - (k % length)
   p.next = dummy.next

   while step > 0:
       p = p.next
       step -= 1
   head = p.next
   p.next = None
   return head
```

##### 7. 倒数第K个节点
经典的双指针解决方案
https://www.nowcoder.com/practice/529d3ae5a407492994ad2a246518148a?tpId=13&tqId=11167&tPage=1&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking
```
    def FindKthToTail(self, head, k):
        if head == None or k < 1:
            return None
        ptr1 = head
        ptr2 = head
        while k > 1:
            if ptr1.next != None:
                ptr1 = ptr1.next
                k -= 1
            else:
                return None
        while ptr1.next:
            ptr1 = ptr1.next
            ptr2 = ptr2.next
        return ptr2
```

##### 8.合并两个排序链表
基本没难度
```
    def Merge(self, pHead1, pHead2):
        dummy = ListNode(0)
        head = dummy
        while pHead1 and pHead2:
            if pHead1.val >= pHead2.val:
                dummy.next = pHead2
                pHead2 = pHead2.next
            else:
                dummy.next = pHead1
                pHead1 = pHead1.next
            dummy = dummy.next
        if pHead1:
            dummy.next = pHead1
        if pHead2:
            dummy.next = pHead2
            
        return head.next
```
go 语言实现，注意构建空节点
```
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var ptr1 *ListNode = &ListNode{Val:0,Next:nil}
    head := ptr1

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            ptr1.Next = l1
            l1 = l1.Next
        } else {
            ptr1.Next = l2
            l2 = l2.Next
        }
        ptr1 = ptr1.Next
    }

    if l1 != nil {
        ptr1.Next = l1
    }

    if l2 != nil {
        ptr1.Next = l2
    }

    return head.Next
}

```

##### 9 合并多个排序链表
上一题的升级版
```
https://leetcode.com/problems/merge-k-sorted-lists/

def mergeLists(lists):
	if not lists:
		return None
	if len(lists) == 1:
		return lists[0]
	mid = len(lists)//2
	left = mergeList(lists[:mid])
	right = mergeList(lists[mid:])
	return merge(left,right)
```


##### 10 实现memcpy
```
void *memcpy(void *dest,const void *src,size_t n)
{
	char *d;
	const char *s;
	if(dest > (src+count) || dest < src)
	{
		d = dest
		s = src
		while(n--){
			*d ++ = *s++;
		}
	}
	else//内存重叠
	{
		d = char*(dest+count-1)
		s = char*(src+count-1)
		while(n--){
			*d-- = *s--;
		}
	}
	return d

}
```

##### 有序链表转二叉搜索树
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树

如果有数组就好办，类似这题https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

这里虽然可以构建有序数组，但是会增加额空间O(n)。如果不适应数组，那么可以每次都遍历链表，找到中间节点，然后递归构建。时间复杂度O(nLogn)

##### 求链表中间节点
最朴素的方法，遍历一遍，求链表总节点个数，然后第二次遍历定位目标节点。
好一点的方法，使用快慢指针，第一个指针一次走一步，第二个指针一次走两步。当第二个指针走到尾部，就可以根据第一个指针得出结果。
下面是Go的解法
```
func middleNode(head *ListNode) *ListNode {
    front := head

    for front != nil && front.Next != nil {
        front = front.Next.Next
        head = head.Next
    }

    return head

}

```



