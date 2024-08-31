// Date arithmetics
package main

import (
	"fmt"
	"time"
)

func main() {
	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	t := time.Date(2017, 11, 30, 11, 10, 20, 0, l)
	fmt.Printf("Default date is: %v\n", t)

	r1 := t.Add(72 * time.Hour) // 3 days
	fmt.Printf("Default date -3HRS is: %v\n", r1)

	// day, mo, yr
	r1 = t.AddDate(1, 3, 2)
	fmt.Printf("Default date +1YR +3MTH +2D is: %v\n", r1)
}
