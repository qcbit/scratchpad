// Using regex to find textamin
package main

import (
	"fmt"
	"regexp"
)

const refString = `[{\"email\":\"email@example.com\" \
						"phone\": 555456789},
						{\"email\":\"other@domain.com\" \
						"phone\": 555457890]`

func main() {
	emailRegexp := regexp.MustCompile("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}")
	first := emailRegexp.FindString(refString)
	fmt.Println("First: ")
	fmt.Println(first)

	all := emailRegexp.FindAllString(refString, -1)
	fmt.Println("All: ")
	for _, val := range all {
		fmt.Println(val)
	}
}
