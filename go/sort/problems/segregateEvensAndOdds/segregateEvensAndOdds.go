package segregateevensandodds

func segregate_evens_and_odds(numbers [] int) [] int {
	i, j := 0, len(numbers)-1 
	for i < j {
		if numbers[i] % 2 == 0 {
			i++
		} else if numbers[j] % 2 == 1 {
			j--
		} else {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
			j--
		}
	}
	return numbers
}
