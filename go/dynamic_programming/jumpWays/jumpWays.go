// Problem

// Jump Ways
// There is a one-dimensional axis. In one turn, you can take a jump of length 1 or 2. Find the total number of distinct ways using which you can reach from 0th coordinate to n-th coordinate?

// Example One
// {
// "n": 3
// }
// Output:

// 3
// There are 3 distinct ways in which you can move from 0 to 3.

// 1 length jump + 1 length jump + 1 length jump.
// 1 length jump + 2 length jump.
// 2 length jump + 1 length jump.
// Example Two
// {
// "n": 4
// }
// Output:

// 5
// There are 5 distinct ways in which you can move from 0 to 5.

// Notes
// Constraints:

// 1 <= n <= 70
package jumpways

func jump_ways(n int) int64 {
	memo := make([]int64, 3)
	memo[1] = 1
	memo[2] = 2
	for i := 3; i <= n; i++ {
		memo[i%3] = memo[(i-1)%3] + memo[(i-2)%3]
	}
	return memo[n%3]
}
