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


