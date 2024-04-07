// Given two strings with the same characters except for one having additional characters in it, return that additional characters.
// https://go.dev/play/p/0W1Bs01bzsz
package main

import (
	"fmt"
)

func main() {
	fmt.Println(xchar("port", "srpot"))
	fmt.Println(xchar("class", "classes"))
}

func xchar(str1, str2 string) []string {

	sstr := func() string {
		if len(str1) < len(str2) {
			return str1
		}
		return str2
	}()

	m := make(map[string]int)
	for _, v := range sstr {
		m[string(v)]++
	}

	lstr := func() string {
		if len(str1) > len(str2) {
			return str1
		}
		return str2
	}()

	x := func() []string {
		var xchars []string
		for _, v := range lstr {
			if m[string(v)] != 1 {
				xchars = append(xchars, string(v))
			}
		}
		return xchars
	}()

	return x
}
