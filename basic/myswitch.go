package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("这是1")
	case 2:
		fmt.Println("这是2")
	case 3:
		fmt.Println("这是3")
	default:
		fmt.Println("这是默认情况")
	}
}
