// Creating the TCP server
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const port = ":8080"

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("Waiting for client...")
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		_, err = io.WriteString(conn, "Received: "+string(msg))
		if err != nil {
			fmt.Println(err)
		}
		conn.Close()
	}
}
