'''
题目：输入一个链表，从尾到头打印链表每个节点的值

思路：遍历链表，一个一个，从头部插入数组

'''

# -*- coding:utf-8 -*-
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


def printListFromTailToHead(self, listNode):
    if not listNode:
        return []

    result = []
    head = listNode
    while head:
        result.insert(0,head.val)
        head = head.next
    return result
