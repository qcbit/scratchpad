// In a given sorted array of integers remove all the duplicates.
// https://go.dev/play/p/Maky13vHIQi
package main

import "fmt"

func main() {
	// case: odd elements
	fmt.Println(filterDuplicates([]int{2, 2, 2, 5, 7, 7, 8, 9, 9, 9, 9}))
	// case: even elements
	fmt.Println(filterDuplicates([]int{2, 2, 2, 5, 7, 7, 8, 9, 9, 9}))
	// case: two elements
	fmt.Println(filterDuplicates([]int{1, 2}))
	// case: two same elements
	fmt.Println(filterDuplicates([]int{1, 1}))
	// case: no duplicates
	fmt.Println(filterDuplicates([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func filterDuplicates(sortedArray []int) []int {
	var a, b, p int

	// Solves case 0 and 1
	if len(sortedArray) < 2 {
		return sortedArray
	}

	for b < len(sortedArray) {
		if sortedArray[p] != sortedArray[b] {
			sortedArray[a+1] = sortedArray[b]
			a++
		}
		b++
		p = b - 1
	}
	return sortedArray[:a+1]
}
