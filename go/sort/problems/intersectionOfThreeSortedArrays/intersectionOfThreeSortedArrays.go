// Problem

// Intersection Of Three Sorted Arrays
// Given three arrays sorted in the ascending order, return their intersection sorted array in the ascending order.

// Example One
// {
// "arr1": [2, 5, 10],
// "arr2": [2, 3, 4, 10],
// "arr3": [2, 4, 10]
// }
// Output:

// [2, 10]
// Example Two
// {
// "arr1": [1, 2, 3],
// "arr2": [],
// "arr3": [2, 2]
// }
// Output:

// [-1]
// Example Three
// {
// "arr1": [1, 2, 2, 2, 9],
// "arr2": [1, 1, 2, 2],
// "arr3": [1, 1, 1, 2, 2, 2]
// }
// Output:

// [1, 2, 2]
// Notes
// If the intersection is empty, return an array with one element -1.
// Constraints:

// 0 <= length of each given array <= 105
// 0 <= any value in a given array <= 2 * 106
package intersectionofthreesortedarrays

const (
	SMALLEST = 0
	MIDDLE   = 1
	LARGEST  = 2
)

func find_intersection(arr1 []int, arr2 []int, arr3 []int) []int {
	aux := make([]int, 0)
	arr := setup(arr1, arr2, arr3)
	// for i := 0; i < len(arr[SMALLEST]); i++ {
		// if float32(len(arr[SMALLEST]))/float32(len(arr[LARGEST])) >= 0.5 {
		// 	if search(arr[MIDDLE], arr[SMALLEST][i], i, len(arr[MIDDLE])-1) && search(arr[LARGEST], arr[SMALLEST][i], i, len(arr[LARGEST])-1) {	
		// 			aux = append(aux, arr[SMALLEST][i])
		// 	}
		// } else {
			var i, j, k int
			for i < len(arr[SMALLEST]) && j < len(arr[MIDDLE]) && k < len(arr[LARGEST]) {
				if arr[SMALLEST][i] == arr[MIDDLE][j] && arr[SMALLEST][i] == arr[LARGEST][k] {
					aux = append(aux, arr[SMALLEST][i])
					i++
					j++
					k++
				} else if arr[SMALLEST][i] < arr[MIDDLE][j] || arr[SMALLEST][i] < arr[LARGEST][k] {
					i++
				} else {
					if arr[MIDDLE][j] < arr[SMALLEST][i] {
						j++
					}
					if arr[LARGEST][k] < arr[SMALLEST][i] {
						k++
					}
				} 
			}
		// }
	// }
	if len(aux) > 0 {
		return aux
	}
	return []int{-1}
}

func search(arr []int, target int, start, end int) bool {
		if start <= end {
			mid := start + (end-start)/2
			if arr[mid] == target {
				return true
			} else if arr[mid] > target {
				return search(arr, target, start, mid-1)
			} else {
				return search(arr, target, mid+1, end)
			}
		}
		return false
}

func setup(arr1, arr2, arr3 []int) map[int][]int {
	arr := make(map[int][]int, 3)
	arr1Size, arr2Size, arr3Size := len(arr1), len(arr2), len(arr3)

	if arr1Size == 0 || arr2Size == 0 || arr3Size == 0 {
		return arr
	}

	if arr1Size >= arr2Size && arr1Size >= arr3Size {
		arr[LARGEST] = arr1
		if arr2Size >= arr3Size {
			arr[MIDDLE] = arr2
			arr[SMALLEST] = arr3
		} else {
			arr[MIDDLE] = arr3
			arr[SMALLEST] = arr2
		}
	} else if arr2Size >= arr1Size && arr2Size >= arr3Size {
		arr[LARGEST] = arr2
		if arr1Size >= arr3Size {
			arr[MIDDLE] = arr1
			arr[SMALLEST] = arr3
		} else {
			arr[MIDDLE] = arr3
			arr[SMALLEST] = arr1
		}
	} else if arr3Size >= arr1Size && arr3Size >= arr2Size {
		arr[LARGEST] = arr3
		if arr1Size >= arr2Size {
			arr[MIDDLE] = arr1
			arr[SMALLEST] = arr2
		} else {
			arr[MIDDLE] = arr2
			arr[SMALLEST] = arr1
		}
	}
	return  arr
}
