package sort

import "math/rand"

// QuickSort sorts an array of integers
func QuickSort(A []int) []int {
	qshelper(A, 0, len(A)-1)
	return A
}

func qshelper(A []int, start, end int) {
	if start >= end {
		return
	}

	pivot := rand.Int()%(end-start) + start
	A[start], A[pivot] = A[pivot], A[start]
	pivot = start // set pivot at start position

	// Lomuto's partition scheme
	i := start
	for j := start+1; j <= end; j++ {
		if A[j] < A[start] {
			i++
			A[j], A[i] = A[i], A[j]
		}
	}

	A[start], A[i] = A[i], A[start]

	qshelper(A, start, i-1)
	qshelper(A, i+1, end)
}
