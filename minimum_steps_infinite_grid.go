package main

//Minimum steps needed to cover a sequence of points on an infinite grid
//Given an infinite grid, initial cell position (x, y) and a sequence of other cell position which needs to be covered in the given order. The task is to find the minimum number of steps needed to travel to all those cells.
//Movement can be done in any of the eight possible directions from a given cell i.e from cell (x, y) you can move to any of the following eight positions:(x-1, y+1), (x-1, y), (x-1, y-1), (x, y-1), (x+1, y-1), (x+1, y), (x+1, y+1), (x, y+1) is possible.

func MinSteps(points_sequence [][]int) int {
	current_point := points_sequence[0]
	moves := 0
	for _, point := range points_sequence[1:] {
		for current_point[0] != point[0] && current_point[1] != point[1] {
			if current_point[0] > point[0] {
				current_point[0]--
			} else if current_point[0] < point[0] {
				current_point[0]++
			}
			if current_point[1] > point[1] {
				current_point[1]--
			} else if current_point[1] < point[1] {
				current_point[1]++
			}
			moves++
		}
	}
	return moves
}
