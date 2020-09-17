// Reading/writing from the child process
// Walk through on how to read output 
// and write to the input of the child process
package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "dir"
	} else {
		cmd = "ls"
	}

	proc := exec.Command(cmd)
	buff, err := proc.Output()
	if err != nil {
		panic(err)
	}

	// The output of child process in the form
	// of byte slice printed as a string
	fmt.Println(string(buff))
}