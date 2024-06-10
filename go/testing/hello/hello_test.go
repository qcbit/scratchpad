package main

import "testing"

// TestHello is a suite of tests for the Hello function
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := hello("Jacque", "French")
		want := "Bonjour, Jacque"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish Camelcase", func(t *testing.T) {
		got := hello("Elodie", "SpAnIsH")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
