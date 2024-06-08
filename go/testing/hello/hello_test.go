package main

import "testing"

// TestHello is a suite of tests for the Hello function
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
