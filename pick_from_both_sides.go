package main

// Given an integer array A of size N. You can pick B elements from either left or right end of the array A to get maximum sum. Find and return this maximum possible sum.
func PickMax(array []int, to_pick int) (string, int) {
	if to_pick >= len(array) {
		sum := 0
		for _, ele := range array {
			sum += ele
		}
		return "left", sum
	}
	sum_left := 0
	sum_right := 0
	for i := 0; i < to_pick; i++ {
		sum_left += array[i]
		sum_right += array[len(array)-i-1]
	}
	if sum_left >= sum_right {
		return "left", sum_left
	} else {
		return "right", sum_right
	}
}

func PickMin(array []int, to_pick int) (string, int) {
	sum_left := 0
	sum_right := 0
	for i := 0; i < to_pick; i++ {
		sum_left += array[i]
		sum_right += array[len(array)-i-1]
	}
	if sum_left <= sum_right {
		return "left", sum_left
	} else {
		return "right", sum_right
	}
}
