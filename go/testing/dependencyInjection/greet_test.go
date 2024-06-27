package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gophers")

	got := buffer.String()
	want := "Hello, Gophers"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
