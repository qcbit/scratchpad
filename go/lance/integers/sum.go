package integers

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, numberList := range numbers {
		sums = append(sums, Sum(numberList))
	}
	return sums
}

// func SumAll(numbers ...[]int) []int {
// 	numberCount := len(numbers)
// 	sums := make([]int, numberCount)
// 	for i, numberList := range numbers {
// 		sums[i] = Sum(numberList)
// 	}
// 	return sums
// }

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, numbersArray := range numbers {
		if len(numbersArray) != 0 {
			tail := numbersArray[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
