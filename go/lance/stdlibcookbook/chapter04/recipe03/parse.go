// Parsing the string into date
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("UTC if no time zone. Parsing 2/1/2006 and 31/7/2015")
	t, err := time.Parse("2/1/2006", "31/7/2015")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	fmt.Println()

	fmt.Println("Timezone given: 2/1/2006 3:04 PM MST, 31/7/2015 1:25 AM DST")
	t, err = time.Parse("2/1/2006 3:04 PM MST", "31/7/2015 1:25 AM DST")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	fmt.Println()

	fmt.Println("Parsing for locale")
	t, err = time.ParseInLocation("2/1/2006 3:04 PM ", "31/7/2015 1:25 AM ", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
