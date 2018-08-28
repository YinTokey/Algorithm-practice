# -*- coding:utf8 -*-
# 直接根据https://github.com/knightsj/awesome-algorithm-question-solution中数组的顺序

# 4.1数组中是否包含重复元素。只要有包含重复元素，就返回true，否则就返回false
def duplicated(array):
    dict = {}
    for num in array:
        if dict.has_key(num):
            dict[num] += 1
        else:
            dict[num] = 1
    for value in dict.itervalues():
        if value > 1:
            return True
    return  False


array1 = [1, 2, 5, 6, 8, 11, 323, 4]
# print(duplicated(array1))


# 4.2 出现次数超过数组长度一半的元素
# 这种题没有意义,如果没有特殊要求，还是像上面那样字典解决

# 4.3 数组中只出现过一次的数字
# Given a **non-empty** array of integers, every element appears *twice* except for one. Find that single one.
# 题目本身不难，但是它有额外的要求 "线性时间复杂度，不使用额外空间"，所以还是值得一做。66题书的第40题简化版
# 记住成对的数字在异或运算中会被抵消，最终剩下单个数字  -------
def singleOne(array):
    result = 0
    for n in array:
        result ^= n
        # print(result)
    return result
array3 = [15,2,2,3,5,1,1,5,15,8,8]
# print(singleOne(array3))

# 4.4 寻找数组中缺失的数字
# Given an array containing *n* distinct numbers taken from `0, 1, 2, ..., n`, find the one that is missing from the array.
# 这是方差为1的等差数列，直接用公式理论求和，然后依次减掉数组中的每一个数，最终剩下的结果，就是缺失的数了。
def findLackNum(array):
    n = len(array)
    total = (n+1)*n/2
    for i in array:
        total -= i
    return total
array4 = [9,6,4,2,3,5,7,0,1]
# print(findLackNum(array4))

# 4.5 将所有的0移动到数组末尾
# 要求非0的元素顺序不变，必须原地处理，不能使用辅助数组,使用尽可能少的操作。也就是时间复杂度一般要O(n). 一般这种"原地处理"的要求，就要依赖指针了。
# 这种思路比较精巧，基本还是靠熟练度了
def moveZeroEnd(array):
    k = 0
    for i in range(len(array)):
        if array[i] != 0:
            array[k],array[i] = array[i],array[k]
            k += 1

array5 =  [9,0,4,2,0,5,7,0,1]
moveZeroEnd(array5)
#print(array5)

# 4.6 移除数组中等于某个值的元素， 返回新的数组长度
# 要求原地解决，空间复杂度为0(1), 其实和上面那题没什么差别
def remvoeSome(array,val):
    k = 0
    for i in range(len(array)):
        if array[i] != 0:
            array[k],array[i] = array[i],array[k]
            k += 1
    return k

#print(remvoeSome(array5,0))

# 4.7 三色旗问题
# Given an array with *n* objects colored red, white or blue, sort them **in-place **so that objects of the same color are adjacent, with the colors in the order red, white and blue.
#
# Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
# 要求： 空间复杂度O(1)， 时间复杂度为一趟遍历
# 这种原地操作的，基本就等同于指针操作
def threeColor(array):
    low,mid = 0,0
    hight = len(array)-1

    while mid<=hight:
        if array[mid] == 0:
            array[low],array[mid] = array[mid],array[low]
            low +=1
            mid +=1
        elif array[mid] == 1:
            mid +=1
        else:
            array[mid],array[hight] = array[hight],array[mid]
            hight-=1


array7 = [0,2,2,2,0,1,1,0,2]
threeColor(array7)
#print(array7)

# 4.8 有序数组内部的两个元素的和为目标值
# 这题很像动态规划的画表格思路，考虑到重复问题，建立二维数组后，只需要遍历上三角
def twoSum(array,val):

    result = []
    for i in range(len(array)-1):
        k = i+1
        for k in range(len(array)):
            if array[i]+array[k] == val:
                result.append(i)
                result.append(k)
                return result

    return []

array8 = [31,2,7,11,5]
val = 9
# print(twoSum(array8,9))

# 4.9 无序数组和大于或等于某值的最小子数组
# Given an array of **n** positive integers and a positive integer **s**,
#  find the minimal length of a **contiguous** subarray of which the sum ≥ **s**. If there isn't one, return 0 instead.
# 其实这题按照动态规划建表的思路也是可以做的，但是时间复杂度比较大，尝试滑动窗口, 它这里求的是子序列的长度，而不是子序列本身
# 这题和 "和至少为K的子数组没什么差别"
def minSubArray(nums,val):
    if nums == None:
        return 0
    start,end,sum,min = 0,0,0,


nums = [2,7,5,11]
val = 14
# print(minSubArray(nums,val))

# 4.10 两数组的交点元素

