package sort

// BubbleSort sorts an array of integers
func BubbleSort(A []int) []int {
	n := len(A)
	if n == 0 {
		return A
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if A[j-1] > A[j] {
				A[j-1], A[j] = A[j], A[j-1]
			}
		}
	}

	return A
}