#递归和循环
#斐波那契数列

# 不使用递归

def fibonacii(num):
    if num <= 0:
        return  0
    elif num == 1:
        return 1
    else:
        j,k,fib = 0,1,0
        for i in range(2,num+1):
            fib = j + k
            j = k
            k = fib
        return fib


print(fibonacii(10))


