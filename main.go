package main

import (
	"fmt"
	"yaojungo/mypkg"
	"yaojungo/mystring"
)

func main() {
	mypkg.CustomPkgFunc()
	fmt.Println("hello world")
	s := "hello"
	ss := mystring.Reverse_string(s)
	fmt.Println(ss)
}
