package main

import "fmt"

func test_func(h []int) []int {
	h[0] = 1000
	h = append(h, 2000)
	return h
}
func main() {
	var hello []int
	hello = append(hello, 10)
	hello = append(hello, 20)
	hello = append(hello, 30)
	fmt.Println(hello)
	hello = test_func(hello)
	fmt.Println(hello)
}
