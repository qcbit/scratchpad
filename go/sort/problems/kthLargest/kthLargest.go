// 215. Kth Largest Element in an Array
// Find the kth largest element in an unsorted array. Note that it is the kth
// largest element in the sorted order, not the kth distinct element.
//
// Example 1:
// Input: [3,2,1,5,6,4] and k = 2
// Output: 5
//
// Example 2:
// Input: [3.,2,3,1,2,4,5,5,6] and k = 4
// Output: 4
//
// Note: You may assume k is always valid. 1 <= k <= array's length
package kthlargest

import "math/rand"

func findKthLargest(input []int, k int) int {
	n := len(input)
	helper(input, 0, n-1, n-k)
	return input[n-k]
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
