# 求和问题

#### Two Sum
https://leetcode.com/problems/two-sum/description/
```
def twoSum(self,nums,target):
	inxDict = dict()
	for idx,num in enumerate(nums):
		if target - num in idxDict:
			return [idxDict[target-num],idx]
		idxDict[num]=idx
```


