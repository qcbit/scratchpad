package generateAllSubsetsOfASet

import (
	"testing"
	"reflect"
)

const ErrorString = `got %v expect %v given %v`

func TestGenerateAllSubsetsOfASet(t *testing.T) {
	testCases := []struct {
		s string
		want  []string
	}{
		{
			"abc",
			[]string{"abc", "ab", "ac", "a", "bc", "b", "c", ""},
		},
	}

	t.Run("Generate All Subsets Of A Set", func(t *testing.T) {
		for _, tt := range testCases {
			assertCorrect(t, generate_all_subsets(tt.s), tt.want, tt.s)
		}
	})
}

func assertCorrect(t *testing.T, got, expect []string, given string) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf(ErrorString, got, expect, given)
	}
}
