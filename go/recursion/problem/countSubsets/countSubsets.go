// Problem

// Count All Subsets Of A Set Of Size N
// Given a number denoting the size of a set, count the number of unique subsets of that set.

// Example
// {
// "n": 3
// }
// Output:

// 8
// If we have a set {1, 2, 3}, then all the possible subsets are: {}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}.

// Notes
// Empty subset should also be counted.

// Constraints:

// 0 <= input number <= 30.
package countsubsets

func count_all_subsets(n int) int {
    if n == 0 {
		return 1
	}
    return 2 * count_all_subsets(n-1)
}