// Problem

// Min Cost Climbing Stairs
// There are n stairs indexed 0 to n – 1. Each stair has a cost associated with it. After paying the cost, it's allowed either to climb one or two steps further. It's allowed to either start from the step with index 0, or the step with index 1.
// Given the cost array, find the minimum cost to reach the top of the floor.
// cost[i] represents the cost of i-th stair.

// Example One
// {
// "cost": [1, 1, 2]
// }
// Output:

// 1
// There are 5 ways to reach the top floor.

// step 0 → step 1 → step 2 → top floor.
// step 0 → step 1 → top floor.
// step 0 → step 2 → top floor.
// step 1 → step 2 → top floor.
// step 1 → top floor.
// Here, the last way(step 1 → top floor) will provide the most optimal cost 1.

// Example Two
// {
// "cost": [3, 4]
// }
// Output:

// 3
// Notes
// Constraints:

// 2 <= length of the input array <= 1000
// 0 <= cost[i] <= 999, for all i.
package mincostclimbingstairs

import "math"

func min_cost_climbing_stairs(cost []int) int {
	n := len(cost)
	cost = append(cost, 0)
	table := make([]int, n+2)
	table[0], table[n+1] = 0, 0
	table[1] = cost[0]

	for i := 2; i < len(table); i++ {
		table[i] = cost[i-1] + int(math.Min(float64(table[i-1]), float64(table[i-2])))
	}
	return table[n+1]
}
