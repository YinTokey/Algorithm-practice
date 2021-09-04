package main

//  每次 arr 都是整个数组，而不是子数组。通过 low high 来区分
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i],arr[j] = arr[j],arr[i]
			i++
		}
	}
	arr[i],arr[high] = arr[high],arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr,low,high)
		arr = quickSort(arr,low,p-1)
		arr = quickSort(arr,p+1,high)
	}
	return arr
}


func getPivot(arr []int) int {
	length := len(arr)
	mid := arr[length/2]
	if mid > arr[0]  && mid < arr[length-1] {
		return length/2
	} else if arr[0] > mid && arr[0] < arr[length-1] {
		return 0
	} else {
		return length-1
	}

}

func findTopK(arr []int, k int) ([]int) {

	low := 0
	high := len(arr)-1

	for  {
		_, p := partition(arr,low, high)
		if p == k - 1 {
			break
		} else if p > k - 1 {
			high = k-1
		} else {
			low = k - 1
		}
	}

	re := arr[:k]

	return re
}