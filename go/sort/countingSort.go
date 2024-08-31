package sort

import (
	"math"
)

func CountingSort(A []int) []int {
	count := make(map[int]int, len(A))
	for _, num := range A {
		count[num]++
	}
	newArr := make([]int, 0, len(A))
	for i := int(-4 * math.Pow(10, 5)); i <= int(4*math.Pow(10, 5)); i++ {
		if _, ok := count[i]; ok {
			for count[i] > 0 {
				newArr = append(newArr, i)
				count[i]--
			}
		}
	}
	return newArr
}
