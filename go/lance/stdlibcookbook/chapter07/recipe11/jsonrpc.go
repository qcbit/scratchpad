// Calling the JSON-RPC service
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Result int

type RPCServer struct{}

func (t RPCServer) Add(args *Args, result *Result) error {
	log.Printf("Adding %d to %d\n", args.A, args.B)
	*result = Result(args.A + args.B)
	return nil
}

const addr = "localhost:6565"

func main() {
	go createServer(addr)
	client, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	args := &Args{
		A: 2,
		B: 3,
	}
	var result Result
	if err = client.Call("RPCServer.Add", args, &result); err != nil {
		log.Fatalf("error in RPCServer", err)
	}
	log.Printf("%d+%d=%d\n", args.A, args.B, result)
}

func createServer(addr string) {
	server := rpc.NewServer()
	if err := server.Register(RPCServer{}); err != nil {
		panic(err)
	}

	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalf("Couldn't start listening on %s errors: %s", addr, e)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
