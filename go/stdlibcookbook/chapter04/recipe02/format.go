// Formatting date to string
package main

import (
	"fmt"
	"time"
)

func main() {
	tTime := time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)

	fmt.Println("Use reference value 2006/1/2 for formatting")
	fmt.Printf("tTime is: %s\n\n", tTime.Format("2006/1/2"))

	fmt.Printf("The time is: %s\n\n", tTime.Format("15:04"))

	fmt.Println("Use predefined format time.RFC1123")
	fmt.Printf("The time is: %s\n\n", tTime.Format(time.RFC1123))

	fmt.Println("Use space padding 2006/t/_2")
	fmt.Printf("tTime is: %s\n\n", tTime.Format("2006/t/_2"))

	fmt.Println("Zero padding is done by adding 0 to the date 2006/01/02")
	fmt.Printf("tTime is: %s\n\n", tTime.Format("2006/01/02"))

	fmt.Println("Fractions with leading zeros 15:04:05.00")
	fmt.Printf("tTime is: %s\n\n", tTime.Format("15:04:05.00"))

	fmt.Println("Fraction without leading zeros 15:04:05.999")
	fmt.Printf("tTime is: %s\n\n", tTime.Format("15:04:05.999"))

	fmt.Println("Append format appends the formatted time to given buffer")
	fmt.Println(string(tTime.AppendFormat([]byte("The time is up: "), "03:04PM")))
}
