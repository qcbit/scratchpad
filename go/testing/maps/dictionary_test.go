package dictionary

import "testing"

func TestSearch(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}
  
  t.Run("known word", func(t *testing.T){
    got, _ := dictionary.Search("test")
    want := "this is just a test"
    assertString(t, got, want)  
  })
  
  t.Run("Unknown word", func(t *testing.T){
    _, got := dictionary.Search("unknown")
    assertError(t, got, ErrNotFound)
  })
}

func assertString(t *testing.T, got, want string) {
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
  dictionary := Dictionary{}
  word := "test"
  definition := "this is just a test"
  dictionary.Add(word, definition)
  
  assertDefinition(t, dictionary, word, definition)
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