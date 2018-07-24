//Review all sorts algoritm written in Javascript


var array = [3,9,11,12,2,5,88,32];
console.log("before:" + array);


quickSort(array,0,array.length-1);
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
		for(var j=i;j<len-1;j++){ // the initial value j should be equal to i, not "j=0"
			if(array[index]>array[j+1]){
				index = j+1;
			}
		}
		if(index !== i){
			swap(array,index,i);

		}
	}

}

//insert sort
function insertSort(array){
 	var len = array.length;
	for(var i=0;i<len;i++){
		var j = i;
		while(j>0){
			if(array[j]<array[j-1]){
				swap(array,j,j-1);
			}
			j--;

		}
	}

}


//quick sort
function quickSort(arr, left, right){
   var len = arr.length, 
   pivot,
   partitionIndex;


  if(left < right){
    pivot = right;
    partitionIndex = partition(arr, pivot, left, right);
    
   //sort left and right
   quickSort(arr, left, partitionIndex - 1);
   quickSort(arr, partitionIndex + 1, right);
  }
  return arr;
}

function partition(arr, pivot, left, right){
   var pivotValue = arr[pivot],
       partitionIndex = left;

   for(var i = left; i < right; i++){
    if(arr[i] < pivotValue){
      swap(arr, i, partitionIndex);
      partitionIndex++;
    }
  }
  swap(arr, right, partitionIndex);
  return partitionIndex;
}


function swap(array,index1,index2){
    var tmp  = array[index2];
    array[index2] = array[index1];
    array[index1] = tmp;

}
