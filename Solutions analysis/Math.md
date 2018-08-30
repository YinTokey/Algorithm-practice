# 数学题
##### 不用加减乘除做加法
剑指offer第6章。python位运算比较坑，这题可以用JavaScript来解决
总结起来就是异或运算，与运算，位移。然后再一直循环。
```
function Add(num1, num2)
{
    var sum, carry;
    while(num2 != 0){
        sum = num1 ^ num2;
        carry = (num1 & num2) << 1;
        num1 = sum;
        num2 = carry;
    }
    return num1;
}
```
##### 孩子们的游戏
参照牛客网的剑指offer专题。核心思路是建立循环链表，然后在循环链表中进行删除节点的操作。也就是约瑟夫环。这里提供一个思路，没能全部通过，应该只是语法问题
```
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def LastRemaining_Solution(self, n, m):
        head = ListNode(0)
        first = head
        for i in range(1,n):
            listHead = ListNode(i)
            head.next = listHead
            head = head.next
        head.next = first
        while first.next is not None:
            for i in range(m-1):
                first = first.next
            tmp = first.next
            first = first.next.next
            tmp.next = None
        return first.val
```


