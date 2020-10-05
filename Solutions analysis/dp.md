# 动态规划

##### 62. 不同路径
二维动态规划，可以通过画表格找到最有子结构。放上Go的实现。
```
func uniquePaths(m int, n int) int {

    opt := make([][]int,m)

    for i:=0; i < m ; i++ {
        opt[i] = make([]int,n)
        for j :=0; j< n; j++ {
            if i == 0 || j == 0 {
                opt[i][j] = 1
            } else {
                opt[i][j] = opt[i-1][j] + opt[i][j-1]
            }
        }
    }
    return opt[m-1][n-1]
}

```



