// Problem

// Unique Paths
// Given a grid of size n x m and a robot initially staying at the top-left position, return the number of unique paths for the robot to reach the bottom-right corner of the grid. The robot is allowed to move either down or right at any point in time.

// Example One
// {
// "n": 3,
// "m": 2
// }
// Output:

// 3
// From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:

// Right -> Down -> Down
// Down -> Down -> Right
// Down -> Right -> Down
// Example Two
// {
// "n": 4,
// "m": 1
// }
// Output:

// 1
// From the top-left corner, there is only one way to reach bottom-right corner:

// Down -> Down -> Down
// Notes
// Return the answer modulo 10^9 + 7.

// Constraints:

// 1 <= n, m <= 103
package uniquepaths

func unique_paths(n int, m int) int {
	table := make([][]int, n)
	for i := 0; i <= n-1; i++ {
		table[i] = make([]int, m)
	}

	// 0th column initialization
	for i := 0; i <= n-1; i++ {
		table[i][0] = 1
	}

	// 0th row initialization
	for i := 0; i <= m-1; i++ {
		table[0][i] = 1
	}

	for row := 1; row <= n-1; row++ {
		for col := 1; col <= m-1; col++ {
			table[row][col] = (table[row-1][col] + table[row][col-1]) % 1000000007
		}
	}
	return table[n-1][m-1]
}
