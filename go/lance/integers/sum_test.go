package integers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("actual: %d\nexpected: %d\ngiven: %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expected := 15
		assertCorrectMessage(t, actual, expected, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		actual := Sum(numbers)
		expected := 6
		assertCorrectMessage(t, actual, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("actual: %v expected: %v", got, want)
		}
	}

	t.Run("any number of collections of any size", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		assertCorrectMessage(t, actual, expected)
	})
}

func ExampleSumAll() {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(actual)
	// Output: [3 9]
}

func TestSumAllTails(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("actual: %v expected: %v", got, want)
		}
	}

	t.Run("Sums the tail", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Safely sums the tail of an empty slice", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		assertCorrectMessage(t, actual, expected)
	})
}
