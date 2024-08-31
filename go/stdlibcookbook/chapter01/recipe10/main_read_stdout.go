package main
import (
	"bytes"
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

	buf := bytes.NewBuffer([]byte{})

	// The buffer implements io.Writer interface
	proc.Stdout = buf

	// Wait till process exits to avoid race conditions in this example
	proc.Run()

	fmt.Println(string(buf.Bytes()))
}