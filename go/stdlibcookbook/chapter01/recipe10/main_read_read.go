package main

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	cmd := "ping"
	timeout := 2 * time.Second

	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	proc := exec.CommandContext(ctx, cmd, "example.com")

	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	proc.Start()
	//Used for more comfortable reading
	s := bufio.NewScanner(stdout)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}