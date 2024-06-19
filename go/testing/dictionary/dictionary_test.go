package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, gotErr := dictionary.Search("unknown")
		assertError(t, gotErr, ErrNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		want := "this is just a test"
		dictionary.Add("test", want)
		assertDefinition(t, dictionary, "test", want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, def)
	})
}

func assertDefinition(t *testing.T, dict Dictionary, word, definition string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}
