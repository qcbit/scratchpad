// Problem

// Attend Meetings
// Given a list of meeting intervals where each interval consists of a start and an end time, check if a person can attend all the given meetings such that only one meeting can be attended at a time.

// Example One
// {
// "intervals": [
// [1, 5],
// [5, 8],
// [10, 15]
// ]
// }
// Output:

// 1
// As the above intervals are non-overlapping intervals, it means a person can attend all these meetings.

// Example Two
// {
// "intervals": [
// [1, 5],
// [4, 8]
// ]
// }
// Output:

// 0
// Time 4 - 5 is overlapping in the first and second intervals.

// Notes
// A new meeting can start at the same time the previous one ended.
// Constraints:

// 1 <= number of intervals <= 105
// 0 <= start time < end time <= 109
package attendmeetings

func can_attend_all_meetings(intervals [][]int) int {
	sorted_times := MergeSort(intervals)
	canAttend := true
	start := 0
	end := sorted_times[0][1]
	for i := 1; i < len(sorted_times); i++ {
		if sorted_times[i][start] < end {
			canAttend = false
			break
		}
		end = sorted_times[i][1]
	}
	return canAttendMeetings(canAttend)
}

func canAttendMeetings(trueVal bool) int {
	if trueVal {
		return 1
	}
	return 0
}

func MergeSort(A [][]int) [][]int {
	helper(A, 0, len(A)-1)
	return A
}

func helper(A [][]int, start, end int){
	if start == end {
		return 
	}

	mid := start+(end-start)/2
	helper(A, start, mid)
	helper(A, mid+1, end)
	merge(A, start, mid, end)
}

func merge(A [][]int, start, mid, end int) {
	aux := make([][]int, 0, end-start+1)
	i := start
	j := mid + 1
	for i <= mid && j <= end {
		if A[i][0] <= A[j][0] {
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