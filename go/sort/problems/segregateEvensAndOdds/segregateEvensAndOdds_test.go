// Problem
// Segregate Even And Odd Numbers
// Given an array of numbers, rearrange them in-place so that even numbers appear before odd ones.

// Example
// {
// "numbers": [1, 2, 3, 4]
// }
// Output:

// [4, 2, 3, 1]
// The order within the group of even numbers does not matter; same with odd numbers. So the following are also correct outputs: [4, 2, 1, 3], [2, 4, 1, 3], [2, 4, 3, 1].

// Notes
// It is important to practice solving this problem by rearranging numbers in-place.
// There is no need to preserve the original order within the even and within the odd numbers.
// We look for a solution of the linear time complexity that uses constant auxiliary space.
// Constraints:

// 1 <= length of the array <= 105
// 1 <= element of the array <= 109
package segregateevensandodds

import (
	"testing"
	"reflect"
)

const ErrorString = `got %v expect %v given %v`

func TestSegregateEvensAndOdds(t *testing.T) {
	testCases := []struct {
		arrays []int
		expect [] int
	}{
		{[]int{4, 9, 5, 2, 9, 5, 7, 10}, []int{10, 4, 2, 5, 9, 5, 7, 9}},
		{[]int{4, 7, 8}, []int{8, 4, 7}},
		{[]int{1,2,3,4}, []int{4,2,3,1}},
	}
	t.Run("intersection", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, tt.arrays, segregate_evens_and_odds(tt.arrays), tt.expect)
		}
	})
}

func assertCorrect(t *testing.T, test []int, got, expect []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf(ErrorString, got, expect, test)
	}
}
