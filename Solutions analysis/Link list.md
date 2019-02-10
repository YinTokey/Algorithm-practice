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
统计两个链表的长度，以及二者的差值。假设差值为N，那么长的那条先走N步，然后再一起走，一旦发现相同，就是第一个公共节点。这种方式可以达到空间复杂度O(n)
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


##### 4.删除链表重复节点
就是使用两个指针操作。这题很巧妙的地方在于建一个辅助空节点，指向第一个节点，这么做的目的是为了解决第一个节点就是重复节点的情况，这种情况如果不建辅助节点，那么第一个节点删不掉。
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





