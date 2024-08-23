package sort

func MergeSort(A []int) []int {
	helper(A, 0, len(A)-1)
	return A
}

func helper(A []int, start, end int){
	if start == end {
		return 
	}

	mid := (start+end)/2
	helper(A, start, mid)
	helper(A, mid+1, end)
	merge(A, start, mid, end)
}

func merge(A []int, start, mid, end int) {
	aux := make([]int, 0, end-start+1)
	i := start
	j := mid + 1
	for i <= mid && j <= end {
		if A[i] <= A[j] {
			aux = append(aux, A[i])
			i++
		} else {
			aux = append(aux, A[j])
			j++
		}
	}
	for i <= mid {
		aux = append(aux, A[i])
		i++
	}
	for j <= end {
		aux = append(aux, A[j])
		j++
	}

	copy(A[start:end+1], aux)
}