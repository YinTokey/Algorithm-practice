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

	return T;


}



var coins = [1,5,6,8];
var total = 11
var n = coins.length

console.log(minCoins(coins,total,n));
