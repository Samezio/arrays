package main

func Unfold(matrix [][]int) []int {
	t := 0
	b := len(matrix) - 1
	l := 0
	r := len(matrix[0]) - 1

	direction := 0 // Left to right: 0, top to bottom: 1, right to left: 2, bottom to top: 3

	result := make([]int, len(matrix)*len(matrix[0]))
	index := 0
	for t <= b && l <= r {
		if direction == 0 {
			for i := l; i <= r; i++ {
				result[index] = matrix[t][i]
				index++
			}
			t++
			direction++
		} else if direction == 1 {
			for i := t; i <= b; i++ {
				result[index] = matrix[i][r]
				index++
			}
			r--
			direction++
		} else if direction == 2 {
			for i := r; i >= l; i-- {
				result[index] = matrix[b][i]
				index++
			}
			b--
			direction++
		} else {
			for i := b; i >= t; i-- {
				result[index] = matrix[i][l]
				index++
			}
			l++
			direction = 0
		}
	}
	return result
}
