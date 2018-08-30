
class ListNode(object):
  def __init__(self, x):
    self.val = x
    self.next = None


n = 15
m = 6

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
    tmp = first
    first = first.next.next
    tmp.next = None

print(first.val)