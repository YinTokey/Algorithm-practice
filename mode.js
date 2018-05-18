//dynamic programming
// 01背包问题
//



// 4个已知量：重量集合，价值集合，总容量（总列数），集合总数（总行数）

function knapSack(w,val,capacity,n){
	var T = []

	for(let i = 0;i < n;i++){
		T[i] = [];
		for(let j=0;j <= capacity;j++){
			if(j === 0){ //容量为0
				T[i][j] = 0;
				continue;
			}	
			if(j < w[i]){ //容量小于物品重量，hold不住
				if(i === 0){
					T[i][j] = 0;

				}else{
					T[i][j] = T[i-1][j];

				}
				continue;
			}
			if(i === 0){
				T[i][j] = val[i];
			}else{
				T[i][j] = Math.max(val[i] + T[i-1][j-w[i]],T[i-1][j]);

			}

		}
	

	}
	return T;
}

// w = [2,3,4].  val = [3,4,5] , n = 3 , capacity = 5
//function knapSack([2,3,4],[3,4,5],5,3);
// 
var values = [3,4,5],
	weights = [2,3,4],
	capacity = 5,
	n = values.length;

console.log(knapSack(weights,values,capacity,n));

