
// 解决方案参考
// https://blog.csdn.net/Faremax/article/details/78078259

/*
二叉树深度
*/

function TreeDepth(pRoot)
{
    if(pRoot === null){
        return 0;
    }
    var left = TreeDepth(pRoot.left);
    var right = TreeDepth(pRoot.right);
    return Math.max(left +1,right + 1 );
}


/*
不用加减乘除做加法
写一个函数，求两个整数之和，要求在函数体内不得使用+、-、*、/四则运算符号。
思路：
两个数异或：相当于每一位相加，而不考虑进位；
两个数相与，并左移一位：相当于求得进位；
将上述两步的结果相加
*/
function Add(num1, num2)
{
    var sum, carry;
    while(num2 != 0){
        sum = num1 ^ num2;
        carry = (num1 & num2) << 1;
        num1 = sum;
        num2 = carry;
    }
    return num1;
}


/*
镜像二叉树

*/
function Mirror(root)
{
    if(!root){
        return null;
    }
    var tempRoot = root.left;
    root.left = root.right;
    root.right = tempRoot;
    
    if(root.left){
        Mirror(root.left);
    }
    if(root.right){
        Mirror(root.right);
    }
}


/*
台阶问题

一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。

可以列举前几个，得出规律。
*/
function jumpFloorII(number)
{
    // write cde here
    return Math.pow(2,number-1);
}

/*
构建乘积数组
画矩阵，分别求上下三角。
B【0】=1是可以先求出来的。

*/




/*
求1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
利用逻辑与的短路求值
*/ 
function Sum_Solution(n)
{
    // write code here
    var sum = n;
    (n >0)&&(sum+=Sum_Solution(n-1))
    return sum;
}


