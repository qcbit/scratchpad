package sort

// SelectionSort sorts an array of integers by selection sort
func SelectionSort(A []int) []int {
	if len(A) == 0 {
		return A
	}
	n := len(A)
	for i := 0; i < n-1; i++ {
		minNdx := i
		for j := i + 1; j < n; j++ {
			if A[j] < A[minNdx] {
				minNdx = j
			}
		}
		A[i], A[minNdx] = A[minNdx], A[i]
	}
	return A
}