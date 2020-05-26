package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Lance!", "")
		want := "Hello, Lance!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("World", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Espanol", func(t *testing.T) {
		actual := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in Francais", func(t *testing.T) {
		actual := Hello("", "French")
		expected := "Bonjour, world"
		assertCorrectMessage(t, actual, expected)
	})
}
