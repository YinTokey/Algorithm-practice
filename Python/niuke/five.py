# -*- coding:utf8 -*-
import heapq

# 连续子数组的最大和
# 输入一个整型数组，数组里有正数也有负数。数组中一个或连续的多个整数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为O(n)。
# 动态规划，画表格解决
# tip 起始这题有个隐藏条件：就是里面既然有整数也有负数，那么最大值必然是大于0的，因为不可能全取负数。所以可以在计算的时候，当累积结果为负数，就舍弃
def array_max_sum(array):
    ret = 0
    if not array:
        return ret

    current = 0
    for i in array:
        if current <= 0: # 当累积结果为负数，就舍弃
            current = i
        else:
            current += i
        ret = max(ret,current)
    return ret

array = [-1,2,3,-4,6]
# print(array_max_sum(array))


# 求从1到n整数的十进制表示中，1出现的次数
# 递归求解  f(n) = 9 * f(n-1) + 10^(n-1)
def get_digits(n):
    # 求整数n的位数
    ret = 0
    while n:
        ret += 1
        n /= 10
    return ret

def get_1_digits(n):
    """
    获取每个位数之间1的总数
    :param n: 位数
    """
    if n <= 0:
        return 0
    if n == 1:
        return 1
    current = 9 * get_1_digits(n-1) + 10 ** (n-1)
    return get_1_digits(n-1) + current

def get_1_nums(n):
    if n < 10:
        return 1 if n >= 1 else 0
    digit = get_digits(n)  # 位数
    low_nums = get_1_digits(digit-1)  # 最高位之前的1的个数
    high = int(str(n)[0])  # 最高位
    low = n - high * 10 ** (digit-1)  # 低位

    if high == 1:
        high_nums = low + 1  # 最高位上1的个数
        all_nums = high_nums
    else:
        high_nums = 10 ** (digit - 1)
        all_nums = high_nums + low_nums * (high - 1)  # 最高位大于1的话，统计每个多位数后面包含的1
    return low_nums + all_nums + get_1_nums(low)


# print(get_1_nums(n))

# 统计数组中出现次数超过一半的数字
def half_num(nums):
    hashes = dict()
    half_length = len(nums)/2

    for i in nums:
        if hashes.has_key(i):
            hashes[i] += 1
        else:
            hashes[i] = 1

        if hashes[i] > half_length:
            return i

arr = [6,6,1,1,6,3,6,6,9,0,6,6,6]
# print(half_num(arr))


# 求数组中的最小k个数 ， 比如 【-1，7，3，4，5，2，6】，求最小4个数为 [-1,2,3,4]。 最粗暴的方法是先排序再取k个数。
# 比较优化的方法，
# 使用快排中的partition思想。

# ①我们设定partition函数的哨兵为key=lists[left]，在partition函数中完成一轮比较的结果是，比key大的数都在其右边，比key小的数放在其左边。完成该轮后返回其left=right时left的值。

# ②我们判断left的值是比k大还是小：

# 如果left的值比k大，说明上轮partition之后，lists中前left个小的数在左边，其余的数在其右边，我们还需要把寻找范围缩小，下次找的时候只在数组前面left个数中找了。

# 如果left的值比k小，说明上轮partition之后，前left个数找的太少了，我们需要再往数组的后面找。

def latest_k(array,k):
    if len(array) < k:
        return []
    ret = []
    for i in array:
        heapq.heappush(ret)

def partion(lists,left,right):
    key = lists[left]
    while left < right:
        while left < right and lists[right] >= key:
            right -= 1
        lists[left] = lists[right]
        while left < right and lists[left] <= key:
            left += 1
        lists[right] = lists[left]
    lists[right] = key
    return left


def function2(lists,k):
    #划分法主要函数部分
    length = len(lists)
    left = 0
    right = length - 1
    index = partion(lists,left,right)
    while k!=index:
        if index > k-1:
            right = index-1
        else:
            left = index+1
        index = partion(lists,left,right)
    return lists[0:k]

# 有点类似二分查找，只是不是均分

arr = [1,2,9,-9,10,5]
print(function2(arr,4))
