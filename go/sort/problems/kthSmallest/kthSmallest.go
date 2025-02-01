// 215. Kth Smallest Element in an Array
// Find the kth smallest element in an unsorted array. Note that it is the kth
// smallest element in the sorted order, not the kth distinct element.
//
// Example 1:
// Input: [3,2,1,5,6,4] and k = 2
// Output: 2
//
// Example 2:
// Input: [3.,2,3,1,2,4,5,5,6] and k = 4
// Output: 6
//
// Note: You may assume k is always valid. 1 <= k <= array's length
package kthsmallest

import "math/rand"

func findKthSmallest(input []int, k int) int {
	n := len(input)
	helper(input, 0, n-1, k-1)
	return input[k-1]
}

func helper(A []int, start, end, index int) {
	if start == end {
		return
	}

	pivot := rand.Int()%(end-start) + start
	A[start], A[pivot] = A[pivot], A[start]
	pivot = start

	i := start
	for j := start + 1; j <= end; j++ {
		if A[j] < A[start] {
			i++
			A[j], A[i] = A[i], A[j]
		}
	}

	A[start], A[i] = A[i], A[start]

	if index == pivot {
		return
	} else if index < pivot {
		helper(A, start, pivot-1, index)
	} else {
		helper(A, pivot+1, end, index)
	}
}
