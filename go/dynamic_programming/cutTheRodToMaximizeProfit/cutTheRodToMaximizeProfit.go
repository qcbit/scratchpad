// Problem

// Cut The Rod To Maximize Profit
// Given the prices for rod pieces of every size between 1 inch and n inches,
// find the maximum total profit that can be made by cutting an n inches long rod inch into pieces.

// Example
// {
// "price ": [1, 5, 8, 9]
// }
// Output:

// 10
// The rod can be cut in the ways given below:

// 1 + 1 + 1 + 1 inches will cost price[0] + price[0] + price[0] + price[0] = 4
// 1 + 1 + 2 inches will cost price[0] + price[0] + price[1] = 7
// 1 + 3 inches will cost price[0] + price[2] = 9
// 2 + 2 inches will cost price[1] + price[1] = 10
// One piece of 4 inches will cost price[3] = 9
// The maximum profit is obtained by cutting it into two pieces 2 inches each.

// Notes
// Constraints:

// 1 <= n <= 103
// 1 <= price of any sized piece of the rod <= 105
package cuttherodtomaximizeprofit

import "math"

func get_maximum_profit(price []int) int {
	// n := len(price)
	// memo := make(map[int]int, n+1)
	// rodCutting(price, memo, n)
	// return memo[n]
	return get_maximum_profit_tabulation(price)
}

func rodCutting(price []int, memo map[int]int, n int) int {
	// if already calculated, then return it
	if n, ok := memo[n]; ok {
		return memo[n]
	}

	// base case
	if n == 0 {
		return 0
	}

	max := int(math.Inf(-1))

	for i := 1; i <= n; i++ {
		max = int(math.Max(float64(max), float64(price[i-1]))) + rodCutting(price, memo, n-i)
	}

	memo[n] = max
	return max
}

func get_maximum_profit_tabulation(price []int) int {
	n := len(price)
	table := make([]int, n+1)

	//base case
	// if rod of length 0, then return 0
	table[0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			table[i] = max(table[i], price[j]+table[i-j-1])
		}
	}

	return table[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
