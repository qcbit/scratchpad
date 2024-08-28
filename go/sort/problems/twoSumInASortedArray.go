// 2 Sum In A Sorted Array
// Given an array sorted in non-decreasing order and a target number, find the 
// indices of the two values from the array that sum up to the given target number.

// Example
// {
// "numbers": [1, 2, 3, 5, 10],
// "target": 7
// }
// Output:

// [1, 3]
// Sum of the elements at index 1 and 3 is 7.

// Notes
// In case when no answer exists, return an array of size two with both values equal to -1, i.e., [-1, -1].
// In case when multiple answers exist, you may return any of them.
// The order of the indices returned does not matter.
// A single index cannot be used twice.
// Constraints:

// 2 <= array size <= 105
// -105 <= array elements <= 105
// -105 <= target number <= 105
// Array can contain duplicate elements.
package twoSumInASortedArray

// BruteForce 
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func BruteForce(numbers []int, target int) []int {
	for i, a := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if a+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1,-1}
}

// BruteForceTwoPointers
// Time Complexity: O(n)
// Space Complexity: O(1)
func BruteForceTwoPointers(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i, j}
		} else if numbers[i] + numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1,-1}
}

// BinarySearch
// Time Complexity: O(log(n))
// Space Complexity: O(1)
func BinarySearch(numbers []int, target int) []int {
	for aNdx := range numbers {
		start := aNdx + 1
		end := len(numbers) - 1
		b := target - numbers[aNdx]
		for start <= end {
			mid := start + (end - start) / 2
			if numbers[mid] == b {
				return []int{aNdx, mid}
			} else if numbers[mid] < b {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return []int{-1,-1}
}

// Transform uses a map to store the value and its index.
// If the value is found in the map, it returns the index of the value.
// Time Complexity:
// Space Complexity: O(n)
func Transform(numbers []int, target int) []int {
	values := make(map[int]int, len(numbers))
	for aNdx, a := range numbers {
		b := target - a
		if _, ok := values[b]; ok {
			return []int{aNdx, values[b]}
		} else {
			values[a] = aNdx
		}
	}
	return []int{-1,-1}
}