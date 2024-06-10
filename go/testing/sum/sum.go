package main

func Sum(numbers [5]int) int {
	var sum int
	for _, number := range numbers { // pointer-semantics
		sum += number
	}
	return sum
}

func Sum2(numbers [5]int) int {
	var sum int
	for i := range numbers { // value-semantics
		sum += numbers[i]
	}
	return sum
}

func Sum3(numbers []int) int {
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func SumAll(numbers [][]int) int {
	var sum int
	for _, numbersSlice := range numbers {
		sum += Sum3(numbersSlice)
	}
	return sum
}

func SumAll2(numbers ...[]int) int {
	var sum int
	for _, numbersSlice := range numbers {
		sum += Sum3(numbersSlice)
	}
	return sum
}

// SumAll3 is a variadic function that takes a slice of slices of integers and returns a slice of integers.
// This function takes a performance hit because it creates a new slice for each iteration.
func SumAll3(numbers ...[]int) []int {
	var sums []int
	for _, numbersSlice := range numbers {
		sums = append(sums, Sum3(numbersSlice))
	}
	return sums
}

// SumAll4 is a variadic function that takes a slice of slices of integers and returns a slice of integers.
// This function is more performant than SumAll3 because it pre-allocates the slice.
func SumAll4(numbers ...[]int) []int {
	sums := make([]int, len(numbers)) // pre-allocate the slice
	for i, numbersSlice := range numbers {
		sums[i] = Sum3(numbersSlice)
	}
	return sums
}

// SumAllTails is a variadic function that takes a slice of slices of integers and returns a slice of integers.
// It sums the tails of the slices.
func SumAllTails(numbers ...[]int) []int {
	sums := make([]int, len(numbers))
	for i := range numbers {
		sliceLen := len(numbers[i])
		if sliceLen == 0 {
			sums[i] = 0
		} else if sliceLen == 1 {
			sums[i] = numbers[i][0]
		} else {
			sums[i] = Sum3(numbers[i][1:])
		}
	}
	return sums
}