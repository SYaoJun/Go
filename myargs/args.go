package main

/*
go run test_argument.go argument1
*/

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os.Args 来获取命令行参数，os.Args[0] 是程序的名称
	args := os.Args[1:] // os.Args 切片中的第一个元素是程序名称，因此我们从第二个元素开始取参数

	// 检查是否提供了参数
	if len(args) == 0 {
		fmt.Println("请提供一个参数")
		return
	}

	// 获取第一个参数
	input := args[0]

	// 在这里可以根据参数执行相应的操作
	// 这里只是简单地打印参数
	fmt.Printf("你提供的参数是：%s\n", input)
}
