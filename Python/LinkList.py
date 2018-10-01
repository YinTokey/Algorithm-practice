
class ListNode(object):
  def __init__(self, x):
    self.val = x
    self.next = None

def reverseKNode(head,k):
    if k < 2:
      return head
    cur = head
    start,pre,next = None,None,None
    count = 1
    while cur:
        next = cur.next
        if count == k:
            start = head if pre is None else pre.next
            head = cur if pre is None else head
            resign(pre,start,cur,next)
            pre = start
            count = 0
        count +=1
        cur = next
    return head


def resign(left,start,end,right):
    pre = start
    cur = start.next
    next = None
    while cur:
        next = cur.next
        cur.next = pre
        pre = cur
        cur = next



# -----------------

n1 = ListNode(1)
n2 = ListNode(2)
n3 = ListNode(3)
n4 = ListNode(4)
n5 = ListNode(5)
n6 = ListNode(6)
n7 = ListNode(7)
n8 = ListNode(8)
n1.next = n2
n2.next = n3
n3.next = n4
n4.next = n5
n5.next = n6
n6.next = n7
n7.next = n8
n8.next = None

p = n1
while p:
  #  print(p.val,end="->")
    p = p.next


def quickSort(array,left,right):
    if left > right:
        return
    low = left
    hight = right
    key = array[low]
    while left < right:
        while left < right and array[right] > key:
            right -= 1
        array[left] = array[right]
        while left < right and array[left] <= key:
            left += 1
        array[right] = array[left]
    # 调整基准数
    array[right] = key
    quickSort(array,low,left -1)
    quickSort(array,left+1,hight)

arr = [2,5,7,3,1,4,9]
print(arr)
quickSort(arr,0,len(arr)-1)
print(arr)

a = int(5)
print(a//2)

a = "1fGH980"
list = [i.lower() for i in a]
print(list)

