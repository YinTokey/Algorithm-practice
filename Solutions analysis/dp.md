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
利用滚动一维数组，每次都存储与更新最优解，可以把空间复杂度优化到O(n)
```
func uniquePaths(m int, n int) int {

    opt := make([]int,n)

    for i:=0; i < m ; i++ {
        for j :=0; j< n; j++ {
            if i == 0 || j == 0 {
                opt[j] = 1
            } else {
                opt[j] = opt[j-1] + opt[j]
            }
        }
    }
    return opt[n-1]
}

```
##### 63. 不同路径2
加了障碍物，需要增加一些判断。且画图领会。
```
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    n, m := len(obstacleGrid), len(obstacleGrid[0])
    opt := make([]int,m)

    if obstacleGrid[0][0] == 0 {
        opt[0] = 1
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if obstacleGrid[i][j] == 1 {
                opt[j] = 0
                continue
            } 

            if j >= 1 && obstacleGrid[i][j-1] == 0 {
                opt[j] += opt[j-1]
            }
         }
    }

    return opt[m-1]
}

```
##### 1143 最长公共子序列
参照js版本的题解 https://juejin.im/post/6844903613861462029
下面是go实现
```
func longestCommonSubsequence(text1 string, text2 string) int {
    n, m := len(text1), len(text2)

    dp := make([][]int,n+1)
    
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m+1)
        for j := 0; j <= m; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = 0
                continue
            }  

            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j],dp[i][j-1])
            }
            
        }
    }
    return dp[n][m]
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

```
由于表格存在左上到右下递增趋势，可以对空间复杂度优化到O(n)。参照国际站Java最优解
https://www.youtube.com/watch?v=DuikFLPt8WQ&feature=emb_title
```
func longestCommonSubsequence(text1 string, text2 string) int {
    n, m := len(text1), len(text2)
    dp := make([]int, m+1)

    for i := 1; i <= n; i++ {
        prev := 0
        for j := 1; j <= m; j++ {
            tmp := dp[j]
            if text1[i-1] == text2[j-1] {
                dp[j] = prev + 1
            } else {
                dp[j] = max(dp[j],dp[j-1])
            }
            prev = tmp
        }
    }
    return dp[m]
}

```


##### 121  买卖股票的最佳时机 1
这题比较简单，只允许买卖一次。可以直接用贪心算法。如果是变种题，则要考虑使用DP
使用一个变量记录历史最低点，最低点买入后，后面就考虑什么时候卖收益最高即可。
时间复杂度O(n), 空间复杂度O(1)
```
import "math"

func maxProfit(prices []int) int {
    var minPrice, maxProfit = math.MaxInt64, 0

    for _, value := range prices {
        if value < minPrice {
            minPrice = value
        } else if value - minPrice > maxProfit {
            maxProfit = value - minPrice
        }
    }
    return maxProfit
}

```
##### 122 买卖股票最佳时机 2
其实这题也是贪心，只是可以多次买卖。那么直接考虑今天只要大于昨天，就卖。
```
func maxProfit(prices []int) int {
    re := 0

    for idx, value := range prices {
        if idx < len(prices) - 1 {
            if prices[idx+1] > value {
                re += prices[idx+1] - value
            }
        }
    }
    return re
}

```

##### 123. 买卖股票的最佳时机 III
股票题通用解答 https://leetcode-cn.com/circle/article/qiAgHn/
这题只能最多两次买卖，需要用到一维动态规划



##### 338. 比特位计数
一般dp是i和i-1的关系。这题也可以，但是因为这题是二进制，所以可以考虑 i和i/2的状态转移方程。
可以得到下面关系 f(i) = f(i/2)+ i%2   (即区分奇偶)，这种题目主要还是靠熟练度。

```
func countBits(num int) []int {
    resutl := make([]int,num+1)
    //第一个为0，从1开始遍历
    // f(i) = f(i/2)+ i%2
    for i := 1; i <= num; i++ {
        resutl[i] = resutl[i>>1] + (i & 1)
    }
    return resutl
}


```


