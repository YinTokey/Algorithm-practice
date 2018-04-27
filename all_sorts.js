//Review all sorts algoritm written in Javascript


var array = [3,9,11,12,2,5,88,32];
console.log("before:" + array);


selectSort(array);
console.log("after:" + array);

//bubble sort
function bubbleSort(array){
	for(var i=0;i<array.length ;i++){
		for(var j=0;j<array.length-1-i;j++){
			// in this circle, sort the bigget number to the right of array
			if(array[j]>array[j+1]){
				swap(array,j+1,j);
			}
		}
	}
}

//select sort
function selectSort(array){
	var len = array.length;
	for(var i=0;i<len;i++){
		var index = i;
		for(var j=i;j<len-1;j++){
			if(array[index]>array[j+1]){
				index = j+1;
			}
		}
		if(index !== i){
			swap(array,index,i);

		}
	}

}


function swap(array,index1,index2){
    var tmp  = array[index2];
    array[index2] = array[index1];
    array[index1] = tmp;

}
