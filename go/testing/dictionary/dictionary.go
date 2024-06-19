package dictionary

import "errors"

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("word already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, definition string) error {
	if d[word] != "" {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}
