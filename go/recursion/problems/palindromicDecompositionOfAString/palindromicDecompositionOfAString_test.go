package palindromicdecompositionofastring

import (
	"reflect"
	"sort"
	"testing"
)

func TestGeneratePalindromicDecompositions(t *testing.T) {
	tests := []struct {
		s        string
		expected []string
	}{
		{"abracadabra", []string{"a|b|r|a|c|ada|b|r|a", "a|b|r|aca|d|a|b|r|a", "a|b|r|a|c|a|d|a|b|r|a"}},
		{"aa", []string{"aa", "a|a"}},
		{"abc", []string{"a|b|c"}},
		{"a", []string{"a"}},
		{"aaaa", []string{"aaaa", "aaa|a", "aa|aa", "aa|a|a", "a|aaa", "a|aa|a", "a|a|aa", "a|a|a|a"}},
		{"abaab", []string{"a|b|a|a|b", "a|b|aa|b", "aba|a|b", "a|baab"}},
		{"abbaeae", []string{"abba|eae", "a|bb|a|eae", "a|b|b|a|eae", "a|bb|aea|e", "a|b|b|aea|e", "abba|e|a|e", "a|bb|a|e|a|e", "a|b|b|a|e|a|e"}},
		{"ikikkiki", []string{"ikikkiki", "iki|kk|iki", "i|k|i|kk|iki", "i|kik|k|iki", "iki|k|k|iki", "i|k|i|k|k|iki", "i|kikkik|i", "i|kik|kik|i", "iki|k|kik|i", "i|k|i|k|kik|i", "i|k|ikki|k|i", "iki|kk|i|k|i", "i|k|i|kk|i|k|i", "i|kik|k|i|k|i", "iki|k|k|i|k|i", "i|k|i|k|k|i|k|i"}},
		{"rooffootfeetteef", []string{"r|ooffoo|t|feetteef", "r|oo|ff|oo|t|feetteef", "r|o|o|ff|oo|t|feetteef", "r|oo|f|f|oo|t|feetteef", "r|o|o|f|f|oo|t|feetteef", "r|o|offo|o|t|feetteef", "r|oo|ff|o|o|t|feetteef", "r|o|o|ff|o|o|t|feetteef", "r|oo|f|f|o|o|t|feetteef", "r|o|o|f|f|o|o|t|feetteef", "r|ooffoo|t|f|eettee|f", "r|oo|ff|oo|t|f|eettee|f", "r|o|o|ff|oo|t|f|eettee|f", "r|oo|f|f|oo|t|f|eettee|f", "r|o|o|f|f|oo|t|f|eettee|f", "r|o|offo|o|t|f|eettee|f", "r|oo|ff|o|o|t|f|eettee|f", "r|o|o|ff|o|o|t|f|eettee|f", "r|oo|f|f|o|o|t|f|eettee|f", "r|o|o|f|f|o|o|t|f|eettee|f", "r|ooffoo|t|f|ee|tt|ee|f", "r|oo|ff|oo|t|f|ee|tt|ee|f", "r|o|o|ff|oo|t|f|ee|tt|ee|f", "r|oo|f|f|oo|t|f|ee|tt|ee|f", "r|o|o|f|f|oo|t|f|ee|tt|ee|f", "r|o|offo|o|t|f|ee|tt|ee|f", "r|oo|ff|o|o|t|f|ee|tt|ee|f", "r|o|o|ff|o|o|t|f|ee|tt|ee|f", "r|oo|f|f|o|o|t|f|ee|tt|ee|f", "r|o|o|f|f|o|o|t|f|ee|tt|ee|f", "r|ooffoo|t|f|e|e|tt|ee|f", "r|oo|ff|oo|t|f|e|e|tt|ee|f", "r|o|o|ff|oo|t|f|e|e|tt|ee|f", "r|oo|f|f|oo|t|f|e|e|tt|ee|f", "r|o|o|f|f|oo|t|f|e|e|tt|ee|f", "r|o|offo|o|t|f|e|e|tt|ee|f", "r|oo|ff|o|o|t|f|e|e|tt|ee|f", "r|o|o|ff|o|o|t|f|e|e|tt|ee|f", "r|oo|f|f|o|o|t|f|e|e|tt|ee|f", "r|o|o|f|f|o|o|t|f|e|e|tt|ee|f", "r|ooffoo|t|f|ee|t|t|ee|f", "r|oo|ff|oo|t|f|ee|t|t|ee|f", "r|o|o|ff|oo|t|f|ee|t|t|ee|f", "r|oo|f|f|oo|t|f|ee|t|t|ee|f", "r|o|o|f|f|oo|t|f|ee|t|t|ee|f", "r|o|offo|o|t|f|ee|t|t|ee|f", "r|oo|ff|o|o|t|f|ee|t|t|ee|f", "r|o|o|ff|o|o|t|f|ee|t|t|ee|f", "r|oo|f|f|o|o|t|f|ee|t|t|ee|f", "r|o|o|f|f|o|o|t|f|ee|t|t|ee|f", "r|ooffoo|t|f|e|e|t|t|ee|f", "r|oo|ff|oo|t|f|e|e|t|t|ee|f", "r|o|o|ff|oo|t|f|e|e|t|t|ee|f", "r|oo|f|f|oo|t|f|e|e|t|t|ee|f", "r|o|o|f|f|oo|t|f|e|e|t|t|ee|f", "r|o|offo|o|t|f|e|e|t|t|ee|f", "r|oo|ff|o|o|t|f|e|e|t|t|ee|f", "r|o|o|ff|o|o|t|f|e|e|t|t|ee|f", "r|oo|f|f|o|o|t|f|e|e|t|t|ee|f", "r|o|o|f|f|o|o|t|f|e|e|t|t|ee|f", "r|ooffoo|t|f|e|ette|e|f", "r|oo|ff|oo|t|f|e|ette|e|f", "r|o|o|ff|oo|t|f|e|ette|e|f", "r|oo|f|f|oo|t|f|e|ette|e|f", "r|o|o|f|f|oo|t|f|e|ette|e|f", "r|o|offo|o|t|f|e|ette|e|f", "r|oo|ff|o|o|t|f|e|ette|e|f", "r|o|o|ff|o|o|t|f|e|ette|e|f", "r|oo|f|f|o|o|t|f|e|ette|e|f", "r|o|o|f|f|o|o|t|f|e|ette|e|f", "r|ooffoo|t|f|ee|tt|e|e|f", "r|oo|ff|oo|t|f|ee|tt|e|e|f", "r|o|o|ff|oo|t|f|ee|tt|e|e|f", "r|oo|f|f|oo|t|f|ee|tt|e|e|f", "r|o|o|f|f|oo|t|f|ee|tt|e|e|f", "r|o|offo|o|t|f|ee|tt|e|e|f", "r|oo|ff|o|o|t|f|ee|tt|e|e|f", "r|o|o|ff|o|o|t|f|ee|tt|e|e|f", "r|oo|f|f|o|o|t|f|ee|tt|e|e|f", "r|o|o|f|f|o|o|t|f|ee|tt|e|e|f", "r|ooffoo|t|f|e|e|tt|e|e|f", "r|oo|ff|oo|t|f|e|e|tt|e|e|f", "r|o|o|ff|oo|t|f|e|e|tt|e|e|f", "r|oo|f|f|oo|t|f|e|e|tt|e|e|f", "r|o|o|f|f|oo|t|f|e|e|tt|e|e|f", "r|o|offo|o|t|f|e|e|tt|e|e|f", "r|oo|ff|o|o|t|f|e|e|tt|e|e|f", "r|o|o|ff|o|o|t|f|e|e|tt|e|e|f", "r|oo|f|f|o|o|t|f|e|e|tt|e|e|f", "r|o|o|f|f|o|o|t|f|e|e|tt|e|e|f", "r|ooffoo|t|f|ee|t|t|e|e|f", "r|oo|ff|oo|t|f|ee|t|t|e|e|f", "r|o|o|ff|oo|t|f|ee|t|t|e|e|f", "r|oo|f|f|oo|t|f|ee|t|t|e|e|f", "r|o|o|f|f|oo|t|f|ee|t|t|e|e|f", "r|o|offo|o|t|f|ee|t|t|e|e|f", "r|oo|ff|o|o|t|f|ee|t|t|e|e|f", "r|o|o|ff|o|o|t|f|ee|t|t|e|e|f", "r|oo|f|f|o|o|t|f|ee|t|t|e|e|f", "r|o|o|f|f|o|o|t|f|ee|t|t|e|e|f", "r|ooffoo|t|f|e|e|t|t|e|e|f", "r|oo|ff|oo|t|f|e|e|t|t|e|e|f", "r|o|o|ff|oo|t|f|e|e|t|t|e|e|f", "r|oo|f|f|oo|t|f|e|e|t|t|e|e|f", "r|o|o|f|f|oo|t|f|e|e|t|t|e|e|f", "r|o|offo|o|t|f|e|e|t|t|e|e|f", "r|oo|ff|o|o|t|f|e|e|t|t|e|e|f", "r|o|o|ff|o|o|t|f|e|e|t|t|e|e|f", "r|oo|f|f|o|o|t|f|e|e|t|t|e|e|f", "r|o|o|f|f|o|o|t|f|e|e|t|t|e|e|f"}},
	}

	for _, test := range tests {
		actual := generate_palindromic_decompositions(test.s)
		if !equalUnordered(actual, test.expected) {
			t.Errorf("For input %q, expected %v but got %v", test.s, test.expected, actual)
		}
	}
}

func equalUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}
