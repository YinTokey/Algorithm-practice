# 链表题
##### 1.给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null。
这题基本上基于数学分析，代码实现逻辑反而很简答。解析直接看这个
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
            if fast.next is not None and fast.next.next is not None:
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
牛客网就有分析，具体不贴了，直接贴AC代码，画个图很好理解
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


