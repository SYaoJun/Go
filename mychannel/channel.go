package main

import "fmt"

/*
发送和接收默认是阻塞
  读：
	data, ok := <- ch
  写：
	ch <- data
*/
func main() {
	var ch chan bool
	ch = make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		ch <- true
		fmt.Println("goroutine over...")
	}()
	data, ok := <-ch
	if !ok {
		fmt.Println("error")
	}
	fmt.Printf("main over...%t\n", data)

}
