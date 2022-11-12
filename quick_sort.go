package main

func partition(array []int, low int, high int) int {
	pivot := array[high]
	i := low - 1
	for j := low; j < high; j++ {
		if array[j] < pivot {
			i++
			//swap
			temp := array[i]
			array[i] = array[j]
			array[j] = temp
		}
	}

	i++
	temp := array[i]
	array[i] = pivot
	array[high] = temp
	return i
}
func QuickSort(array []int, low int, high int) {
	if low < high {
		pivot_index := partition(array, low, high)
		QuickSort(array, low, pivot_index-1)
		QuickSort(array, pivot_index+1, high)
	}
}
