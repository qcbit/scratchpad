// Problem

// Power
// Given a base a and an exponent b. Your task is to find ab. The value could be large enough. So, calculate ab % 1000000007.

// Example
// {
// "a": 2,
// "b": 10
// }
// Output:

// 1024
// Notes
// Constraints:

// 0 <= a <= 104
// 0 <= b <= 109
// a and b together won't be 0
package power

const mod = 1000000007

// calculate_power returns a^b
// Time complexity: O(log b)
// Space complexity: O(log b)
func calculate_power(a int64, b int64) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return int(a % mod)
	}
	a = a % mod
	tmp := calculate_power(a, b/2)
	if b%2 == 0 {
		return int((tmp%mod * tmp%mod) % mod)
	}

	return int((int64((tmp * tmp) % mod) * a) % mod)
}
