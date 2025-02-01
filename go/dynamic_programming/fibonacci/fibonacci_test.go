package fibonacci

import "testing"

const ErrorString = `got %v want %v`

func TestFibonnaci(t *testing.T) {
	t.Helper()
	tests := []struct {
		n    int
		want int
	}{
		{4, 3},
		{5, 5},
		{6, 8},
		{8, 21},
		{10, 55},
		{11, 89},
		{13, 233},
		{14, 377},
		{16, 987},
		{18, 2584},
		{20, 6765},
		{21, 10946},
		{23, 28657},
		{25, 75025},
		{27, 196418},
		{29, 514229},
		{30, 832040},
		{31, 1346269},
		{33, 3524578},
		{34, 5702887},
		{35, 9227465},
		{37, 24157817},
		{38, 39088169},
		{40, 102334155},
		{41, 165580141},
		{42, 267914296},
		{43, 433494437},
		{44, 701408733},
		{46, 1836311903},
	}

	t.Run("fibonacci", func(t *testing.T) {
		for _, tt := range tests {
			got := find_fibonacci(tt.n)
			if got != tt.want {
				t.Errorf(ErrorString, got, tt.want)
			}
		}
	})
}
