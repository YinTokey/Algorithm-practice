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

