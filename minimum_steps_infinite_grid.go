package main

//Minimum steps needed to cover a sequence of points on an infinite grid
//Given an infinite grid, initial cell position (x, y) and a sequence of other cell position which needs to be covered in the given order. The task is to find the minimum number of steps needed to travel to all those cells.
//Movement can be done in any of the eight possible directions from a given cell i.e from cell (x, y) you can move to any of the following eight positions:(x-1, y+1), (x-1, y), (x-1, y-1), (x, y-1), (x+1, y-1), (x+1, y), (x+1, y+1), (x, y+1) is possible.
func move(start_point []int, end_point []int) (int, []int) {
	if start_point[0] == end_point[0] {
		moves := start_point[1] - end_point[1]
		if moves < 0 {
			moves *= -1
		}
		return moves, end_point
	} else if start_point[1] == end_point[1] {
		moves := start_point[0] - end_point[0]
		if moves < 0 {
			moves *= -1
		}
		return moves, end_point
	} else {
		moves1 := start_point[0] - end_point[0]
		if moves1 < 0 {
			moves1 *= -1
		}
		moves2 := start_point[0] - end_point[0]
		if moves2 < 0 {
			moves2 *= -1
		}
		if moves1 < moves2 {
			x := start_point[0] + moves1
			y := start_point[1] + moves1
			if end_point[0] < start_point[0] {
				x = start_point[0] - moves1
			}
			if end_point[1] < start_point[1] {
				y = start_point[1] - moves1
			}
			return moves1, []int{x, y}
		} else {
			x := start_point[0] + moves2
			y := start_point[1] + moves2
			if end_point[0] < start_point[0] {
				x = start_point[0] - moves2
			}
			if end_point[1] < start_point[1] {
				y = start_point[1] - moves2
			}
			return moves2, []int{x, y}
		}
	}
}
func MinSteps(points_sequence [][]int) int {
	current_point := points_sequence[0]
	moves := 0
	for _, point := range points_sequence[1:] {
		for current_point[0] != point[0] && current_point[1] != point[1] {
			moved, p := move(current_point, point)
			moves += moved
			current_point = p
		}
	}
	return moves
}
