// Problem

// Pascals Triangle
// Pascal’s triangle is a triangular array of the binomial coefficients.
// Write a function that takes an integer value n as input and returns
// a two-dimensional array representing pascal’s triangle.

// pascalTriangleArray is a two-dimensional array of size n * n, where
// pascalTriangleArray[i][j] = BinomialCoefficient(i, j); if j <= i,
// pascalTriangleArray[i][j] = 0; if j > i

// Example
// {
// "n": 4
// }
// Output:

// [
// [1],
// [1, 1],
// [1, 2, 1],
// [1, 3, 3, 1]
// ]
// Notes
// All values in the 2D output array result must be modulo with (10^9 + 7) and size of result[i] for 0 <= i < n should be (i + 1) i.e. 0s for pascalTriangleArray[i][j] = 0; if j > i, should be ignored.
// Constraints:

// 1 <= n <= 1700
package pascalstriangle

func find_pascal_triangle(n int) [][]int {
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = 1
		}
	}

	for row := 2; row < n; row++ {
		for col := 1; col < len(triangle[row])-1; col++ {
			triangle[row][col] = (triangle[row-1][col-1] + triangle[row-1][col]) % 1000000007
		}
	}

	return triangle
}
