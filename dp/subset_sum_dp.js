
// https://www.youtube.com/watch?v=s6FhG--P7z0&index=5&list=PLrmLmBdmIlpsHaNTPP_jHHDx_os9ItYXr&t=0s
//子集和问题
//给定一个包含N个非负数的set， 并且给定K, 从集合中找出一组元素子集， 使得这组子集的各个元素相加起来和是K。
/*
 给定非负数集: 2,3,7,8,10，记为 input[i]，k = 11。求元素子集（每个元素只能使用1次）

*/

function subsetSum(input,k,n){
	var T = []; //T[i][j] 为1或0，表示true or false

	for(let i=0;i<n;i++){
		T[i] = [];
		for(let j=0;j<=k;j++){
			if(j==0){
				T[i][j] = 1;  //第0列，都是true，此时子集合为空集
				continue;
			}

			if(i==0){//第0行无法使用i-1,单独处理
				if(input[i]!=j){
					T[i][j]=0;
				}else{
					T[i][j]=1
				}

			}else{
				//常规情况
				if(j<input[i]){
					T[i][j] = T[i-1][j];
				}else{
					T[i][j] = T[i-1][j] || T[i-1][j-input[i]]
				}
			}

		}

	}

	console.log(T);
	findValue(input,k,n,T);
	return T;
}


function findValue(input,k,n,T){
	//老套路，反向寻找
	var i = n-1,j=k;
	var result = []; //结果保存在数组中
	while(i>0&&j>0){
		if(T[i][j] == T[i-1][j]){ //如果T[i][j]的值来自上方，则回退
			i--;
		}else{
			result.unshift(input[i])//直到T[i][j]的值与上方不同，选中它。此时总和j减去被选中的input[i]
			j = j-input[i];
			i--;
		}
		//虽然上面 i-- 可以简化，但是建议这样写，算法显得比较直观。
	}
	console.log(result);
}


var input = [2,3,7,8,10];
var n = input.length;
var k = 11;
subsetSum(input,k,n);


