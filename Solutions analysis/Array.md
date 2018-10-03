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

##### 旋转数组的最小数字
https://www.nowcoder.com/practice/9f3231a991af4f55b95579b44b7a01ba?tpId=13&tqId=11159&tPage=1&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking

题目说的有点模糊，本质是查找数组中最小的数字。由于是旋转数组，只要找到分割点，把数组分割成两部分，那么这两部分是有序的。并且它是左选择，最小数字存在于分割点左边。

```
    def minNumberInRotateArray(self, rotateArray):
        if len(rotateArray) == 0:
            return 0
        elif len(rotateArray) == 1:
            return rotateArray[0]
        else:
            low = 0
            high = len(rotateArray) - 1
            while low < high:
                mid = int((low + high)/2)
                if rotateArray[mid] < rotateArray[high]:
                    high = mid
                elif rotateArray[mid] > rotateArray[high]:
                    low = mid+1
                else:
                    high = high - 1
            return rotateArray[high]
```

##### 调整数组顺序，使得奇数位于偶数前面
https://www.nowcoder.com/practice/beb5aa231adc45b2a5dcc5b62c93f593?tpId=13&tqId=11166&tPage=1&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking

这种题直接开两个数组解决。


```
    def reOrderArray(self, array):
        res1 = []
        res2 = []
        for i in array:
            if i % 2 == 0:
                res2.append(i)
            else:
                res1.append(i)
        return res1+res2
```

