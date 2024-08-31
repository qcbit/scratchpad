package main

import (
	"runtime"
	"os/exec"
	"fmt"
	"time"
)

func main() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	proc := exec.Command(cmd, "1")
	proc.Start()
	proc.Wait()

	// After the process terminates the *os.ProcessState contains
	// simple information about the process run
	fmt.Printf("PID: %d\n", proc.ProcessState.Pid())
	fmt.Printf("Process took: %dms\n", proc.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("Exited successfully: %t\n", proc.ProcessState.Success())
}