// Problem

// Generate Binary Strings Of Length N
// Given a number n, generate all possible binary strings of length n.

// Example
// {
// "n": 3
// }
// Output:

// ["000", "001", "010", "011", "100", "101", "110", "111"]
// Notes
// A string consisting of only 0s and 1s is called a binary string.
// Return the output list in any order.
// Constraints:

// 1 <= n <= 16
package binaryStringsOfLengthN

func get_binary_strings(n int) []string {
	results := make([]string, 0)
	helper(n, "", &results)
	return results
}

func helper(n int, slate string, result *[]string) {
	if n == 0 {
		*result = append(*result, slate)
		return
	}
	helper(n-1, slate+"0", result)
	helper(n-1, slate+"1", result)
}