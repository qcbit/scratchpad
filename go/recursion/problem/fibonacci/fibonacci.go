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

func find_fibonacci(n int) int {
	return addseq(n, 0, 1)
}

func addseq(n, b1, b2 int) int {
	if n == 0 {
			return b1
	}
	return addseq(n-1, b2, b1+b2)
}