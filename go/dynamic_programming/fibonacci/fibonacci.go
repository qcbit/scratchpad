// Problem

// Fibonacci Number
// Given a number n, find the n-th Fibonacci Number.

// Example
// {
// "n": 2
// }
// Output:

// 1
// 2nd Fibonacci Number is the sum of the 0th and 1st Fibonacci Number = 0 + 1 = 1.

// Notes
// In mathematics, the Fibonacci numbers, commonly denoted Fn, form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
// F(0) = 0, F(1) = 1 and F(n) = F(n − 1) + F(n − 2) for n > 1.

// Constraints:

// 0 <= n <= 46
package fibonacci

// find_fibonacci returns the nth Fibonacci number.
// The function uses a technique called memoization to store the results of expensive function calls and return the cached result when the same inputs occur again.
// T(n) = O(n)
// S(n) = O(1)
func find_fibonacci(n int) int {
	memo := make([]int, 3)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i%3] = memo[(i-1)%3] + memo[(i-2)%3]
	}
	return memo[n%3]
}
