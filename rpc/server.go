// server.go
package rpc

import (
	"fmt"
	"net"
	"net/rpc"
)

type Arith struct{}

func (t *Arith) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func Run_server() {
	arith := new(Arith)
	// 注册服务
	rpc.Register(arith)
	// 启动监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 1234...")
	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// 为每个连接启动一个 goroutine
		go rpc.ServeConn(conn)
	}
}
