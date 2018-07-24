// https://www.youtube.com/watch?v=CE2b_-XfVDk&index=7&list=PLrmLmBdmIlpsHaNTPP_jHHDx_os9ItYXr
// largest increaseing subsequense    最长递增子序列
/*
和之前的子序列问题一样，所求子序列是不连续的
e.g.   3 4 -1 0 6 2 3  最长递增子序列为  -1 0 2 3
2 5 1 8 3  最长为  2 5 8
*/

 /*
	if arr[j]<arr[i]
	T[i] = max(T[i],T[j] + 1)

 */

function largestIncre(arr,n){
	if(n==1){
		return arr;
	}
	var T = []
	//构建初始T，值全部为1
	for(let k = 0;k<n;k++){
		T[k] = 1;
	}

	for(let i=1;i<n;i++){
		for(let j=0;j<i;j++){
			if(arr[j]<arr[i]){
				T[i] = Math.max(T[i],T[j]+1)
			}
		}
	}
	console.log(T);
	console.log(largestNum(T,n))

}


function findValue(arr,n,T){
	var result = []
	//从右往左搜索。起始位置为表中最大值的位置。 原则：搜索它左边离它最近的，值为 T[i]-1的。
	var i = largestNum(arr,n);
//	var k = 

	// for(;i>=0;i--){

	// }

}

//查找列表最大值
function largestNum(T,n){
	var result = 1;
	for(let i=n-1;i>=0;i--){
		result = Math.max(result,T[i])
	}
	return result
}

var arr = [2,-1,3,11,5,0,9,8];
var n = arr.length;

largestIncre(arr,n);

