// Palindromic Decomposition Of A String
// Find all palindromic decompositions of a given string s.

// A palindromic decomposition of string is a decomposition of the string into substrings, such that all those substrings are valid palindromes.

// Example
// {
// "s": "abracadabra"
// }
// Output:

// ["a|b|r|a|c|ada|b|r|a", "a|b|r|aca|d|a|b|r|a", "a|b|r|a|c|a|d|a|b|r|a"]
// Notes
// Any string is its own substring.
// Output should include ALL possible palindromic decompositions of the given string.
// Order of decompositions in the output does not matter.
// To separate substrings in the decomposed string, use | as a separator.
// Order of characters in a decomposition must remain the same as in the given string. For example, for s = "ab", return ["a|b"] and not ["b|a"].
// Strings in the output must not contain whitespace. For example, ["a |b"] or ["a| b"] is incorrect.
// Constraints:

// 1 <= length of s <= 20
// s only contains lowercase English letters.
package palindromicdecompositionofastring

func generate_palindromic_decompositions(s string) []string {
	result := make([]string, 0)
	helper(s, 1, string(s[0]), string(s[0]), &result)
	return result
}

func helper(s string, i int, slate, lastString string, result *[]string) {
	// Base case
	if i == len(s) {
		if isPalindrome(lastString) {
			*result = append(*result, slate)
		}
		return
	}

	// Recursive case
	// Concatenate case
	slate = slate + string(s[i])
	helper(s, i+1, slate, lastString+string(s[i]), result)
	slate = slate[:len(slate)-1]

	// Prune - backtrack
	if !isPalindrome(lastString) {
		return
	}

	// Partition case
	slate = slate + "|" + string(s[i])
	lastString = string(s[i])
	helper(s, i+1, slate, lastString, result)
	slate = slate[:len(slate)-2]
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
