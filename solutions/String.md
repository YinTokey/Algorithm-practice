# 字符串
##### 手动翻转字符串
大部分与语言其实都提供翻转字符串的函数支持，特别是python这种。如果要用最朴素的方式的手动翻转的话，其实核心思路就是遍历一半，然后第一个字符和最后一个交换，第二个和倒数第二个交换，一直下去。
下面这个其实是翻转数组，但是字符串可以转换成数组的
```
def reverse(s,start,end):
	while start < end:
		s[start],s[end] = s[end],s[start]
		start += 1
		end -= 1
```
下面是Go的实现
```
func reverseString(s []byte)  {
    start, end := 0, len(s) - 1
    for start < end {
        s[start], s[end] = s[end], s[start]
        start++
        end--
    }
}

```

##### 翻转单词顺序列
https://www.nowcoder.com/practice/3194a4f4cf814f63919d0790578d51f3?tpId=13&tqId=11197&tPage=3&rp=3&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
最粗暴的方式就是直接利用函数，分割单词，再转置。
现在用一种比较朴素的方式来实现：先翻转整个字符串，然后翻转每个单词。两个指针表示一个单词的头尾，遍历时遇到空格，说明需要翻转这个单词。

```
  def ReverseSentence(self, s):
        if not s:
            return s
        s = list(s)
        self.reverse(s,0,len(s)-1)
        start,end = 0,0
        while start < len(s):
            if s[start] == ' ': # 首位空
                start +=1
                end += 1
            elif end == len(s) or s[end] == ' ': #到达翻转节点
                self.reverse(s,start,end -1)
                end +=1
                start = end
            else:  #常规移动
                end += 1
        return ''.join(s)
```

##### 字符串排列
题目 https://www.nowcoder.com/practice/fe6b651b66ae47d7acce78ffdd9a96c7?tpId=13&tqId=11180&tPage=2&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking
思路参考：https://blog.csdn.net/u010005281/article/details/79920046
```
def Permutation(self, ss):
   list = []
   if len(ss) <= 1:
       return ss
   for i in range(len(ss)):
       for j in map(lambda x: ss[i] + x,self.Permutation(ss[:i]+ss[i+1:])):
           if j not in list:
               list.append(j)
   return list
```

##### 第一个只出现一次的字符
因为强调第一个，又要返回它的位置。字典又是无序的，但是又想用字典的特性，所以就把数组稍微整整，假装成字典来用，满足有序性。

https://www.nowcoder.com/practice/1c82e8cf713b4bbeb2a5b31cf5b0417c?tpId=13&tqId=11187&tPage=2&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking

```
    def FirstNotRepeatingChar(self, s):
        if s == None or len(s) <= 0:
            return -1
        alist = list(s)
        blist = {}
        for i in alist:
            if i not in blist:
                blist[i] = 1
            else:
                blist[i] += 1
        for index,i in enumerate(s):
            if blist[i] == 1:
                return index
        return -1
```

##### 左旋转字符串
直接把原字符串扩充为2倍，然后切割。
https://www.nowcoder.com/practice/12d959b108cb42b1ab72cef4d36af5ec?tpId=13&tqId=11196&tPage=3&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking

```

```

##### 字符串转换成整数
https://www.nowcoder.com/practice/1277c681251b4372bdef344468e4f26e?tpId=13&tqId=11202&tPage=3&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking

思路就是构建字典，然后做合法性匹配。

```
    def StrToInt(self, s):
        if not s:
            return 0
        num = []
        numbers = {'1':1,'2':2,'3':3,'4':4,'5':5,'6':6,'7':7,'8':8,'9':9}
        for i in s:
            if i in numbers.keys():
                num.append(numbers[i])
            elif i == '-' or i == '+':
                continue
            else:
                return 0
        ans = 0
        for j in num:
            ans = ans * 10 + j
        if s[0] == '-':
            ans = 0 - ans
        return ans
```

##### 表示数值的字符串
这题没有那么难，只要对合法性条件做充分考虑就可以了。
https://www.nowcoder.com/practice/6f8c901d091949a5837e24bb82a731f2?tpId=13&tqId=11206&tPage=3&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking

合法性条件有哪些？
以e为界，e后面不能有.或者为空。然后利用数组切片把e的前后两部分放到一个判断函数里。
在判断函数内部，+- 不能出现在首位。字符串里面的 . 不能出现超过1次。

总和上面条件，就可以对字符串的合法性做处理了。
```
    def isNumeric(self, s):
        if not s:
            return False
        alist = [i.lower() for i in s]
        if 'e' in alist:
            index = alist.index('e')
            front = alist[:index]
            behind = alist[index+1:]
            if '.' in behind or len(behind) == 0:
                return False
            return self.Digit(front) and self.Digit(behind)
            
        else:
            return self.Digit(alist)
        
        
    def Digit(self,alist):
        dotnum = 0
        allownum = ['0','1','2','3','4','5','6','7','8','9','+','-','.']
        for i in range(len(alist)):
            if alist[i] not in allownum:
                return False
            if alist[i] == '.':
                dotnum += 1
            if alist[i] in '+-' and i != 0:
                return False
        if dotnum > 1:
            return False
        return True
```

##### 字符流中第一个不重复的字符
原题：
https://www.nowcoder.com/practice/00de97733b8e4f97a3fb5c680ee10720?tpId=13&tqId=11207&tPage=3&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking

好好理解一下这个思路，写起来并不难

https://github.com/leeguandong/Interview-code-practice-python/blob/master/%E5%89%91%E6%8C%87offer/%E5%AD%97%E7%AC%A6%E6%B5%81%E4%B8%AD%E7%AC%AC%E4%B8%80%E4%B8%AA%E4%B8%8D%E9%87%8D%E5%A4%8D%E7%9A%84%E5%AD%97%E7%AC%A6.py

```
    def __init__(self):
        self.alist = []
        self.adict = {}
    # 返回对应char
    def FirstAppearingOnce(self):
        while len(self.alist) > 0 and self.adict[self.alist[0]] == 2:
            self.alist.pop(0)
        if len(self.alist) == 0:
            return '#'
        else:
            return self.alist[0]
        
    def Insert(self, char):
        if char not in self.adict.keys():
            self.adict[char] = 1
            self.alist.append(char)
        else:
            self.adict[char] = 2
```
##### 罗马数字转整数
原题：
https://leetcode.com/problems/roman-to-integer/description/

思路是反向遍历，除了最后一个以外，其他位如果前一位大于后一位，就累加，如果小于后一位，就减它

```
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        if not s:
            return 0
        
        dict = {'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
        result = 0
        for i in range(len(s)-1,-1,-1):
            if i!= len(s)-1:
                if dict[s[i]] < dict[s[i+1]]:
                    result -= dict[s[i]]
                else:
                    result += dict[s[i]]
            else:
                result += dict[s[i]]
        return result
```

##### 最长不重复子序列
https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/
参考思路
https://blog.csdn.net/yurenguowang/article/details/77839381

时间和空间复杂度都为O(n)

##### Implement strStr()
leetcode 28 题， 字符串匹配。
最简单的解法是，遍历第一个字符串，使用字典存储索引。再次遍历第二个字符串，使用字典键值对做对比。但是这样做会有O(n)的空间复杂度。这里其实考察的是KMP算法。


