'''
A栈是入栈，B栈是出栈。入栈的时候，直接进入A。出栈的时候，判断B是否为空，如果不为空，B就直接Pop。 如果为空，A全部出栈，进入到B中

'''

class Solution:
    def __init__(self):
      self.stackA = []
      self.stackB = []

    def push(self, node):
          # write code here
      self.stackA.append(node)

    def pop(self):
        if self.stackB:
            self.stackB.pop()
        elif not self.stackA:
            return None
        else:
            while self.stackA:
                self.stackA.append(self.stackA.pop())
            return self.stackB.pop()