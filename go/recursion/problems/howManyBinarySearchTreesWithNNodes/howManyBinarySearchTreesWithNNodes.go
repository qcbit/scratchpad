// How Many Binary Search Trees With N Nodes
// Write a function that returns the number of distinct binary search trees that can be constructed with n nodes. For the purpose of this exercise, do solve the problem using recursion first even if you see some non-recursive approaches.

// Example One
// {
// "n": 1
// }
// Output:

// 1
// Example Two
// {
// "n": 2
// }
// Output:

// 2
// Suppose the values are 1 and 2, then the two trees that are possible are

//    (2)            (1)
//   /       and       \
// (1)                  (2)
// Example Three
// {
// "n": 3
// }
// Output:

// 5
// Suppose the values are 1, 2, 3 then the possible trees are

//        (3)
//       /
//     (2)
//    /
// (1)

//    (3)
//   /
// (1)
//    \
//    (2)

// (1)
//    \
//     (2)
//       \
//        (3)

// (1)
//    \
//     (3)
//    /
// (2)

//    (2)
//   /   \
// (1)    (3)
// Notes
// Constraints:

// 1 <= n <= 16
package howmanybinarysearchtreeswithnnodes

func how_many_bsts(n int) int64 {
	return helper(n)
}

func helper(n int) int64 {
	if n <= 1 {
		return 1
	}

	var ans int64
	for i := 1; i <= n; i++ {
		ans += helper(i-1) * helper(n-i)
	}
	return ans
}
