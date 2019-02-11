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
**O(n)空间复杂度**：
如果可以修改数组，且相对位置可以改变，那么使用快排的思路，可以达到常数级复杂度。
```
def reOrderArray(self, array):
    left = 0
    right = len(array)-1
    while left < right:
        while left < right and array[right]%2 == 0:
            right -=1
        while left < right and array[left]%2 != 0:
            left += 1
        array[left],array[right]=array[right],array[left]
    return array

```

##### 顺时针打印矩阵
https://www.nowcoder.com/practice/9b4c81a02cd34f76be2659fa0d54342a?tpId=13&tqId=11172&tPage=1&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking

模拟魔方旋转，打印第一行，然后切掉，逆时针旋转90度。一直循环。

```

```


##### top k 问题
比如最小K个数

使用heapq来快速进行堆操作。如果不想走捷径的话，可以使用快排，或者更加粗暴的冒泡排序
```
    def GetLeastNumbers_Solution(self, tinput, k):
        import heapq
        if tinput == None or len(tinput) < k or len(tinput) <= 0 or k <= 0:
            return []

        # 建立最小堆，最上面那个数是最小的，返回一个列表，这个列表就是从最小值开始的k个数
        return heapq.nsmallest(k, tinput)
```

top k 其实和选择排序的思路是一样的，可以使用选择排序。相对来说没有那种作弊的感觉。

```
    def GetLeastNumbers_Solution(self, tinput, k):
        res = []
        n = len(tinput)
        if k > n:
            return res
        count = 0
        for j in range(0,n-1):
            minIndex = j
            for i in range(j+1,n):
                if tinput[minIndex] > tinput[i]:
                    minIndex = i
            count +=1
            tinput[minIndex],tinput[j] = tinput[j],tinput[minIndex]
            if count >= k:
                break
        return tinput[:k]
        
```

##### 连续子数组最大和
https://www.nowcoder.com/practice/459bd355da1549fa8a49e350bf3df484?tpId=13&tqId=11183&tPage=2&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking

一定要记得一个隐藏条件：所求和的最小值是0，计算应该是以0位分节点。即使整个数组都是负数，空数组也可以作为它的子数组，和位0。记住这点，分析就比较容易了。它只要求值，不要求子数组，所以不用考虑滑动窗口或者动态规划。

这里采用的思路是用一个数值sum存储当前和，一个result用来保存最大值。
```
        if array == None or len(array) == 0:
            return 0
        result = array[0]
        sum = 0
        for i in range(len(array)):
            if sum <= 0:
                sum = array[i]
            else:
                sum += array[i]
            if sum > result:
                result = sum
        return result
```
##### 数组中重复的数字
https://www.nowcoder.com/practice/623a5ac0ea5b4e5f95552655361ae0a8?tpId=13&tqId=11203&tPage=2&rp=2&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking

###### tip1 如果可以修改元素
https://blog.csdn.net/u010579068/article/details/49514745
按照这个排序思路，时间复杂度O(n),空间复杂度O(1)




