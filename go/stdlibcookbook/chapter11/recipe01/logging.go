// Logging customization
package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "custom1: ", log.Ldate|log.Ltime)
	logger.Println("Hello I'm customized")
	loggerEnh := log.New(os.Stdout, "custom2: ", log.Ldate|log.Lshortfile)
	loggerEnh.Println("Hello I'm customized logger 2")
}
