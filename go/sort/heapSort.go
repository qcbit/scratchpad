package sort

func HeapSort(A []int) []int {
	BuildMaxHeap(A)
	i := len(A)
	for i > 0 {
		DeleteMax(A[0:i])
		i--
		if i != 0 {
			heapify(A[:i], 0)
		}
	}
	return A
}

func BuildMaxHeap(A []int) {
	for i := len(A)/2; i >= 0; i--  {
		heapify(A, i)
	}
}

// heapify builds the heap with the given array A and index i
func heapify(A []int, i int) {
	leftChild := 2*i+1
	rightChild := 2*i+2
	largest := i
	if leftChild < len(A) && A[leftChild] > A[largest] {
		largest = leftChild
	}
	if rightChild < len(A) && A[rightChild] > A[largest] {
		largest = rightChild
	}

	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		heapify(A, largest)
	}
}

func DeleteMax(A []int) {
	A[0], A[len(A)-1] = A[len(A)-1], A[0]
}