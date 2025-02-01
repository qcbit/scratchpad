// Shutdown process gracefully to prevent resource leakage
// The reading from a sigChan is blocking so the program keeps running
// until the Signal is sent through the channel. The sigChan is the
// channel where the Notify function sends the signals.
//
// The main code of the program runs in a new goroutine. This way,
// the work continues while the main function is blocked on the sigChan.
// Once the signal from operation system is sent to process, the sigChan
// receives the signal and the code below the line where the reading from
// the sigChan channel is executed. This code section could be considered
// as the cleanup section.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var writer *os.File

func main() {
	var err error
	writer, err = os.OpenFile(fmt.Sprintf("test_%d.log", time.Now().Unix()), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// The code is running in a goroutine independently.
	// If program is terminated from outside,
	// then let the goroutine know via the closeChan
	closeChan := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-closeChan:
				log.Println("Goroutine close")
				return
			default:
				log.Println("Writing to log")
				io.WriteString(writer, fmt.Sprintf("Logging access %s\n", time.Now().String()))
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT)

	// Blocking read from sigChan where the Notify function sends the signal
	<-sigChan

	// Clean up section after signal received
	close(closeChan)
	releaseAllResources()
	fmt.Println("The application shut down gracefully")
}

func releaseAllResources() {
	io.WriteString(writer, "Application releasing all resources\n")
	writer.Close()
}
