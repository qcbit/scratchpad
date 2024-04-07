// In a given sorted array of integers remove all the duplicates.
// https://go.dev/play/p/Maky13vHIQi
package main

import "fmt"

func main() {
	// case: odd elements
	fmt.Println(filter_duplicates([]int{2, 2, 2, 5, 7, 7, 8, 9, 9, 9, 9}))
	// case: even elements
	fmt.Println(filter_duplicates([]int{2, 2, 2, 5, 7, 7, 8, 9, 9, 9}))
	// case: two elements
	fmt.Println(filter_duplicates([]int{1, 2}))
	// case: two same elements
	fmt.Println(filter_duplicates([]int{1, 1}))
	// case: no duplicates
	fmt.Println(filter_duplicates([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func filter_duplicates(sorted_array []int) []int {
	var a, b, p int

	// Solves case 0 and 1
	if len(sorted_array) < 2 {
		return sorted_array
	}

	for b < len(sorted_array) {
		if sorted_array[p] != sorted_array[b] {
			sorted_array[a+1] = sorted_array[b]
			a++
		}
		b++
		p = b - 1
	}
	return sorted_array[:a+1]
}
