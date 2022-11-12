package main

import "fmt"

func conquer(array []int, start_index int, mid_index int, end_index int, isAscending bool) {
	var merged_array []int = make([]int, end_index-start_index+1)
	index1 := start_index
	index2 := mid_index + 1
	var x int = 0

	for index1 <= mid_index && index2 <= end_index {
		if array[index1] <= array[index2] {
			merged_array[x] = array[index1]
			x++
			index1++
		} else {
			merged_array[x] = array[index2]
			x++
			index2++
		}
	}
	for index1 <= mid_index {
		merged_array[x] = array[index1]
		x++
		index1++
	}

	for index2 <= end_index {
		merged_array[x] = array[index2]
		x++
		index2++
	}

	//Copying to original array
	for i, val := range merged_array {
		array[i+start_index] = val
	}
	fmt.Print("Merge: ")
	fmt.Println(array[start_index : end_index+1])
}
func divide(array []int, start_index int, end_index int, isAscending bool) {
	if start_index >= end_index {
		return
	}
	var mid int = start_index + (end_index-start_index)/2
	fmt.Print("Divide: ")
	fmt.Print(array[start_index : mid+1])
	fmt.Print("  ----  ")
	fmt.Println(array[mid+1 : end_index])
	divide(array, start_index, mid, isAscending)
	divide(array, mid+1, end_index, isAscending)
	conquer(array, start_index, mid, end_index, isAscending)
}
func MergeSort(array []int, isAscending bool) []int {
	divide(array, 0, len(array)-1, isAscending)
	return array
}
