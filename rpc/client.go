// client.go
package rpc

import (
	"fmt"
	"net/rpc"
)

func Run_client() {
	// 连接到 RPC 服务器
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer client.Close()

	// 构造请求参数
	args := Args{A: 7, B: 3}
	var reply int
	// 调用 Add 方法
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		fmt.Println("Error calling Arith.Add:", err)
		return
	}

	// 输出结果
	fmt.Printf("Result: %d + %d = %d\n", args.A, args.B, reply)
}
