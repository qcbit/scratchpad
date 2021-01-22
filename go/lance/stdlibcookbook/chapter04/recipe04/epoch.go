// Converting dates to epoch
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Get epoch time")
	t := time.Unix(0, 0)
	fmt.Println(t)
	fmt.Println()

	fmt.Println("Get the epoch")
	epoch := t.Unix()
	fmt.Println(epoch)
	fmt.Println()

	fmt.Println("Current epoch time")
	epochNow := time.Now().Unix()
	fmt.Printf("Eposh time in seconds: %d\n\n", epochNow)

	epochNano := time.Now().UnixNano()
	fmt.Printf("Epoch time in nano-seconds: %d\n", epochNano)
}
