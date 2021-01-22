// Periodic delayed code execution
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	ticker := time.NewTicker(time.Second)
	stop := make(chan bool)

	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			case <-stop:
				fmt.Println("Goroutine closing")
				return
			}
		}
	}()

	// Block unitl the signal is received
	<-c
	ticker.Stop()

	// Stop the groutine
	stop <- true
	// Wait until the
	<-stop
	fmt.Println("Application stopped")
}
