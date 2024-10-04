// Problem

// Minimum Coins
// Given a variety of coin types defining a currency system, find the minimum number of coins
// required to express a given amount of money. Assume infinite supply of coins of every type.

// Example
// {
// "coins": [1, 3, 5],
// "value": 9
// }
// Output:

// 3
// Here are all the unique ways to express 9 as a sum of coins 1, 3 and 5:

// 1, 1, 1, 1, 1, 1, 1, 1, 1
// 1, 1, 1, 1, 1, 1, 3
// 1, 1, 1, 1, 5
// 1, 1, 1, 3, 3
// 1, 3, 5
// 3, 3, 3
// Last two ways use the minimal number of coins, 3.

// Notes
// There will be no duplicate coin types in the input.

// Constraints:

// 1 <= number of coin types <= 102
// 1 <= coin value <= 102
// 1 <= amount of money to express <= 104
package minimumcoins

import "math"

func minimum_coins(coins []int, value int) int {
	table := make([]int, value+1)
	n := len(table)
	// Base case
	table[0] = 0

	// Initialization
	for i := 1; i < n; i++ {
		table[i] = math.MaxInt32
	}

	// Compute
	for i := 1; i < n; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				table[i] = Min(table[i], table[i-c])
			}
		}
		table[i]++
	}
	return table[value]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
