package main

import (
	"fmt"
	"time"
)

func array_generator(n int, queue chan int) {
	for i := 0; i < n; i++ {
		queue <- i
	}
}

func scheduler(queue chan int, quit chan int) {
	for {
		select {
		case task := <-queue:
			fmt.Println("Get a task: ", task)
		case <-quit:
			fmt.Println("Quit scheduler.")
			return
		default:
			fmt.Println("Idle...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("===== Task sheduler START =====")
	queue := make(chan int)
	quit := make(chan int)
	go scheduler(queue, quit)
	array_generator(10, queue)
	quit <- 0
	fmt.Println("===== Task sheduler END =====")
}
