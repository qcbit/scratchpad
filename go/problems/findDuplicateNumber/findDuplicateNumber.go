// LeetCode 287
package findduplicatenumber

func detectCycle(nums []int) int {
	f := func(x int) int {
		return nums[x]
	}
	hare, tortoise := 0, 0
	for {
		hare = f(f(hare))
		tortoise = f(tortoise)
		if hare == tortoise { // cycle detected
			// find the index where the cycle begins
			var newTortoise int
			for newTortoise != tortoise {
				newTortoise = f(newTortoise)
				tortoise = f(tortoise)
			}
			return tortoise
		}
	}
}