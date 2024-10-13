// LeetCode 202 Happy Number
package happynumber

import "math"

func happyNumber(n int) bool {
	slow, fast := n, n
	for {
		fast = isHappy(isHappy(fast))
		slow = isHappy(slow)
		if fast == slow {
			return fast == 1
		}
	}
}

func isHappy(x int) int {
	var total int
	for x != 0 {
		total += int(math.Pow(float64(x % 10),2))
		x = x/10
	}
	return total
}