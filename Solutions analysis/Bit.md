# 位运算
##### 191. 位1的个数

```
func hammingWeight(num uint32) int {
    count := 0
    for num != 0 {
        if num & 1 != 0 {
            count++
        }
        num = num >> 1
    }
    return count
}

```

##### 231.   2的冥
满足条件  n & (n-1) == 0 。记住这个原则就简单了。原理归纳

2^x | n | n-1 | n & (n-1) |
------- | ------- | ------- | ----
2^0 | 0001 | 0000 | 0
2^1 | 0010 | 0001 | 0
2^2 | 0100 | 0011 | 0 
2^3 | 1000 | 0111 | 0

```
func isPowerOfTwo(n int) bool {
    
    return n > 0 && n & (n-1) == 0
}


```

##### 190 颠倒二进制位
```
func reverseBits(num uint32) uint32 {
    
    var ret uint32 = 0
    var power uint32 = 31

    for num != 0 {
        ret += (num & 1) << power
        num = num >> 1
        power -= 1
    }
    
    return ret
}
```
由于固定32位，所以即使有循环，时间复杂度也是O(1)

下面是另一种不循环的解法，使用分治代替循环。（待定）
```

```

##### 51 N皇后问题


##### 136. 只出现一次的数字
利用异或运算3个性质：1. 相同为0  2.任何数和0异或，还是它原来的值 3.无限多个数异或运算，满足结合律，即可以随意交换运算位置，顺序。因此全部进行异或运算，相同为0，两两互消，就留下最后一个了。

这里使用一个特殊技巧，即用数组第一个元素作为累计计算，这样连常数空间都不用分配了。不过正常来说，来时分配个变量保存结果合理一点。
```
func singleNumber(nums []int) int {
    
    for  i:= 1; i< len(nums); i++ {
        nums[0] ^= nums[i]
    }
    return nums[0]
}

```

