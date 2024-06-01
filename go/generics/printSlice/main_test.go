package main

import (
	"testing"
)

func TestPrint(t *testing.T) {
	slice := []int{1}
	got := printSlice[int](slice)
	want := "0: hello - 1"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
