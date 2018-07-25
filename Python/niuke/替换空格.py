# 请实现一个函数，将一个字符串中的每个空格替换成“%20”。
# 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。

def replaceSpace(string):
    result = []
    for i in string:
        if i == " ":
            result.append("%20")
        else:
            result.append(i)
    return ''.join(result)


str = "We Are Happy"
result = replaceSpace(str)
print(result)