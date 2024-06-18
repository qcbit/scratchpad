// Given a tolerance value to the system to check if the value is within the tolerance range
// If the value is outside the tolerance range, then return the set of values that are within the range.
// tolerance = 5
// values = [1, 2, 4, 10, 11, 15, 22, 23, 40, 50, 100, 101, 120]
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Values within the tolerance range: ", findToleranceSet(5, []int{1, 2, 4, 10, 11, 15, 22, 23, 40, 50, 100, 101, 120}))
	fmt.Println("Values within the tolerance range: ", findToleranceSet(5, []int{1, 2, 4, 10, 11, 15, 22, 23, 27, 35, 40, 46, 50}))
}

// findToleranceSet returns the set of values that are within the t tolerance range given the values v
func findToleranceSet(t int, v []int) [][]int {
	var set [][]int
	var i, j int
	j = 1
	for i = 0; i < len(v); i++ {
		var pair []int
		pair = append(pair, v[i])
		for ; j < len(v); j++ {
			// If the difference between the current value and the next value is greater than the tolerance
			if v[j]-v[i] > t {
				// If the previous value is not the same as the current value, then add it to the pair
				if j != i && v[j-1] != v[i] {
					pair = append(pair, v[j-1])
				}

				break
			}
			// Handle the last element
			if j == len(v)-1 && v[j] != v[i] {
				pair = append(pair, v[j])
			}
		}
		set = append(set, pair)
		i = j - 1
	}

	return set
}
