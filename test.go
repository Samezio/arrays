package main

import "fmt"

func Test_merge_sort() {
	fmt.Println("Merge sort")

	array := []int{14, 7, 3, 12, 9, 11, 6, 2}
	fmt.Print("Array: ")
	fmt.Println(array)
	sorted_array := MergeSort(array, true)
	fmt.Print("Sorted Array: ")
	fmt.Println(sorted_array)
}

func Test_quick_sort() {
	fmt.Println("Quick sort")

	array := []int{14, 7, 3, 12, 9, 11, 6, 2}
	fmt.Print("Array: ")
	fmt.Println(array)
	QuickSort(array, 0, len(array)-1)
	fmt.Print("Sorted Array: ")
	fmt.Println(array)
}

func Test_spiral_unfolding() {
	fmt.Println("Spiral unfolding of matrix")

	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("Matrix: ")
	for _, row := range matrix {
		fmt.Println(row)
	}
	unfolded := Unfold(matrix)
	fmt.Print("Unfolded: ")
	fmt.Println(unfolded)
}

func Test_pick_from_both_side() {
	fmt.Println("Pick from both side.")
	fmt.Println("??: Given an integer array A of size N. You can pick B elements from either left or right end of the array A to get maximum sum. Find and return this maximum/minimum possible sum.")

	array := []int{14, 7, 102, 3, 12, 9, 11, 6, 91, 2, 10}
	fmt.Print("Array: ")
	fmt.Println(array)
	fmt.Print("Number of elements to pick: ")
	fmt.Println(4)
	fmt.Print("Max Pick: ")
	fmt.Println(PickMax(array, 4))
	fmt.Print("Min Pick: ")
	fmt.Println(PickMin(array, 4))
}

func main() {
	Test_merge_sort()
	fmt.Println("--------------------------")
	Test_quick_sort()
	fmt.Println("--------------------------")
	Test_spiral_unfolding()
	fmt.Println("--------------------------")
	Test_pick_from_both_side()
	fmt.Println("--------------------------")
}
