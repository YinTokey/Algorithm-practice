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

