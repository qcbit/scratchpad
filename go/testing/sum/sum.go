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
