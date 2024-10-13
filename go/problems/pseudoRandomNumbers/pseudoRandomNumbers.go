package pseudorandomnumbers

// Given sets of values for Z, I, M, and the seed, L. Each having no more than 4 digits.
// For each set of values, determine the length of the cycle of pseudo-random numbers
// that will be generated. But be careful: the cycle might not begin with the seed.

func findCycleLength(Z, I, seed, M int) int {
	fast := seed
	slow := seed
	var counter int
	for {
		fast = f(Z, I, f(Z, I, fast, M), M)
		slow = f(Z, I, slow, M)
		if fast == slow {
			third := slow
			for third != slow {
				third = f(Z, I, slow, M)
				counter++
			}
			return counter
		}
	}
}

func f(Z, I, curr, M int) int {
	return (Z*curr + I) % M
}
