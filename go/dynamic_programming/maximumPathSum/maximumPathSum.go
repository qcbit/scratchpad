// Problem

// Maximum Path Sum
// Given a two dimensional grid of numbers. Find a path from top-left corner to bottom-right corner, which maximizes the sum of all numbers along its path.

// You can only move either down or right from your current position.

// Example One
// {
// "grid": [
// [4, 5, 8],
// [3, 6, 4],
// [2, 4, 7]
// ]
// }
// Output:

// 28
// Example 1 illustration

// The path 4 -> 5 -> 8 -> 4 -> 7 maximizes the sum. Every other path from top left to bottom right has sum less than 28.

// Example Two
// {
// "grid": [
// [1, -4, 3],
// [4, -2, 2]
// ]
// }
// Output:

// 5
// The path 1 -> 4 -> -2 -> 2 maximizes the sum. Note that there can be negative values in the grid as well.

// Notes
// Constraints:

// 1 <= number of rows <= 300
// 1 <= number of columns <= 300
// -104 <= numbers in the grid <= 104
package maximumpathsum

import "math"

func maximum_path_sum(grid [][]int) int {
	// Initialization
	rows := len(grid)
	cols := len(grid[0])
	table := make([][]int, rows)
	for i := 0; i < rows; i++ {
		table[i] = make([]int, cols)
	}
	table[0][0] = grid[0][0]

	// init 0th row
	for j := 1; j < cols; j++ {
		table[0][j] = table[0][j-1] + grid[0][j]
	}

	// init 0th column
	for i := 1; i < rows; i++ {
		table[i][0] = table[i-1][0] + grid[i][0]
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			table[row][col] = grid[row][col] + int(math.Max(float64(table[row-1][col]), float64(table[row][col-1])))
		}
	}
	return table[rows-1][cols-1]
}
