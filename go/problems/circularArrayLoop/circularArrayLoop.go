// LeetCode 457
package circulararrayloop

func detectCycle(nums []int) bool {
	f := func(x int) int {
		index := (x + nums[x]) % len(nums)
		if index < 0 {
			index = len(nums) + index
		}
		return index
	}
	for i := 0; i < len(nums); i++ {
		hare, tortoise := i, i
		for {
			hare = f(f(hare))
			tortoise = f(tortoise)
			if hare == tortoise {
				third := tortoise
				positive := nums[third] > 0
				count := 1
				for f(third) != tortoise {
					third = f(third)
					if (positive && nums[third] < 0) || (!positive && nums[third] > 0) {
						count = 1
						break
					}
					count++
				}
				if count > 1 {
					return true
				} else {
					// count == 1, not a special cycle
					break
				}
			}
		}
	}
	return false
}