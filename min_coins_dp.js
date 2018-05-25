//dynamic programming
// 最少硬币问题
//

function minCoins(coins,total,n){
	var T = [];

	for(let i = 0;i<n;i++){
		T[i] = []
		for (let j=0;j<= total;j++){
			if(j == 0){
				T[i][j] = 0;
				continue;
			}

			if(i == 0){
				T[i][j] = j/coins[i]; //硬币找零一定要有个 最小面额1，否则会无解
			}else{
				if(j >= coins[i]){
					T[i][j] = Math.min(T[i-1][j],1+T[i][j-coins[i]])
			
				}else{
					T[i][j] = T[i-1][j];
				}
			}

		}

	}
	findValue(coins,total,n,T);

	return T;


}

function findValue(coins,total,n,T){
	var i = n-1, j = total;
	var result = []; //结果保存在数组中
	while(i>0 && j >0){
		if(T[i][j]!=T[i-1][j]){
			//锁定位置，开始找构成结果的硬币组合
			console.log(T[i][j]);
			break
		}else{
			i--;
		}
	}

}


var coins = [1,5,6,8];
var total = 11
var n = coins.length

console.log(minCoins(coins,total,n));
