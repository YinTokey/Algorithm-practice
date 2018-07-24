
# 在一个二维数组中（每个一维数组的长度相同），
# 每一行都按照从左到右递增的顺序排序，
# 每一列都按照从上到下递增的顺序排序。
# 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

# 解答：从左下角开始查找。如果大了就往右，如果小了就往上。

# array 二维列表
def Find(self, target, array):
    col = len(array[0]) - 1
    row = len(array) -1
    j = 0
    i = row
    while j <= row and i >= 0:
        if target < array[i][j]:
            i -= 1
        elif target > array[i][j]:
            j += 1
        else
            return true
    return false