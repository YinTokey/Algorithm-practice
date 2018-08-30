# 数组题
##### 和为S的两个数字
建立矩阵，做上三角遍历，这个套路在动态规划中很常见，虽然效率不太高，但算是一个万金油的套路。
比如数组 1，3，5，6，8 。 tsum = 1。建表

 | 1 | 3 | 5 | 6 | 8
------- | ------- | ------- | ------- | ------- | -------
1 | x | 4 | 6 | 7 | 9
3 | x | x | 8 | 9 | 11
5 | x | x | x | 11| 13
6 | x | x | x | x | 14
8 | x | x | x | x | x

表中数据表示求和，因为数组中相同元素不会互相加，所以只需要遍历上三角.找到满足条件的组合，存起来。如果超过1对组合，在进行对比筛选，保证保存组合的临时变量永远是最优解即可。


```
# -*- coding:utf-8 -*-
class Solution:
    def FindNumbersWithSum(self, array, tsum):
        if array == None:
            return None
        n = len(array)
        ret = []
        q = 0
        for i in range(n-1):
            
            for j in range(i+1,n):
                if array[i] + array[j] == tsum:
                    if q != 0 and array[i] * array[j] < q :
                        q = array[i] * array[j]
                        ret.pop()
                        ret.pop()
                        ret.append(array[i])
                        ret.append(array[j])
                    if q == 0:
                        q = array[i] * array[j]
                        ret.append(array[i])
                        ret.append(array[j])
        return ret            
```




