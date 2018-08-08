'''
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非减排序的数组的一个旋转，输出旋转数组的最小元素。
 例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
  NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。

  思路：使用二分法解决

'''


class Solution:
    def minNumberInRotateArray(self, rotateArray):
        if len(rotateArray) == 0:
            return 0
        elif len(rotateArray) == 1:
            return rotateArray[0]

        else:
            low = 0
            high = len(rotateArray) - 1
            while low < high:
                mid = int((low + high) / 2)
                if rotateArray[mid] < rotateArray[high]:
                    high = mid
                elif rotateArray[mid] > rotateArray[high]:
                    low = mid + 1
                else:
                    high = high - 1
            return rotateArray[high]

