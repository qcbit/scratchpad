package dutchnationalflag

import (
	"testing"
	"reflect"
)

const ErrorString = `got %v expect %v given %v`

func TestSegregateEvensAndOdds(t *testing.T) {
	testCases := []struct {
		arrays [] string
		expect [] string
	}{
		{[]string{"G", "B", "G", "G", "R", "B", "R", "G"}, []string{"R", "R", "G", "G", "G", "G", "B", "B"}},
		{[]string{"G", "G", "B", "B", "R", "R"}, []string{"R", "R", "G", "G", "B", "B"}},
		{[]string{"G", "B", "R"}, []string{"R", "G", "B"}},
	}
	t.Run("intersection", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, tt.arrays, dutch_flag_sort(tt.arrays), tt.expect)
		}
	})
}

func assertCorrect(t *testing.T, test []string, got, expect []string) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf(ErrorString, got, expect, test)
	}
}
