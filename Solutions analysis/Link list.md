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





