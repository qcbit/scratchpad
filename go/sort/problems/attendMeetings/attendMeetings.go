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

import "sort"

// ByTime implements sort.Interface for []Interval based on the start field.
type ByTime [][]int
func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)	  { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i][0] < a[j][0] || a[i][0] == a[j][0] && a[i][1] < a[j][1] }

const (
	START = 0
	END = 1
)

func can_attend_all_meetings(intervals [][]int) int {
	sort.Sort(ByTime(intervals))
	canAttend := true
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][START] < intervals[i][END] {
			canAttend = false
			break
		}
	}
	return canAttendMeetings(canAttend)
}

func canAttendMeetings(trueVal bool) int {
	if trueVal {
		return 1
	}
	return 0
}