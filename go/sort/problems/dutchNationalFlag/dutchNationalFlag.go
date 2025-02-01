// Problem

// Dutch National Flag
// Given some balls of three colors arranged in a line, rearrange them such that all the red balls go first, then green and then blue ones.

// Do rearrange the balls in place. A solution that simply counts colors and overwrites the array is not the one we are looking for.

// This is an important problem in search algorithms theory proposed by Dutch computer scientist Edsger Dijkstra. Dutch national flag has three colors (albeit different from ones used in this problem).

// Example
// {
// "balls": ["G", "B", "G", "G", "R", "B", "R", "G"]
// }
// Output:

// ["R", "R", "G", "G", "G", "G", "B", "B"]
// There are a total of 2 red, 4 green and 2 blue balls. In this order they appear in the correct output.

// Notes
// Constraints:

// 1 <= n <= 100000
// Do this in ONE pass over the array, NOT TWO passes
// Solution is only allowed to use constant extra memory
package dutchnationalflag

func dutch_flag_sort(balls [] string) [] string {
	r, g := -1, -1
	for i := range balls {
		if balls[i] == "G" {
			g++
			balls[i], balls[g] = balls[g], balls[i]
		} else if balls[i] == "R" {
			g++
			balls[i], balls[g] = balls[g], balls[i]
			r++
			balls[g], balls[r] = balls[r], balls[g]
		} 
	}	

	return balls
}
