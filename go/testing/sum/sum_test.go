package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}    // fixed size array
	numbers2 := [...]int{1, 2, 3, 4, 5} // fixed size array
	numbers3 := []int{1, 2, 3, 4, 5}    // slice

	assertCorrect := func(t *testing.T, got, want int, numbers any) { // any is a placeholder for any type
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers) // %v is used to print the value of the array. This is the default format.
		}
	}

	t.Run("an array of fixed size", func(t *testing.T) {
		got := Sum(numbers)
		want := 15
		assertCorrect(t, got, want, numbers)
	})

	t.Run("another array of fixed size", func(t *testing.T) {
		got := Sum2(numbers2)
		want := 15
		assertCorrect(t, got, want, numbers2)
	})

	t.Run("a slice", func(t *testing.T) {
		got := Sum3(numbers3)
		want := 15
		assertCorrect(t, got, want, numbers3)
	})

	t.Run("a collection of any size", func(t *testing.T) {
		numbers4 := []int{1, 2, 3}
		got := Sum3(numbers4)
		want := 6
		assertCorrect(t, got, want, numbers4)
	})
}
