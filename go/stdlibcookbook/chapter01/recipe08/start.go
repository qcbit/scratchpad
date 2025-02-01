// Basics of how to execute and handle the child process
package main
import (
	"fmt"
	"bytes"
	"os/exec"
)

func main() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
    prc.Stdout = out
	err := prc.Start() // Start process asynchronously
	if err != nil {
		fmt.Println(err)
	}
	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:\n")
		fmt.Println(out.String())
	}
}