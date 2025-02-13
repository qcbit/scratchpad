// Single Number
// Easy
// Given a non-empty array of integers nums, every element appears twice except for one.
// Find that single one.
// Implement a solution with a linear runtime complexity and
// use only constant extra space.
package singlenumber

func find_single_number(nums []int) int {
	var number int
	for _, num := range nums {
		number ^= num
	}
	return number
}
