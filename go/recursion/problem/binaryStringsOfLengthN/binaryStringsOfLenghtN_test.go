package binaryStringsOfLengthN

import (
	"reflect"
	"testing"
)

const ErrorString = `got %v expect %v given %v`

func TestBinaryStringsOfLengthN(t *testing.T) {
	testCases := []struct {
		n    int
		want []string
	}{
		{
			5,
			[]string{"00000", "00001", "00010", "00011", "00100", "00101", "00110", "00111", "01000", "01001", "01010", "01011", "01100", "01101", "01110", "01111", "10000", "10001", "10010", "10011", "10100", "10101", "10110", "10111", "11000", "11001", "11010", "11011", "11100", "11101", "11110", "11111"},
		},
	}

	t.Run("Generate Binary Strings Of Length N", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, get_binary_strings(tt.n), tt.want, tt.n)
		}
	})
}

func assertCorrect(t *testing.T, got, expect []string, given int) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf(ErrorString, got, expect, given)
	}
}
