// Problem

// N Choose R
// Consider you have n distinct elements, find the number of ways through which you can choose exactly r number of elements out of those.

// Example One
// {
// "n": 5,
// "r": 3
// }
// Output:

// 10
// There is a set of 5 elements {1, 2, 3, 4, 5}. You need to pick a subset of size 3. Eligible subsets are {1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}. There are 10 subsets of size 3.

// Example Two
// {
// "n": 3,
// "r: 5
// }
// Output:

// 0
// There is a set of 3 elements {1, 2, 3}. You need to pick a subset of size 5. Which is not possible, that's why the answer is 0.

// Notes
// Here the answer can be very big, find it modulo 109 + 7.

// Constraints:

// 0 <= n, r <= 1000
package nchooser

import "math"

func ncr(n int, r int) int {
	if r == 0 || r == n {
		return 1
	}
	if r > n {
		return 0
	}
	table := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		table[i] = make([]int, r+1)
	}

	for i := 0; i <= n; i++ {
		table[i][0] = 1
	}
	for i := 0; i <= r; i++ {
		table[i][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= int(math.Min(float64(i), float64(r))); j++ {
			table[i][j] = (table[i-1][j] + table[i-1][j-1]) % 1000000007
		}
	}
	return table[n][r]
}
