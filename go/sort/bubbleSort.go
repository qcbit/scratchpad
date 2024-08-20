package sort

// BubbleSort sorts an array of integers
func BubbleSort(A []int) []int {
	if len(A) == 0 {
		return A
	}
    
    n := len(A)
    for i := 0; i < len(A); i++ {
    	for j := n; j > i; j-- {
    		if A[j-1] < A[j] {
    			A[j-1], A[j] = A[j], A[j-1]
    		}
    	}
    }

	return A
}