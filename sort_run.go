package main

import "fmt"

func main() {

	// array := []int{14, 7, 3, 12, 9, 11, 6, 2}

	// fmt.Print("Input: ")
	// fmt.Println([]int{14, 7, 3, 12, 9, 11, 6, 2})
	// sorted_array := MergeSort([]int{14, 7, 3, 12, 9, 11, 6, 2}, true)
	// fmt.Println(sorted_array)

	// fmt.Print("Input: ")
	// fmt.Println(array)
	// QuickSort(array, 0, len(array)-1)
	// fmt.Println(array)

	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Print("Input: ")
	fmt.Println(matrix)
	unfolded := Unfold(matrix)
	fmt.Println(unfolded)

}
