package sort

// InsertionSort sorts an Aay of integers
func InsertionSort(A []int) []int {
	for i := 0; i < len(A); i++ {
		temp := A[i]
		j := i - 1
		for j >= 0 && A[j] > temp {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
	return A
}
