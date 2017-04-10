package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffer size is 2
	ch <- 1
	ch <- 2
	//ch <- 3 // buffer is full, it will block!

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
