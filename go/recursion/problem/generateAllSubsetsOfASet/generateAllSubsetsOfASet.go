// Problem

// Generate All Subsets Of A Set
// Generate ALL possible subsets of a given set. The set is given in the form of a string s containing distinct lowercase characters 'a' - 'z'.

// Example
// {
// "s": "xy"
// }
// Output:

// ["", "x", "y", "xy"]
// Notes
// Any set is a subset of itself.
// Empty set is a subset of any set.
// Output contains ALL possible subsets of given string.
// Order of strings in the output does not matter. E.g. s = "a", arrays ["", "a"] and ["a", ""] both will be accepted.
// Order of characters in any subset must be same as in the input string. For s = "xy", array ["", "x", "y", "xy"] will be accepted, but ["", "x", "y", "yx"] will not be accepted.
// Constraints:

// 0 <= length of s <= 19
// s only contains distinct lowercase English letters.
package generateAllSubsetsOfASet

func generate_all_subsets(s string) []string {
	results := make([]string, 0)
	helper(s, 0, "", &results)

	return results
}

func helper(s string, ndx int, slate string, results *[]string) {
	if ndx == len(s) {
		*results = append(*results, slate)
		return
	}
	
	slate += string(s[ndx])
	helper(s, ndx+1, slate, results)
	slate = slate[:len(slate)-1]
	helper(s, ndx+1, slate, results)
}