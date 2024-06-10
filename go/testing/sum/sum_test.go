package main

import (
	"reflect"
	"testing"
)

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

	t.Run("sum all slices of any size", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{1, 2, 3, 4, 5}
		numbers := [][]int{numbers1, numbers2}
		got := SumAll(numbers)
		want := 21
		assertCorrect(t, got, want, numbers)
	})

	t.Run("sum all slices of any size 2", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{1, 2, 3, 4, 5}
		numbers3 := []int{1, 2, 3, 4, 5, 6}
		numbers := [][]int{numbers1, numbers2, numbers3}
		got := SumAll2(numbers1, numbers2, numbers3)
		want := 42
		assertCorrect(t, got, want, numbers)
	})

	t.Run("sum all slices of any size 3", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{1, 2, 3, 4, 5}
		numbers3 := []int{1, 2, 3, 4, 5, 6}
		got := SumAll3(numbers1, numbers2, numbers3)
		want := []int{6, 15, 21}
		// Go does not allow us to compare slices directly. We need to use the reflect.DeepEqual function.
		// The reflect package is part of the Go standard library and provides a way to look at the type of an interface object at runtime.
		// The reflect.DeepEqual is not type safe and will compile even if we pass in the wrong types.
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given, %v", got, want, numbers)
		}
	})

	t.Run("sum all slices of any size 4", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{1, 2, 3, 4, 5}
		numbers3 := []int{1, 2, 3, 4, 5, 6}
		got := SumAll4(numbers1, numbers2, numbers3)
		want := []int{6, 15, 21}
		// Go does not allow us to compare slices directly. We need to use the reflect.DeepEqual function.
		// The reflect package is part of the Go standard library and provides a way to look at the type of an interface object at runtime.
		// The reflect.DeepEqual is not type safe and will compile even if we pass in the wrong types.
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	assertCorrect := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertCorrect(t, got, want)
	})

	t.Run("empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertCorrect(t, got, want)
	})

	t.Run("1 element slices", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{3,4,5})
		want := []int{1, 9}
		assertCorrect(t, got, want)
	})
}
